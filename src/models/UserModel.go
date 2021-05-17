package models

type UserModel struct {
	UserId   int  `uri:"id" binding:"required,gt=0"`
	UserName string
}

func (this *UserModel) String() string {
	return "usermodel"
}
