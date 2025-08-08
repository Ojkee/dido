package command

type None struct{}

func NewNone() *None {
	return &None{}
}

func (c *None) Execute() error { return nil }
func (c *None) Undo() error    { return nil }
