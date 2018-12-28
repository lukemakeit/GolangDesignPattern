package Factory

import "fmt"

/*
* 包子店的接口
* 包子店的接口可以产生多个包子店类
*/
type BunShopInterface interface {
	Generate(t string) Bun
}
/*
* 包子的接口
* 包子的接口可以产生多种包子类
*/
type Bun interface {
	Eat()
}

//北京包子店
//当然程序也可以让北京包子店卖广东味的包子
type BJBunShop struct {}
func (bj BJBunShop)Generate(t string) Bun {
	if t == "pig"{
		return BJPigMeatBuns{}
	}else if(t == "sun"){
		return BJSunStuffingBuns{}
	}
	return nil
}
//广东包子店
type GDBunShop struct {}
func (bj GDBunShop)Generate(t string) Bun {
	if t == "pig"{
		return GDPigMeatBuns{}
	}else if(t == "sun"){
		return GDSunStuffingBuns{}
	}
	return nil
}

//北京猪肉馅包子
type BJPigMeatBuns struct {}
func (bjpm BJPigMeatBuns)Eat(){
	fmt.Println("老板,一份北京猪肉馅的包子")
}
//北京三鲜馅包子
type BJSunStuffingBuns struct {}
func (bjsss BJSunStuffingBuns)Eat(){
	fmt.Println("老板,一份北京三鲜馅的包子")
}

//广东猪肉馅包子
type GDPigMeatBuns struct {}
func (gdpb GDPigMeatBuns)Eat(){
	fmt.Println("老板,一份广东猪肉馅的包子")
}
//广东三鲜馅包子
type GDSunStuffingBuns struct {}
func (gdss GDSunStuffingBuns)Eat(){
	fmt.Println("老板,一份广东三鲜馅的包子")
}