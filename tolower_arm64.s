// ARM64 SWAR (SIMD-within-a-register) implementation for Go Assembler
#include "textflag.h"

// func ToLowerAsm(buf []byte)
TEXT ·ToLowerAsm(SB), NOSPLIT, $0-24
    MOVD    buf_data+0(FP), R0      // R0 = pointer to buf
    MOVD    buf_len+8(FP), R1       // R1 = length of buf

    // Create constants for the SWAR calculation
    MOVD    $0x4141414141414141, R2  // 'A' repeated
    MOVD    $0x5A5A5A5A5A5A5A5A, R3  // 'Z' repeated
    MOVD    $0x2020202020202020, R4  // Space char (0x20) repeated
    MOVD    $0x8080808080808080, R5  // High-bit mask

loop:
    CMP     $8, R1
    BLT     tail                    // If less than 8 bytes, go to tail processing

    // Load 8 bytes (64 bits) from memory
    MOVD    (R0), R6

    // This is a branchless, arithmetic way to identify uppercase letters
    // within the 64-bit register R6.
    // 1. Create a mask for bytes >= 'A'.
    //    (R6 - 'A') will have its high bit clear for bytes >= 'A'.
    //    We add 0x80 to flip this, so the high bit is SET for bytes >= 'A'.
    SUB     R2, R6, R7
    ADD     R5, R7, R7              // R7 has high bit set for bytes >= 'A'

    // 2. Create a mask for bytes <= 'Z'.
    //    ('Z' - R6) will have its high bit clear for bytes <= 'Z'.
    SUB     R6, R3, R8
    ADD     R5, R8, R8              // R8 has high bit set for bytes <= 'Z'

    // 3. Combine the masks. The high bit is set only if both conditions are true.
    AND     R7, R8, R7              // R7 has high bit set for 'A' <= byte <= 'Z'

    // 4. Shift the high bit to the 0x20 position.
    LSR     $2, R7, R7
    AND     R4, R7, R7              // R7 now contains 0x20 for uppercase bytes, 0 otherwise

    // 5. Apply the mask to the original data to convert to lowercase.
    ORR     R6, R7, R6

    // Store the result back to memory
    MOVD    R6, (R0)

    // Advance pointers and counters
    ADD     $8, R0
    SUB     $8, R1
    B       loop

tail:
    // Process remaining bytes one by one
    CBZ     R1, done
    MOVBU   (R0), R2
    SUB     $'A', R2, R3
    CMP     $25, R3
    BHI     notupper
    ORR     $0x20, R2, R2
    MOVB    R2, (R0)
notupper:
    ADD     $1, R0
    SUB     $1, R1
    B       tail

done:
    RET
