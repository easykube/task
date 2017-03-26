package model

type Type int8

const T_Int = 1
const T_String = 2
const T_Float = 3

func GetNameByType(t Type) string {
	switch t {
	case T_Int:
		return "int"
	case T_String:
		return "string"
	case T_Float:
		return "float"
	default:
		return "int"
	}
}
