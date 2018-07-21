package def

type UserCredential struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	Role string `json:"role"`
	QQ   int    `json:"qq"`
	Desc string `json:"desc"`
}

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Sort    string `json:"sort"`
	Time    string `json:"time"`
}
