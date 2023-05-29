package simpleRestAPI

type TodoList struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Desc  string `json:"desc" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}
