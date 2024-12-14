package rest_err

import "net/http"

// Objeto para dados contidos nos retornos de requisiçoes
type RestErr struct {
	Message string		`json:"message"`
	Err 		string		`json:"message"`
	Code 		int 			`json:"message"`
	Causes 	[]Causes	`json:"causes,omitempty"`
}


// Objeto para causa de retorno de erro
type Causes struct{
	Field   string   `json:"field"`
	Message string   `json:"message"`
}

// Funçao para retornar apenas a mensagem de erro.
func (r *RestErr) Error() string {
	return r.Message
}

// Criando um construtor
func NewRestErr (message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message:	message,
		Err:			err,
		Code:			code,
		Causes:		causes,
	}
}
//Retorno caso a requisição seja falha
func NewBadResquestError(message string) *RestErr {
	return &RestErr{
		Message:	message,
		Err:			"bad_request",
		Code:			http.StatusBadRequest,
	}
}
//Retorno caso a requisição seja tenha erro de validação
func NewBadResquestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message:	message,
		Err:			"bad_request",
		Code:			http.StatusBadRequest,
		Causes: 	causes,	
	}
}
//Retorno caso a requisição cause erro interno
func NewInternalServerError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message:	message,
		Err:			"internal_server_error",
		Code:			http.StatusInternalServerError,
		Causes: 	causes,	
	}
}
//Retorno caso a requisição nao tenha rota
func NewNotFoundError(message string) *RestErr{
	return &RestErr{
		Message: 	message,
		Err: 			"not_found",
		Code: 		http.StatusNotFound,
	}
}
//Retorno caso a requisição nao seja autorizada
func NewForbiddenError(message string) *RestErr{
	return &RestErr{
		Message:	message,
		Err:			"forbidden",
		Code: 		http.StatusForbidden,
	}
}