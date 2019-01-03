package Composite2

import "fmt"
//只抽象了Price()方法和Print()方法,Add() Remove() Child()等只在Menu中实现
type MenuComponent interface {
	Price() float32
	Print()
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
//MenuItem 并没有实现Add() Remove() Child()这三种方法
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
//NewMenuItem 返回的是一个*MenuItem 而不是一个MenuComponent
func NewMenuItem(name,description string,price float32) *MenuItem {
	return &MenuItem{
		MenuDesc{
			name:name,
			description:description,
		},
		price,
	}
}

type MenuGroup struct {
	children []MenuComponent
}
//菜单
//菜单实现Add() Remove() Child()等方法
type Menu struct {
	MenuDesc
	MenuGroup
}
func (m *Menu) Add(c MenuComponent) {
	m.children = append(m.children, c)
}
func (m *Menu) Remove(idx int) {
	m.children = append(m.children[:idx], m.children[idx+1:]...)
}
func (m *Menu) Child(idx int) MenuComponent {
	return m.children[idx]
}
func (m *Menu) Price() (price float32) {
	for _, v := range m.children {
		price += v.Price()
	}
	return
}
func (m *Menu) Print() {
	fmt.Printf("%s, %s, ￥%.2f\n", m.name, m.description, m.Price())
	fmt.Println("------------------------")
	for _, v := range m.children {
		v.Print()
	}
	fmt.Println()
}
//NewMenu 返回的是一个*Menu 而不是一个MenuComponent
func NewMenu(name, description string) *Menu {
	return &Menu{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
	}
}
//从Menu中读取某个child,其可用方法只有Price()和Price(),一样可以安全调用,如果想在MenuComponent是Menu的情况下添加子项呢?
//if m, ok := all.Child(1).(*Menu); ok {
//	m.Add(NewMenuItem("玩具", "Hello Kitty", 5.0))
//}