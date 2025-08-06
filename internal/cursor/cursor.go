package cursor

type Cursor struct {
	pos int
}

func NewCursor() Cursor {
	return Cursor{
		pos: 0,
	}
}

func (c *Cursor) CurrentPos() int {
	return c.pos
}

func (c *Cursor) Move(pos int) {
	c.pos = pos
}

func (c *Cursor) MoveRight() {
	c.pos += 1
}

func (c *Cursor) MoveLeft() {
	c.pos = max(0, c.pos-1)
}
