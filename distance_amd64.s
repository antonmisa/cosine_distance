#include "textflag.h"

// func _cosineSimAsm(a []float32, b []float32) float32
// Requires: AVX, FMA3, SSE
TEXT ·_cosineSimAsm(SB), NOSPLIT, $0-52
        MOVQ   a_base+0(FP), AX
        MOVQ   b_base+24(FP), CX
        MOVQ   a_len+8(FP), DX
        VXORPS Y0, Y0, Y0
        VXORPS Y1, Y1, Y1
        VXORPS Y2, Y2, Y2

blockloop:
        CMPQ        DX, $0x00000008
        JL          tail
        VMOVUPS     (AX), Y3
        VMOVUPS     (CX), Y4
        VFMADD231PS (CX), Y3, Y0
        VFMADD231PS (AX), Y3, Y1
        VFMADD231PS (CX), Y4, Y2
        ADDQ        $0x00000020, AX
        ADDQ        $0x00000020, CX
        SUBQ        $0x00000008, DX
        JMP         blockloop

tail:
        VXORPS X3, X3, X3
        VXORPS X4, X4, X4
        VXORPS X5, X5, X5

tailloop:
        CMPQ        DX, $0x00000000
        JE          reduce
        VMOVSS      (AX), X6
        VMOVSS      (CX), X7
        VFMADD231SS (CX), X6, X3
        VFMADD231SS (AX), X6, X4
        VFMADD231SS (CX), X7, X5
        ADDQ        $0x00000004, AX
        ADDQ        $0x00000004, CX
        DECQ        DX
        JMP         tailloop

reduce:
        VEXTRACTF128 $0x01, Y0, X6
        VADDPS       X0, X6, X0
        VADDPS       X0, X3, X0
        VHADDPS      X0, X0, X0
        VHADDPS      X0, X0, X0
        VEXTRACTF128 $0x01, Y1, X3
        VADDPS       X1, X3, X1
        VADDPS       X1, X4, X1
        VHADDPS      X1, X1, X1
        VHADDPS      X1, X1, X1
        VEXTRACTF128 $0x01, Y2, X3
        VADDPS       X2, X3, X2
        VADDPS       X2, X5, X2
        VHADDPS      X2, X2, X2
        VHADDPS      X2, X2, X2
        VSQRTPS      X1, X1
        VSQRTPS      X2, X2
        VMULPS       X2, X1, X1
        VDIVPS       X1, X0, X0
        MOVSS        X0, ret+48(FP)
        RET
