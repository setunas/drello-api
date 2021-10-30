package board

type Board struct {
	id    int
	title string
}

func New(id int, title string) *Board {
	return &Board{id: id, title: title}
}

func (w *Board) ID() int {
	return w.id
}

func (w *Board) Title() string {
	return w.title
}
