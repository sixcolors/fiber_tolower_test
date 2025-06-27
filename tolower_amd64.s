// x86-64 SSE2 implementation for Go Assembler
#include "textflag.h"

// Define constants in memory that we can load into SSE registers.
DATA masks<>+0(SB)/8, $0x4141414141414141  // 'A' repeated
DATA masks<>+8(SB)/8, $0x5A5A5A5A5A5A5A5A  // 'Z' repeated
DATA masks<>+16(SB)/8, $0x2020202020202020 // 0x20 repeated
GLOBL masks<>(SB), (RODATA+NOPTR), $24

// func ToLowerAsm(buf []byte)
TEXT ·ToLowerAsm(SB), NOSPLIT, $0-24
    MOVQ    buf_data+0(FP), DI      // DI = pointer to buf
    MOVQ    buf_len+8(FP), CX       // CX = length of buf

    // Load constants into SSE registers
    MOVOU   masks<>+0(SB), X2       // X2 contains 'A's
    MOVOU   masks<>+8(SB), X3       // X3 contains 'Z's
    MOVOU   masks<>+16(SB), X4      // X4 contains 0x20s

loop16:
    CMPQ    CX, $16
    JL      tail                    // If less than 16 bytes, go to tail processing

    // Load 16 bytes of data from memory
    MOVOU   (DI), X0

    // Create a copy to work with
    MOVOU   X0, X1

    // Create a mask for bytes that are uppercase letters.
    // This is a branchless, arithmetic way to identify them.
    // 1. Find bytes >= 'A'. PCMPGTB finds > not >=, so we compare to 'A'-1.
    //    The constant for 'A' is already loaded, so we can use it directly.
    //    We want to find where X1 > 'A'-1, which is equivalent to X1 >= 'A'.
    //    To do this without another constant, we can use PCMEQ and then invert.
    //    A simpler way is to use the logic from the ARM version:
    //    (c >= 'A') AND (c <= 'Z')
    PCMPGTB X2, X1                  // X1 = (X0 > 'A'), effectively (X0 >= 'A') if we treat 'A' as 'A'-1
    PCMPEQB X2, X0                  // Find bytes equal to 'A'
    POR     X0, X1                  // Combine to get (X0 >= 'A')
    
    // 2. Find bytes <= 'Z'
    MOVOU   X0, X5
    PCMPGTB X3, X5                  // X5 = (X0 > 'Z'), so ~X5 is (X0 <= 'Z')
    
    // 3. Combine the masks
    PAND    X1, X5                  // X5 now has high bits set for uppercase letters
    
    // 4. Create the conversion mask
    PAND    X4, X5                  // X5 now contains 0x20 for uppercase bytes

    // 5. Apply the mask to the original data
    POR     X0, X5, X0

    // Store the result back to memory
    MOVOU   X0, (DI)

    // Advance pointers and counters
    ADDQ    $16, DI
    SUBQ    $16, CX
    JMP     loop16

tail:
    // Process remaining bytes one by one
    TESTQ   CX, CX
    JZ      done
    MOVBLU  (DI), AX
    SUBQ    $'A', AX
    CMPQ    $25, AX
    JA      notupper
    ORB     $0x20, (DI)
notupper:
    INCQ    DI
    DECQ    CX
    JMP     tail

done:
    RET
