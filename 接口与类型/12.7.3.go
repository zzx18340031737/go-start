type Phone struct {
	Name string
}

func (p *Phone) Call() error {
	fmt.Println(p.Name,"Start Call")
	return nil
}

func (p *Phone) Video() error {
fmt.Println(p.Name,"Start Video")
return nil
}

func (p *Phone) Game() error {
fmt.Println(p.Name,"Start Game")
return nil
}