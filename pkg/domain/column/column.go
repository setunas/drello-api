package column

type Column struct {
	id      int
	title   string
	boardId int
}

func New(id int, title string, boardId int) *Column {
	return &Column{id: id, title: title, boardId: boardId}
}

func (c *Column) ID() int {
	return c.id
}

func (c *Column) Title() string {
	return c.title
}

func (c *Column) BoardId() int {
	return c.boardId
}
