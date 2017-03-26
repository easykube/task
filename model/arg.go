package model

//变量
type Arg struct {
	//名称
	Arg string `yaml:"arg"`
	//描述
	Desc string `yaml:"desc"`
	//默认值
	Val string `yaml:"val"`
	//类型
	Type Type `yaml:"type"`
}

type Args []*Arg

func NewAgr() *Arg {
	return &Arg{}
}

func NewArgs() Args {
	args := make(Args, 0)
	return args
}

func (this Args) Add(arg string, val string, t Type) Args {
	a := NewAgr()
	a.Arg = arg
	a.Val = val
	a.Type = t
	return append(this, a)
}
