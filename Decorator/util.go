package Decorator

import (
	"sync"
)

var monthSaleMoney map[string]float64
var once sync.Once
//模拟DB中记录每个人的月度销售额
func GetTempDB() map[string]float64  {
	once.Do(func() {
		monthSaleMoney=make(map[string]float64)
		monthSaleMoney["张三"]=10000.0
		monthSaleMoney["李四"]=20000.0
		monthSaleMoney["王五"]=30000.0
	})
	return monthSaleMoney
}
