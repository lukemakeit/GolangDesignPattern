package Observer

import "fmt"

//报纸
type NewsPaper struct {
	Readers []Observer //读者
	Content string //报纸内容
}
func (np *NewsPaper)Attach(o Observer)  {
	np.Readers=append(np.Readers,o)
}
func (np *NewsPaper)Notify()  {
	for _,reader:=range np.Readers{
		reader.Update(np)
	}
}
//观察者接口
type Observer interface {
	Update(paper *NewsPaper)
}
type Reader struct {
	Name string //读者姓名
}
func (r *Reader)Update(pager *NewsPaper)  {
	fmt.Printf("%s 收到报纸啦,阅读他,报纸内容是===%s\n",r.Name,pager.Content)
}
