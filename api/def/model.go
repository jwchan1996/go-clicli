package def

type UserCredential struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Pwd  string `json:"pwd,omitempty"`
	Role string `json:"role"`
	QQ   int    `json:"qq"`
	Desc string `json:"desc"`
}

type Post struct {
	Id      int    `json:"id,omitempty"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Sort    string `json:"sort"`
	Time    string `json:"time"`
	Uid     int    `json:"uid,omitempty"`
	Uname   string `json:"uname,omitempty"`
	Uqq     string `json:"uqq,omitempty"`
}

type Comment struct {
	Id      int    `json:"id,omitempty"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Pid     int    `json:"pid"`
	Uid     int    `json:"uid,omitempty"`
	Uname   string `json:"uname,omitempty"`
	Uqq     string `json:"uqq,omitempty"`
}

type Posts struct {
	Posts []*Post `json:"posts"`
}

type Users struct {
	Users []*UserCredential `json:"users"`
}

type Comments struct {
	Posts []*Comment `json:"comments"`
}
