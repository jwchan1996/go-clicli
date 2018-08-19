package def

type Count struct {
	Pid int `json:"pid"`
	Pv  int `json:"pv,omitempty"`
	Cv  int `json:"cv"`
}
