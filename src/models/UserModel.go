package models

type UserModel struct {
	UserId   int `uri:"id" binding:"required,gt=0"`
	UserName string
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Email    string `db:"email"`
	ID       int64  `db:"id"`
}

func (this *UserModel) String() string {
	return "usermodel"
}
