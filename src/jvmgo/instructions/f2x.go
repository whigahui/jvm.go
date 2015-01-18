package instructions

import "jvmgo/rtda"

// Convert float to double
type f2d struct {}
func (self *f2d) fetchOperands(bcr *BytecodeReader) {}
func (self *f2d) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    f := stack.PopFloat()
    d := float64(f)
    stack.PushDouble(d)
}

// Convert float to int
type f2i struct {}
func (self *f2i) fetchOperands(bcr *BytecodeReader) {}
func (self *f2i) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    f := stack.PopFloat()
    i := int32(f)
    stack.PushInt(i)
}

// Convert float to long
type f2l struct {}
func (self *f2l) fetchOperands(bcr *BytecodeReader) {}
func (self *f2l) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    f := stack.PopFloat()
    l := int64(f)
    stack.PushLong(l)
}
