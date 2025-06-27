// ARM64 optimized SWAR (SIMD-within-a-register) implementation for Go Assembler
// Cache-line aligned processing: 64, 32, 16, 8 bytes
#include "textflag.h"

// func ToLowerAsm(buf []byte)
TEXT ·ToLowerAsm(SB), NOSPLIT, $0-24
    MOVD    buf_data+0(FP), R0      // R0 = pointer to buf
    MOVD    buf_len+8(FP), R1       // R1 = length of buf
    
    // Early return if empty
    CBZ     R1, done

    // Create constants for the SWAR calculation
    MOVD    $0x4141414141414141, R2   // 'A' repeated
    MOVD    $0x5A5A5A5A5A5A5A5A, R3   // 'Z' repeated
    MOVD    $0x2020202020202020, R4   // Space char (0x20) repeated
    MOVD    $0x8080808080808080, R5   // High-bit mask

// Main loop - process 64 bytes (aligned with cache line) per iteration
loop64:
    CMP     $64, R1
    BLT     loop32                   // If less than 64 bytes, handle smaller chunks

    // Process 8 chunks of 8 bytes each (64 bytes total)
    // Chunk 1
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
    
    // Chunk 2
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
    
    // Chunk 3
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
    
    // Chunk 4
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
    
    // Chunk 5
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
    
    // Chunk 6
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
    
    // Chunk 7
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
    
    // Chunk 8
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

    // Advance pointer and counter
    ADD     $64, R0
    SUB     $64, R1
    CMP     $64, R1
    BGE     loop64

// Process 32 bytes (half cache line)
loop32:
    CMP     $32, R1
    BLT     loop16                   // If less than 32 bytes, handle smaller chunks

    // Process 4 chunks of 8 bytes each (32 bytes total)
    // Chunk 1
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
    
    // Chunk 2
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
    
    // Chunk 3
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
    
    // Chunk 4
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

    // Advance pointer and counter
    ADD     $32, R0
    SUB     $32, R1

// Process 16 bytes
loop16:
    CMP     $16, R1
    BLT     loop8                    // If less than 16 bytes, handle smaller chunks

    // Process 2 chunks of 8 bytes each (16 bytes total)
    // Chunk 1
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
    
    // Chunk 2
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

    // Advance pointer and counter
    ADD     $16, R0
    SUB     $16, R1

// Process 8 bytes
loop8:
    CMP     $8, R1
    BLT     tail                     // If less than 8 bytes, handle individually
    
    // Process 1 chunk of 8 bytes
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

    // Advance pointer and counter
    ADD     $8, R0
    SUB     $8, R1

// Process remaining bytes individually
tail:
    CBZ     R1, done                 // If no bytes left, we're done
    
loop_tail:
    MOVBU   (R0), R2
    SUB     $'A', R2, R3
    CMP     $25, R3
    BHI     notupper
    ORR     $0x20, R2, R2
    MOVB    R2, (R0)
notupper:
    ADD     $1, R0
    SUB     $1, R1
    CBNZ    R1, loop_tail            // Continue if there are more bytes

done:
    RET
