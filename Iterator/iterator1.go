package Iterator

type IteratorApi interface {
	First()
	Next()
	IsDone() bool
	CurrentItem() interface{}
}
//聚合类,如slice map 甚至二叉树等
type ConcreteAggregate struct {
	List []string
}
func (ca *ConcreteAggregate)Get(idx int) interface{} {
	return ca.List[idx]
}
func (ca *ConcreteAggregate)Size() int  {
	return len(ca.List)
}
func (ca *ConcreteAggregate) CreateIterator()IteratorApi{
	return &ConcreateIterator{
		ConcreteAgg:ca,
		Idx:0,
	}
}
//具体迭代器
type ConcreateIterator struct {
	Idx int
	ConcreteAgg *ConcreteAggregate
}

func (ci *ConcreateIterator)First()  {
	ci.Idx=0
}
func (ci *ConcreateIterator)Next()  {
	if ci.Idx < ci.ConcreteAgg.Size() {
		ci.Idx=ci.Idx+1
	}
}
func (ci *ConcreateIterator)IsDone() bool {
	return ci.Idx>=ci.ConcreteAgg.Size()
}
func (ci *ConcreateIterator)CurrentItem() interface{}  {
	return ci.ConcreteAgg.Get(ci.Idx)
}
