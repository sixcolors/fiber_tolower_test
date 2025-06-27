// x86-64 SSE2 branchless arithmetic implementation for Go Assembler
#include "textflag.h"

// Define constants in memory that we can load into SSE registers.
DATA masks<>+0(SB)/8, $0x4040404040404040
DATA masks<>+8(SB)/8, $0x4040404040404040  // 'A'-1 repeated for 16 bytes
DATA masks<>+16(SB)/8, $0x5A5A5A5A5A5A5A5A
DATA masks<>+24(SB)/8, $0x5A5A5A5A5A5A5A5A // 'Z' repeated for 16 bytes
DATA masks<>+32(SB)/8, $0x2020202020202020
DATA masks<>+40(SB)/8, $0x2020202020202020 // 0x20 repeated for 16 bytes
GLOBL masks<>(SB), (RODATA+NOPTR), $48

// func ToLowerAsm(buf []byte)
TEXT ·ToLowerAsm(SB), NOSPLIT, $0-24
    MOVQ    buf_data+0(FP), DI      // DI = pointer to buf
    MOVQ    buf_len+8(FP), CX       // CX = length of buf

    // Skip empty buffers
    TESTQ   CX, CX
    JZ      done

    // Load constants into SSE registers
    MOVOU   masks<>+0(SB), X2       // X2 contains 'A'-1
    MOVOU   masks<>+16(SB), X3      // X3 contains 'Z'
    MOVOU   masks<>+32(SB), X4      // X4 contains 0x20s

    // Process 48 bytes per loop iteration (3x unrolling)
    CMPQ    CX, $48
    JL      loop16_check            // If less than 48 bytes, skip unrolled loop

loop48:
    // Prefetch data that will be needed soon
    PREFETCHT0 128(DI)

    // Process first 16 bytes
    MOVOU   (DI), X0
    MOVOU   X0, X6                  // Preserve original data in X6
    MOVOU   X0, X1                  // X1 = copy of original data for comparison
    PCMPGTB X2, X1                  // X1 has 0xFF where c > 'A'-1 (i.e., c >= 'A')
    MOVOU   X0, X5                  // X5 = copy of original data for comparison
    PCMPGTB X3, X5                  // X5 has 0xFF where c > 'Z'
    PCMPEQB X7, X7                  // X7 = all 1s (0xFF...FF)
    PXOR    X5, X7                  // X7 has 0xFF where c <= 'Z'
    PAND    X7, X1                  // Combine masks
    PAND    X4, X1                  // X1 now contains 0x20 for uppercase bytes
    POR     X1, X6                  // Apply to original data
    MOVOU   X6, (DI)                // Store result back

    // Process second 16 bytes
    MOVOU   16(DI), X0
    MOVOU   X0, X6
    MOVOU   X0, X1
    PCMPGTB X2, X1
    MOVOU   X0, X5
    PCMPGTB X3, X5
    PCMPEQB X7, X7
    PXOR    X5, X7
    PAND    X7, X1
    PAND    X4, X1
    POR     X1, X6
    MOVOU   X6, 16(DI)

    // Process third 16 bytes
    MOVOU   32(DI), X0
    MOVOU   X0, X6
    MOVOU   X0, X1
    PCMPGTB X2, X1
    MOVOU   X0, X5
    PCMPGTB X3, X5
    PCMPEQB X7, X7
    PXOR    X5, X7
    PAND    X7, X1
    PAND    X4, X1
    POR     X1, X6
    MOVOU   X6, 32(DI)

    // Advance pointers and counters
    ADDQ    $48, DI
    SUBQ    $48, CX
    CMPQ    CX, $48
    JAE     loop48                  // Continue if at least 48 bytes remain

loop16_check:
    CMPQ    CX, $16
    JL      tail                    // If less than 16 bytes, go to tail processing

loop16:
    // Load 16 bytes of data from memory
    MOVOU   (DI), X0
    MOVOU   X0, X6                  // Preserve original data in X6

    // This is a branchless, arithmetic way to identify uppercase letters ('A' <= c <= 'Z')
    MOVOU   X0, X1                  // X1 = copy of original data for comparison
    PCMPGTB X2, X1                  // X1 has 0xFF where c > 'A'-1 (i.e., c >= 'A')
    MOVOU   X0, X5                  // X5 = copy of original data for comparison
    PCMPGTB X3, X5                  // X5 has 0xFF where c > 'Z'
    PCMPEQB X7, X7                  // X7 = all 1s (0xFF...FF)
    PXOR    X5, X7                  // X7 has 0xFF where c <= 'Z'
    PAND    X7, X1                  // Combine masks
    PAND    X4, X1                  // X1 now contains 0x20 for uppercase bytes
    POR     X1, X6                  // Apply to original data
    MOVOU   X6, (DI)                // Store result back

    // Advance pointers and counters
    ADDQ    $16, DI
    SUBQ    $16, CX
    CMPQ    CX, $16
    JAE     loop16                  // Continue if at least 16 bytes remain

tail:
    // Process remaining bytes one by one
    TESTQ   CX, CX
    JZ      done

    // Use SIMD for tail if we have at least 8 bytes
    CMPQ    CX, $8
    JL      byte_tail

    // Process 8 bytes using the same SIMD technique
    MOVQ    (DI), R8                // Load 8 bytes into general register
    MOVQ    R8, R9                  // Preserve original data
    
    // Convert to XMM register for SIMD operations
    MOVQ    R8, X0
    MOVOU   X0, X1
    PCMPGTB X2, X1                  // X1 has 0xFF where c > 'A'-1
    MOVOU   X0, X5
    PCMPGTB X3, X5                  // X5 has 0xFF where c > 'Z'
    PCMPEQB X7, X7
    PXOR    X5, X7
    PAND    X7, X1
    PAND    X4, X1
    POR     X1, X0
    
    // Convert back to general register
    MOVQ    X0, R8
    MOVQ    R8, (DI)
    
    ADDQ    $8, DI
    SUBQ    $8, CX

byte_tail:
    // Process remaining bytes one by one
    TESTQ   CX, CX
    JZ      done

loop_tail:
    MOVB    (DI), AL                // Move byte into 8-bit AL register
    CMPB    AL, $'A'
    JB      notupper                // Jump if below 'A'
    CMPB    AL, $'Z'
    JA      notupper                // Jump if above 'Z'
    ORB     $0x20, AL               // It's an uppercase letter, convert it
    MOVB    AL, (DI)
notupper:
    INCQ    DI
    DECQ    CX
    JNZ     loop_tail               // Jump if CX is not zero

done:
    RET
