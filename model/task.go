package model

//模块
type Task struct {
	//名称
	Task string `yaml:"task"`
	//描述
	Desc string `yaml:"desc"`
	//功能组合
	Exec []Func `yaml:"exec"`
}

func NewTask() *Task {
	return &Task{}
}

//在主机上执行任务
func (this *Task) Execute(host IHost) {

}
