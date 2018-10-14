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
	Type    string `json:"type,omitempty"`
	Time    string `json:"time"`
	Uid     int    `json:"uid,omitempty"`
	Uname   string `json:"uname,omitempty"`
	Uqq     string `json:"uqq,omitempty"`
	Count   *Count `json:"count,omitempty"`
}

type Comment struct {
	Id      int    `json:"id,omitempty"`
	Content string `json:"content"`
	Ctime   string `json:"ctime"`
	Pid     int    `json:"pid"`
	Vid     int    `json:"vid,omitempty"`
	Tuid    int    `json:"tuid,omitempty"`
	Time    int `json:"time,omitempty"`
	Uid     int    `json:"uid,omitempty"`
	Uname   string `json:"uname,omitempty"`
	Uqq     string `json:"uqq,omitempty"`
}

type Video struct {
	Id      int    `json:"id,omitempty"`
	Oid     int    `json:"oid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Pid     int    `json:"pid"`
	Ptitle  string `json:"ptitle,omitempty"`
	Uid     int    `json:"uid"`
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
	Comments []*Comment `json:"comments"`
}

type Videos struct {
	Videos []*Video `json:"videos"`
}

type Token struct {
	Name string
	Pwd  string
}
