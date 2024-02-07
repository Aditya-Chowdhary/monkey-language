package evaluator

import (
	"github.com/Aditya-Chowdhary/Monkey-Interpreter/object"
)

var builtins = map[string]*object.Builtin{
	// TODO if passed multiple arguments, return length of all values?
	"len":   object.GetBuiltinByName("len"),
	"first": object.GetBuiltinByName("first"),
	"last":  object.GetBuiltinByName("last"),
	"rest":  object.GetBuiltinByName("rest"),
	"push":  object.GetBuiltinByName("push"),
	"puts":  object.GetBuiltinByName("puts"),
}
