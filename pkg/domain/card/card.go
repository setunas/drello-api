package card

type Card struct {
	id          int
	title       string
	description string
}

func New(id int, title string, description string) *Card {
	return &Card{id: id, title: title, description: description}
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
