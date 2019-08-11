package def

type Count struct {
	Pid int `json:"pid"`
	Pv  int `json:"pv,omitempty"`
	Cv  int `json:"cv"`
}

type Cookie struct {
	Uid  int    `json:"uid"`
	Hcy  string `json:"hcy"`
	Quqi string `json:"quqi,omitempty"`
}
