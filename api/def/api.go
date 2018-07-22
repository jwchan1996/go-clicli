package def

type UserCredential struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pwd  string `json:"pwd,omitempty"`
	Role string `json:"role"`
	QQ   int    `json:"qq"`
	Desc string `json:"desc"`
}

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Sort    string `json:"sort"`
	Time    string `json:"time"`
}
