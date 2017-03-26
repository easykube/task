package model

import (
	io "io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//项目
type Project struct {
	//项目名称
	Project string `yaml:"project"`

	//项目描述
	Desc string `yaml:"desc"`

	//主任务名称
	Main string `yaml:"main"`

	//变量定义
	vars Vars

	//任务列表
	Tasks []Task `yaml:"tasks"`
}

const (
	projectExt string = ".yaml"
)

//创建项目实例
func NewProject() *Project {
	return &Project{vars: NewVars()}
}

//在主机上执行项目任务
func (this *Project) Execute(host IHost) {

}

//加载项目
//projectFile,含路径，不含后缀
func (this *Project) Load(projectFile string) error {
	vs, err := this.vars.Load(projectFile)
	if err == nil {
		this.vars = vs
	}
	file := projectFile + projectExt
	data, err := io.ReadFile(file)
	if err != nil {
		return err
	}
	datayaml := []byte(data)
	err = yaml.Unmarshal(datayaml, this)
	if err != nil {
		return err
	}
	return nil

}

//保存项目
//projectFile,含路径，不含后缀
func (this *Project) Save(projectFile string) error {
	this.vars.Save(projectFile)
	file := projectFile + projectExt
	data, err := yaml.Marshal(this)
	if err != nil {
		return err
	}
	err = io.WriteFile(file, data, 0777)
	return err

}
