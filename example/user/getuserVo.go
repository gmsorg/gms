package user

type GetUserReq struct {
	Id int64
}

type GetUserRes struct {
	Name     string
	Barthday string
	Sex      string
}
