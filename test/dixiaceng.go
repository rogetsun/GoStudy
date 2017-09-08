//食物木材换算计算
package main

import "fmt"

type makeInterface interface {
	make(person int) (resultFood, resultWood int)
}
type res struct {
	name  string
	price float32
	id    int
}

var (
	//食物单价
	food = res{name: "food", price: 45.0 / 1000.0, id: 1}
	//木材单价,生产方法
	wood = res{name: "wood", price: 140.0 / 1000.0, id: 2}
)

/**
生产模型
*/
type mode struct {
	p2food, p2wood, hour int
}

func main() {
	var ms = make([]mode, 0, 10)
	fmt.Println(ms)
	person := 200
	for i := 0; i < 200; i += 10 {
		ms = append(ms, mode{p2food: i, p2wood: person - i, hour: 1})
	}
	fmt.Println(ms)
	for _, m := range ms {
		produce(m)
	}

}

func produce(m mode) int {

	personForFood := m.p2food
	personForWood := m.p2wood
	person := personForFood + personForWood

	seconds := m.hour * 60 * 60
	pc := seconds / 30
	//fmt.Println("produce", pc, " times")

	var initFood, initWood = 0, 0
	for i := 0; i < pc; i++ {
		tFood, tWood := makeFood(personForFood)
		initFood += tFood
		initWood += tWood
		tFood, tWood = makeWood(personForWood)
		initFood += tFood
		initWood += tWood
		//fmt.Printf("%v[food:%v,wood:%v]\n", (i + 1) * 30, initFood, initWood)
	}
	foodValue := int(float32(initFood) * food.price)
	woodValue := int(float32(initWood) * wood.price)
	value := foodValue + woodValue
	fmt.Printf("%v人生产 %v 小时(%v人食物,%v人木材):\n"+
		"食物产出:%-10v\t  价值:%v\t\n"+
		"木材产出:%-10v\t  价值:%v\t\n"+
		"总价值:%v金\n\n",
		person, m.hour, personForFood, personForWood,
		initFood, foodValue,
		initWood, woodValue,
		value)

	return value
}
func makeFood(person int) (resultFood, resultWood int) {
	resultFood = person * 1
	resultWood = 0
	return
}
func makeWood(person int) (resultFood, resultWood int) {
	resultFood = person * 2 * -1
	resultWood = person * 1
	return
}
