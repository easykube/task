package model

import (
	io "io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

//变量
type Var struct {
	//名称
	Var string `yaml:"var"`
	//描述
	Desc string `yaml:"desc"`
	//值或默认值
	Val string `yaml:"val"`

	Type Type `yaml:"type`
}

const (
	varsFilename string = ".var.yaml"
)

type Vars []*Var

func NewVar() *Var {
	return &Var{}
}

func NewVars() Vars {
	return make(Vars, 0)
}

func GetEnvVars() Vars {
	vs := NewVars()
	lines := os.Environ()
	for _, line := range lines {
		list := strings.Split(line, "=")
		if len(list) > 1 && list[0] != "" && list[1] != "" {
			a := NewVar()
			a.Var = list[0]
			a.Val = list[1]
			a.Desc = "环境变量"
			a.Type = T_String
			vs = append(vs)
		}
	}
	return vs
}

func (vs Vars) Add(name, val string) {
	a := NewVar()
	a.Var = name
	a.Val = val
	vs = append(vs, a)
}

func (vs Vars) Load(baseFileName string) (Vars, error) {
	file := baseFileName + varsFilename
	data, err := io.ReadFile(file)
	if err != nil {
		return vs, err
	}
	datayaml := []byte(data)
	err = yaml.Unmarshal(datayaml, vs)
	if err != nil {
		return vs, err
	}
	return vs, err

}

func (vs Vars) Save(baseFileName string) error {

	file := baseFileName + varsFilename
	data, err := yaml.Marshal(vs)
	if err != nil {
		return err
	}
	err = io.WriteFile(file, data, 0777)
	return err

}
