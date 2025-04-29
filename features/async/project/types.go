package project

type ProjectEvent struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
