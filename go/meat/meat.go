package meat

type Meat interface {
	GetName() string
	GetProcessingSeconds() int
}

type Beef struct{}

func (Beef) GetName() string {
	return "牛肉"
}
func (Beef) GetProcessingSeconds() int {
	return 1
}

type Pork struct{}

func (Pork) GetName() string {
	return "豬肉"
}
func (Pork) GetProcessingSeconds() int {
	return 2
}

type Chicken struct{}

func (Chicken) GetName() string {
	return "雞肉"
}
func (Chicken) GetProcessingSeconds() int {
	return 3
}
