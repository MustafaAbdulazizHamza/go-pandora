package goPandora

type PandoraClient struct {
	url        string
	username   string
	password   string
	privateKey string
	publicKey  string
}
type user struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Secret struct {
	SecretID string `json:"secret_id" binding:"required"`
	Secret   string `json:"secret" binding:"required"`
}
type Response struct {
	Status string `json:"status" binding:"required"`
	Text   string `json:"text" binding:"required"`
}
type RequestedSecret struct {
	SecretID string `json:"secret_id" binding:"required"`
}
