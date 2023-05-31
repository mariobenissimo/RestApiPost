package models

type Response map[string]interface{}

type ResponseLogin struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Token   string `json:"token"`
	User    User   `json:"user"`
}
