type IPhoner interface {
	Call() error
	Video() error
	Game() error
}

type Apple interface{
	Call() error
	Video() error
}

type Huawei interface {
	Call() error
	Game() error
}

