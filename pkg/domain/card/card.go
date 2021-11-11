package card

type Card struct {
	id          int
	title       string
	description string
	columnId    int
}

func New(id int, title string, description string, columnId int) *Card {
	return &Card{id: id, title: title, description: description, columnId: columnId}
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

func (c *Card) ColumnId() int {
	return c.columnId
}
