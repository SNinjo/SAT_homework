package meat

type Meat interface {
	GetName() string
	GetProcessingSeconds() int
}

type Beef struct{}

func (_ Beef) GetName() string {
	return "牛肉"
}
func (_ Beef) GetProcessingSeconds() int {
	return 1
}

type Pork struct{}

func (_ Pork) GetName() string {
	return "豬肉"
}
func (_ Pork) GetProcessingSeconds() int {
	return 2
}

type Chicken struct{}

func (_ Chicken) GetName() string {
	return "雞肉"
}
func (_ Chicken) GetProcessingSeconds() int {
	return 3
}
