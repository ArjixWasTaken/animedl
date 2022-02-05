package commands

type Command struct {
	Name        string
	Description string
	Callback    func()
}

var Commands = []string{}
