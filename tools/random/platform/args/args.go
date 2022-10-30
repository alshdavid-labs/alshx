package args

type Args struct {
	Command string
	Args    []string
}

// NewArgs helps with getting the args of nested commands.
// eg: "foo --bar foobar" becomes { Command: "foo", Args "--bar foobar" }
func NewArgs(a []string) *Args {
	command := ""
	args := []string{}

	if len(a) >= 1 {
		command = a[0]
	}
	if len(a) >= 2 {
		args = append(args, a[1:]...)
	}

	return &Args{
		Command: command,
		Args:    args,
	}
}
