package request

type MoveRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
