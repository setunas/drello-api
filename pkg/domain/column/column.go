package column

type Column struct {
	id       int
	title    string
	position float64
	boardId  int
}

func New(id int, title string, position float64, boardId int) *Column {
	return &Column{id: id, title: title, position: position, boardId: boardId}
}

func (c *Column) ID() int {
	return c.id
}

func (c *Column) Title() string {
	return c.title
}

func (c *Column) Positon() float64 {
	return c.position
}

func (c *Column) BoardId() int {
	return c.boardId
}
