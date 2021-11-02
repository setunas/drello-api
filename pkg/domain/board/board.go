package board

type Board struct {
	id    int
	title string
}

func New(id int, title string) *Board {
	return &Board{id: id, title: title}
}

func (b *Board) ID() int {
	return b.id
}

func (b *Board) Title() string {
	return b.title
}
