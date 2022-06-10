package contract

type VMCallData struct {
	From string `json:"from"`
	To   string `json:"to"`
	Data string `json:"data"`
}
