package model

import "testing"

func Test1(t *testing.T) {
	f := NewFunc()
	f.Func = "TestHello"
	f.Desc = "测试Hello"
	f.Args = f.Args.Add("name", "tom", T_String)
	f.CmdLine = "cmd /k echo hello,{{name}}"
	f.Return = "{{out}}"
	f.Execute(f.Args)
	println("return ===")
	println(f.Return)

}
