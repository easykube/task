package model

//项目
type Project struct {
	//项目名称
	Project string `yaml:"project"`

	//项目描述
	Desc string `yaml:"desc"`

	//主任务名称
	Main string `yaml:"main"`

	//变量定义
	Vars []Var `yaml:"vars"`

	//任务列表
	Tasks []Task `yaml:"tasks"`
}

func NewProject() *Project {
	return &Project{}
}

//在主机上执行项目任务
func (this *Project) Execute(host IHost) {

}
