package user

type User struct {
	id          int
	username    string
	boardID     int
	firebaseUID string
}

func New(id int, username string, boardID int, firebaseUID string) *User {
	return &User{id: id, username: username, boardID: boardID, firebaseUID: firebaseUID}
}

func (u *User) ID() int {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) BoardID() int {
	return u.boardID
}

func (u *User) FirebaseUID() string {
	return u.firebaseUID
}
