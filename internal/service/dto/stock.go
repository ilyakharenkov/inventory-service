package dto

type Stock struct {
	ActionType Action `json:"actionType"`
	Quantity   int64  `json:"quantity"`
}

type Action string

const (
	Add      Action = "add"
	Subtract Action = "subtract"
)
