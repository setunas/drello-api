package column

type Column struct {
	id    int
	title string
}

func New(id int, title string) *Column {
	return &Column{id: id, title: title}
}

func (c *Column) ID() int {
	return c.id
}

func (c *Column) Title() string {
	return c.title
}
