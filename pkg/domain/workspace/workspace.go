package workspace

type Workspace struct {
	id    int
	title string
}

func New(id int, title string) *Workspace {
	return &Workspace{id: id, title: title}
}

func (w *Workspace) Title() string {
	return w.title
}
