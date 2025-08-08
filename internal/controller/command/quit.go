package command

type Quit struct{}

func NewQuit() *Quit {
	return &Quit{}
}

func (c *Quit) Execute() error { return nil }
func (c *Quit) Undo() error    { return nil }
