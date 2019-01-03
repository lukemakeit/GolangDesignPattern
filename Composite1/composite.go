package Composite1

import "fmt"

type MenuComponent interface {
	Name() string
	Description() string
	Price() float32
	Print()

	Add(component MenuComponent)
	Remove(int)
	Child(int) MenuComponent
}
//MenuItem Menu相同的属性抽象出来
type MenuDesc struct {
	name string
	description string
}
func (md *MenuDesc)Name() string {
	return md.name
}
func (md *MenuDesc)Description() string  {
	return md.description
}
//菜单项
type MenuItem struct {
	MenuDesc
	price float32
}
func (mi *MenuItem)Price() float32  {
	return mi.price
}
func (mi *MenuItem)Print()  {
	fmt.Printf("  %s,￥%.2f\n",mi.name,mi.price)
	fmt.Printf("    -- %s\n",mi.description)
}
func (mi *MenuItem)Add(component MenuComponent)  {
	panic("not implement")
}
func (mi *MenuItem)Remove(idx int)  {
	panic("not implement")
}
func (mi *MenuItem)Child(idx int) MenuComponent  {
	panic("not implement")
}

func NewMenuItem(name,description string,price float32) MenuComponent  {
	return &MenuItem{
		MenuDesc{
			name:name,
			description:description,
		},
		price,
	}
}

//菜单
type Menu struct {
	MenuDesc
	children []MenuComponent
}
func (m *Menu)Price()(price float32)  {
	for _,v:=range m.children{
		price+=v.Price()
	}
	return
}
func (m *Menu)Print()  {
	fmt.Printf("%s,%s, ￥%.2f\n",m.name,m.description,m.Price())
	fmt.Println("----------------------------")
	for _,v:=range m.children{
		v.Print()
	}
	fmt.Println()
}
func (m *Menu)Add(c MenuComponent)  {
	m.children=append(m.children,c)
}
func (m *Menu)Remove(idx int)  {
	m.children=append(m.children[:idx],m.children[idx+1:]...)
}
func (m *Menu)Child(idx int) MenuComponent  {
	return m.children[idx]
}
func NewMenu(name,description string) MenuComponent {
	return &Menu{
		MenuDesc{
			name: name,
			description:description,
		},
		[]MenuComponent{},
	}
}