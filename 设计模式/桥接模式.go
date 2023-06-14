package 设计模式

import "fmt"

// ICoffee 咖啡接口
type ICoffee interface {
	// OrderCoffee 订购咖啡
	OrderCoffee()
}

// LargeCoffee 大杯咖啡，实现ICoffee接口
type LargeCoffee struct {
	ICoffeeAddition
}

// MediumCoffee 中杯咖啡，实现ICoffee接口
type MediumCoffee struct {
	ICoffeeAddition
}

// SmallCoffee 小杯咖啡，实现ICoffee接口
type SmallCoffee struct {
	ICoffeeAddition
}

// OrderCoffee 订购大杯咖啡
func (lc LargeCoffee) OrderCoffee() {
	fmt.Println("订购了大杯咖啡")
	lc.AddSomething()
}

// OrderCoffee 订购中杯咖啡
func (mc MediumCoffee) OrderCoffee() {
	fmt.Println("订购了中杯咖啡")
	mc.AddSomething()
}

// OrderCoffee 订购小杯咖啡
func (sc SmallCoffee) OrderCoffee() {
	fmt.Println("订购了小杯咖啡")
	sc.AddSomething()
}

// CoffeeCupType 咖啡容量类型
type CoffeeCupType uint8

const (
	// CoffeeCupTypeLarge 大杯咖啡
	CoffeeCupTypeLarge = iota
	// CoffeeCupTypeMedium 中杯咖啡
	CoffeeCupTypeMedium = iota
	// CoffeeCupTypeSmall 小杯咖啡
	CoffeeCupTypeSmall = iota
)

// CoffeeFuncMap 全局可导出变量，咖啡类型与创建咖啡对象的map，用于减小圈复杂度
var CoffeeFuncMap = map[CoffeeCupType]func(coffeeAddition ICoffeeAddition) ICoffee{
	CoffeeCupTypeLarge:  NewLargeCoffee,
	CoffeeCupTypeMedium: NewMediumCoffee,
	CoffeeCupTypeSmall:  NewSmallCoffee,
}

// NewCoffee 创建咖啡接口对象的简单工厂，根据咖啡容量类型，获取创建接口对象的func
func NewCoffee(cupType CoffeeCupType, coffeeAddition ICoffeeAddition) ICoffee {
	if handler, ok := CoffeeFuncMap[cupType]; ok {
		return handler(coffeeAddition)
	}
	return nil
}

// NewLargeCoffee 创建大杯咖啡对象
func NewLargeCoffee(coffeeAddition ICoffeeAddition) ICoffee {
	return &LargeCoffee{coffeeAddition}
}

// NewMediumCoffee 创建中杯咖啡对象
func NewMediumCoffee(coffeeAddition ICoffeeAddition) ICoffee {
	return &MediumCoffee{coffeeAddition}
}

// NewSmallCoffee 创建小杯咖啡对象
func NewSmallCoffee(coffeeAddition ICoffeeAddition) ICoffee {
	return &SmallCoffee{coffeeAddition}
}

// ICoffeeAddition 咖啡额外添加接口
type ICoffeeAddition interface {
	//AddSomething 添加某种食物
	AddSomething()
}

// Milk 加奶，实现ICoffeeAddition接口
type Milk struct{}

// Sugar 加糖，实现ICoffeeAddition接口
type Sugar struct{}

// AddSomething Milk实现加奶
func (milk Milk) AddSomething() {
	fmt.Println("加奶")
}

// AddSomething Sugar实现加糖
func (sugar Sugar) AddSomething() {
	fmt.Println("加糖")
}

// CoffeeAdditionType 咖啡额外添加类型
type CoffeeAdditionType uint8

const (
	// CoffeeAdditionTypeMilk 咖啡额外添加牛奶
	CoffeeAdditionTypeMilk = iota
	// CoffeeAdditionTypeSugar 咖啡额外添加糖
	CoffeeAdditionTypeSugar = iota
)

// CoffeeAdditionFuncMap 全局可导出变量，咖啡额外添加类型与创建咖啡额外添加对象的map，用于减小圈复杂度
var CoffeeAdditionFuncMap = map[CoffeeAdditionType]func() ICoffeeAddition{
	CoffeeAdditionTypeMilk:  NewCoffeeAdditionMilk,
	CoffeeAdditionTypeSugar: NewCoffeeAdditionSugar,
}

// NewCoffeeAddition 创建咖啡额外添加接口对象的简单工厂，根据咖啡额外添加类型，获取创建接口对象的func
func NewCoffeeAddition(addtionType CoffeeAdditionType) ICoffeeAddition {
	if handler, ok := CoffeeAdditionFuncMap[addtionType]; ok {
		return handler()
	}
	return nil
}

// NewCoffeeAdditionMilk 创建咖啡额外加奶
func NewCoffeeAdditionMilk() ICoffeeAddition {
	return &Milk{}
}

// NewCoffeeAdditionSugar 创建咖啡额外加糖
func NewCoffeeAdditionSugar() ICoffeeAddition {
	return &Sugar{}
}
