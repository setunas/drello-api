package user

type User struct {
	id       int
	username string
	boardID  int
}

func New(id int, username string, boardID int) *User {
	return &User{id: id, username: username, boardID: boardID}
}

func (c *User) ID() int {
	return c.id
}

func (c *User) Username() string {
	return c.username
}

func (c *User) BoardID() int {
	return c.boardID
}
