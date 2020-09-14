package user

type RegisterUserReq struct {
	Name     string
	Barthday string
	Sex      string
}

type RegisterUserRes struct {
	Id int64
}
