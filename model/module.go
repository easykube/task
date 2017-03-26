package model

import (
	"fmt"

	"github.com/flosch/pongo2"
)

//模块
type Module struct {
	//名称
	Module string `yaml:"module"`
	//描述
	Desc string `yaml:"desc"`
	//变量定义
	Vars Vars `yaml:"vars"`
	//参数列表
	Args Args `yaml:"args"`
	//功能组合
	Task []Func `yaml:"task"`
	//返回值定义
	Return string `yaml:"return"`
}

//创建Func
func NewModule() *Module {
	return &Module{Args: NewArgs()}
}

//IFunc
//获取名称
func (this *Module) GetName() string {
	return this.Module
}

//IFunc
//获取描述
func (this *Module) GetDesc() string {
	return this.Desc
}

//IFunc
//获取参数定义
func (this *Module) GetArgs() Args {
	return this.Args
}

//IFunc
//执行操作
func (this *Module) Execute(host IHost, args Args) {
	strTemp := ""
	tpl, err := pongo2.FromString(strTemp)
	if err != nil {
		panic(err)
	}
	context := pongo2.Context{}
	AddArgsToContext(args, context)
	println("context")
	for k, v := range context {
		println(k + "=")
		println(v)
	}
	line, err := tpl.Execute(context)
	if err != nil {
		panic(err)
	}
	fmt.Println("cmdline===")
	fmt.Println(line)
	out, err := host.Run(line)
	if err != nil {
		panic(err)
	}
	fmt.Println("out===")
	fmt.Println(out)

	context = pongo2.Context{}
	context["out"] = out
	tpl, err = pongo2.FromString(this.Return)
	out, err = tpl.Execute(context)
	if err != nil {
		panic(err)
	}
	this.Return = out
}
