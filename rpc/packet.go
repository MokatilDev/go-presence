package rpc

type HandShake struct {
	V        string `json:"v"`
	ClientID string `json:"client_id"`
}

type Packet struct {
	Nonce string `json:"nonce"`
	Cmd   string `json:"cmd"`
	Args  Args   `json:"args"`
}

type Args struct {
	Pid      int       `json:"pid"`
	Activity *Activity `json:"activity,omitempty"`
}
