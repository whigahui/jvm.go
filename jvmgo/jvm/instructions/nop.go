package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
)

// Do nothing
type nop struct{ NoOperandsInstruction }

func (self *nop) Execute(frame *rtda.Frame) {
	// really do nothing
}
