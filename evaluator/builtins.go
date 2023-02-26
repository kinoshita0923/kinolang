package evaluator

import (
	"io"
	"monkey/object"
	"os"
	"strconv"
	"strings"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to 'len' not supported, got %s",
					args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return  newError("argument to 'first' must be ARRAY, got=%s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'last' must be ARRAY, got=%s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if len(arr.Elements) > 0 {
				return arr.Elements[length - 1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to 'rest' must be ARRAY, got=%s",
				args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if len(arr.Elements) > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to 'push' must be ARRAY, got=%s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			
			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"sum": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to 'sum' must be ARRAY. got=%s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			var sum int64 = 0
			for i := 0; i < length; i++ {
				integer, ok := arr.Elements[i].(*object.Integer)
				if !ok {
					return newError("element of array to 'sum' must be INTEGER. got=%s",
						arr.Elements[i].Type())
				}
				sum += integer.Value
			}

			return &object.Integer{Value: sum}
		},
	},
	"max": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to 'max' must be ARRAY. got=%s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			var max int64 = 0

			for _, i := range arr.Elements {
				integer, ok := i.(*object.Integer)
				if !ok {
					return newError("element of array to 'sum' must be INTEGER. got=%s",
						i.Type())
				}
				if integer.Value > max {
					max = integer.Value
				}
			}

			return &object.Integer{Value: max}
		},
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			length := len(args)
			for index, arg := range args {
				io.WriteString(os.Stdout, arg.Inspect())

				if index < length - 1 {
					io.WriteString(os.Stdout, ", ")
				}
			}

			return NULL
		},
	},
	"println": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			length := len(args)
			for index, arg := range args {
				io.WriteString(os.Stdout, arg.Inspect())

				if index < length - 1 {
					io.WriteString(os.Stdout, ", ")
				} else {
					io.WriteString(os.Stdout, "\n")
				}
			}

			return NULL
		},
	},
	"printf": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to 'printf' must be STRING. got=%s",
					args[0].Type())
			}

			length := len(args)
			paramaters := args[1:length]
			character := args[0].(*object.String)
			for _, paramater := range paramaters {
				switch paramater := paramater.(type) {
				case *object.Integer:
					character.Value = strings.Replace(character.Value, "%d", strconv.Itoa(int(paramater.Value)), 1)
				case *object.String:
					character.Value = strings.Replace(character.Value, "%s", paramater.Value, 1)
				}
			}

			io.WriteString(os.Stdout, character.Value)
			io.WriteString(os.Stdout, "\n")

			return NULL
		},
	},
}