// ARM64 optimized implementation with offset support
#include "textflag.h"

// func toLowerAsmWithOffset(buf []byte, offset int, length int)
// FP: data+0, len+8, cap+16, offset+24, length+32  → $0-40
TEXT ·toLowerAsmWithOffset(SB), NOSPLIT, $0-40
    MOVD    buf_data+0(FP), R0      // R0 = pointer to buf
    MOVD    offset+24(FP), R2       // R2 = starting offset  (FIXED)
    MOVD    length+32(FP), R1       // R1 = length to process (FIXED)

    // Early return if length is zero
    CBZ     R1, done

    // Apply offset to data pointer
    ADD     R2, R0, R0

    // Create constants for the SWAR calculation
    MOVD    $0x4141414141414141, R2   // 'A' repeated
    MOVD    $0x5A5A5A5A5A5A5A5A, R3   // 'Z' repeated
    MOVD    $0x2020202020202020, R4   // 0x20 repeated
    MOVD    $0x8080808080808080, R5   // high-bit mask

// Main loop - process 64 bytes per iteration
loop64:
    CMP     $64, R1
    BLT     loop32

    MOVD    (R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, (R0)

    MOVD    8(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 8(R0)

    MOVD    16(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 16(R0)

    MOVD    24(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 24(R0)

    MOVD    32(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 32(R0)

    MOVD    40(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 40(R0)

    MOVD    48(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 48(R0)

    MOVD    56(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 56(R0)

    ADD     $64, R0
    SUB     $64, R1
    CMP     $64, R1
    BGE     loop64

loop32:
    CMP     $32, R1
    BLT     loop16

    MOVD    (R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, (R0)

    MOVD    8(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 8(R0)

    MOVD    16(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 16(R0)

    MOVD    24(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 24(R0)

    ADD     $32, R0
    SUB     $32, R1

loop16:
    CMP     $16, R1
    BLT     loop8

    MOVD    (R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, (R0)

    MOVD    8(R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, 8(R0)

    ADD     $16, R0
    SUB     $16, R1

loop8:
    CMP     $8, R1
    BLT     tail

    MOVD    (R0), R6
    SUB     R2, R6, R7
    ADD     R5, R7, R7
    SUB     R6, R3, R8
    ADD     R5, R8, R8
    AND     R7, R8, R7
    LSR     $2, R7, R7
    AND     R4, R7, R7
    ORR     R6, R7, R6
    MOVD    R6, (R0)

    ADD     $8, R0
    SUB     $8, R1

tail:
    CBZ     R1, done

loop_tail:
    MOVBU   (R0), R6
    SUB     $'A', R6, R7
    CMP     $25, R7
    BHI     notupper
    ORR     $0x20, R6, R6
    MOVB    R6, (R0)
notupper:
    ADD     $1, R0
    SUB     $1, R1
    CBNZ    R1, loop_tail

done:
    RET
