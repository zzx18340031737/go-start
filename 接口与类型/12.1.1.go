type Anunal interface {
	Name() string
	Speak() string
}

type Cat struct {

}

func(cat Cat)Name() string {
	return "Cat"
}

func (cat Cat)Speak() string  {
	return "喵喵喵"
}
