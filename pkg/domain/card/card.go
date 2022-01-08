package card

type Card struct {
	id          int
	title       string
	description string
	position    float64
	columnId    int
}

func New(id int, title string, description string, position float64, columnId int) *Card {
	return &Card{id: id, title: title, description: description, position: position, columnId: columnId}
}

func (c *Card) ID() int {
	return c.id
}

func (c *Card) Title() string {
	return c.title
}

func (c *Card) Description() string {
	return c.description
}

func (c *Card) Position() float64 {
	return c.position
}

func (c *Card) ColumnId() int {
	return c.columnId
}
