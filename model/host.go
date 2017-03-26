package model

import (
	io "io/ioutil"
	"path"

	yaml "gopkg.in/yaml.v2"

	"github.com/easykube/sh/remote"
	"github.com/easykube/util"
	"github.com/flosch/pongo2"
)

type IHost interface {
	//获取名称
	GetName() string
	//执行命令
	Run(cmdLine string) (string, error)
	//是否本地主机
	IsLocal() bool
	//关闭连接，清理资源
	Close()
}

//主机定义
type Host struct {
	//主机标识名
	Name string `yaml:"name"`
	//主机名或ip
	Host string `yaml:"host"`
	//端口
	Port int `yaml:"port"`
	//登录的用户
	User string `yaml:"user"`
	//密码
	Password string `yaml:"password"`
	//是否使用WinRm远程连接，否则使用ssh
	UseWinRm bool `yaml:"winrm"`

	//是否是本地主机
	isLocal bool
	//远程session
	session remote.Session
}

type Hosts []*Host

//本机
var LocalHost = &Host{Name: "localhost", Host: "localhost", isLocal: true}

const (
	hostsFilename string = ".host.yaml"
)

//创建Host
func NewHost(name, ip string, port int, user, password string) *Host {
	return &Host{
		Name:     name,
		Host:     ip,
		Port:     port,
		User:     user,
		Password: password,
	}
}

//创建Host数组
func NewHosts() Hosts {
	return make(Hosts, 0)
}

//IHost
//获取名称
func (this *Host) GetName() string {
	return this.Name
}

//IHost
//是否本地主机
func (this *Host) IsLocal() bool {
	return this.isLocal
}

func (this *Host) initSession() {
	if this.session == nil {
		if this.UseWinRm {
			this.session = remote.NewWinRmSession()
		} else {
			this.session = remote.NewSSHSession()
		}
		config := remote.NewSessionConfig()
		config.Host = this.Host
		config.Port = this.Port
		config.User = this.User
		config.Password = this.Password
		this.session.Open(config)
	}
}

//IHost
//执行命令
func (this *Host) Run(cmdLine string) (string, error) {
	if this.isLocal {
		out, err := util.ExecCmdLine(cmdLine)
		return out, err
	} else {
		this.initSession()
		return this.session.Run(cmdLine)
	}
	return "", nil
}

//IHost
//关闭连接，清理资源
func (this *Host) Close() {
	if this.isLocal {

	} else {
		if this.session != nil {
			err := this.session.Close()
			if err != nil {
				util.LogError("Host.Close", err)
			}
		}
	}

}

//添加主机到数组
func (hs Hosts) AddHost(h *Host) Hosts {
	hs = append(hs, h)
	return hs
}

func (hs Hosts) write(rootPath string) (string, error) {
	file := path.Join(rootPath, hostsFilename)
	tpl, err := pongo2.FromFile(file)
	if err != nil {
		panic(err)
	}

	out, err := tpl.Execute(pongo2.Context{"hosts": hs})
	if err != nil {
		panic(err)
	}

	return out, nil
}

func (hs Hosts) Load(baseFileName string) error {
	file := baseFileName + hostsFilename
	tpl, err := pongo2.FromFile(file)
	if err != nil {
		return err
	}

	data, err := tpl.Execute(pongo2.Context{"hosts": hs})
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(data), &hs)

	return err

}

func (hs Hosts) Save(baseFileName string) error {
	file := baseFileName + hostsFilename
	data, err := yaml.Marshal(hs)
	if err != nil {
		return err
	}
	err = io.WriteFile(file, data, 0777)
	return err
}
