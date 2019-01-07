package Iterator
//工资描述模型
type PayModel struct {
	UserName string
	Pay float64
}
//工资管理类
type PayManager struct {
	//工资列表
	List []PayModel
}
func (pm *PayManager)CalcPay()  {
	//计算工资,并把工资填充到工资列表
	pm1:=PayModel{
		UserName:"张三",
		Pay:3800,
	}
	pm2:=PayModel{
		UserName:"李四",
		Pay:5800,
	}
	pm3:=PayModel{
		UserName:"王五",
		Pay:6000,
	}
	pm.List=append(pm.List,pm1)
	pm.List=append(pm.List,pm2)
	pm.List=append(pm.List,pm3)
}
func (pm *PayManager)Get(index int) interface{} {
	return pm.List[index]
}
func (pm *PayManager)Size() int {
	return len(pm.List)
}
func (pm *PayManager)CreateIterator() IteratorApi  {
	return &PayIteratorImpl{
		PayM:pm,
		Idx:0,
	}
}

//工资类迭代器
type PayIteratorImpl struct {
	PayM *PayManager
	Idx int
}
func (pi *PayIteratorImpl)First()  {
	pi.Idx=0
}
func (pi *PayIteratorImpl)Next()  {
	if pi.Idx<pi.PayM.Size(){
		pi.Idx=pi.Idx+1
	}
}
func (pi *PayIteratorImpl)IsDone() bool  {
	return pi.Idx>=pi.PayM.Size()
}
func (pi *PayIteratorImpl)CurrentItem() interface{} {
	return pi.PayM.Get(pi.Idx)
}

//工资类过滤式迭代器
type PayFilterIterator struct {
	pms []PayModel
	Idx int
}
func (pf *PayFilterIterator)First()  {
	pf.Idx=0
}
func (pf *PayFilterIterator)Next()  {
	if pf.Idx<len(pf.pms){
		pf.Idx=pf.Idx+1
	}
}
func (pf *PayFilterIterator)IsDone() bool  {
	return pf.Idx>=len(pf.pms)
}
func (pf *PayFilterIterator)CurrentItem() interface{} {
	//我们甚至可以在这里将每个PayModel的Pay置为0
	return pf.pms[pf.Idx]
}
//工资类过滤式迭代器 构造函数
func NewPayFilterIterator(payList []PayModel) IteratorApi  {
	//只查看工资5000以上的人
	pf:=PayFilterIterator{
		pms:make([]PayModel,0),
		Idx:0,
	}
	for _,pm:=range payList{
		if pm.Pay>5000{
			pf.pms=append(pf.pms,pm)
		}
	}
	return &pf
}
