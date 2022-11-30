package environment

type Args struct {
	Version bool
}

func NewArgs() *Args {
	return &Args{}
}
