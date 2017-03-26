package model

//管理中心
type Mgr struct {
	//功能列表
	funcList map[string]IFunc
}

var mgr = &Mgr{}

//获取实例
func GetMgr() *Mgr {
	return mgr
}

//注册Func
func (this *Mgr) Register(f IFunc) {
	this.funcList[f.GetName()] = f
}

//按名称查询Func
func (this *Mgr) GetFunc(name string) IFunc {
	return this.funcList[name]
}
