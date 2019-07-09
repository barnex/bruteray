
#include "textflag.h"

// func Sqrt(x float64) float64
TEXT ·Sqrt(SB), NOSPLIT, $0
	XORPS  X0, X0 // break dependency
	SQRTSD x+0(FP), X0
	MOVSD  X0, ret+8(FP)
	RET

TEXT ·Add4(SB), NOSPLIT, $0
	VADDPD Y0, Y0, Y0
	RET
