package vm

import (
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/code"
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/object"
)

type Frame struct {
	fn *object.CompiledFunction
	ip int
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
