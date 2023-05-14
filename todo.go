package simpleRestAPI

type TodoList struct {
	Id    int    `json:"-"`
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}
