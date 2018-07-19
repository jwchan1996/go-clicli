package def

type UserCredential struct {
	UserName string `json:"name"`
	Pwd string `json:"pwd"`
	Role string `json:"role"`
	QQ string `json:"qq"`
	Desc string `json:"desc"`
}

type Post struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Status string `json:"status"`
	Sort string `json:"sort"`
	Time int `json:"time"`
}