package model

import (
	"fmt"

	"github.com/flosch/pongo2"
)

//功能接口
type IFunc interface {
	//功能名称
	GetName() string
	//描述
	GetDesc() string
	//参数定义
	GetArgs() Args
	//执行
	Execute(host IHost, args Args)
}

//功能
type Func struct {
	//名称
	Func string `yaml:"func"`
	//描述
	Desc string `yaml:"desc"`
	//参数列表
	Args Args `yaml:"args"`
	//命令行cmd-line
	CmdLine string `yaml:"desc"`
	//返回值定义
	Return string `yaml:"return"`
}

//创建Func
func NewFunc() *Func {
	return &Func{Args: NewArgs()}
}

//IFunc
//获取名称
func (this *Func) GetName() string {
	return this.Func
}

//IFunc
//获取描述
func (this *Func) GetDesc() string {
	return this.Desc
}

//IFunc
//获取参数定义
func (this *Func) GetArgs() Args {
	return this.Args
}

//IFunc
//执行操作
func (this *Func) Execute(host IHost, args Args) {
	tpl, err := pongo2.FromString(this.CmdLine)
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
