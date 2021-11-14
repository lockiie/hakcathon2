package models

//Errors é uma estrutura personalizada para ser retornada na api
type ErrorsDB struct {
	// Code    uint16 `json:"code"`
	Message         string `json:"message"`
	InternalMessage string `json:"internal_message"`
	//MessageInternal string `json:"internal_message"`
}

type Errors struct {
	// Code    uint16 `json:"code"`
	Message string `json:"message"`
	//MessageInternal string `json:"internal_message"`
}

//SendError cria uma mensagem de erro
func SendError(message string) Errors {
	return Errors{message}
}

//SendError cria uma mensagem de erro
func SendErrorDB(message string, InternalMessage string) ErrorsDB {
	return ErrorsDB{message, InternalMessage}
}

// func (err *Errors) SentError(c *fiber.Ctx) error {
// 	return
// }
