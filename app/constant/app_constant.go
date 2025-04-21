package constant

type ResponseStatus int
type Headers int
type General int
type Gender string

// Constant Api
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
	Duplicated
	ValidateError
	BadRequest
	RequiredQuery
	DataIsExit
	MethodNotFound
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED", "DUPLICATED", "ValidateError", "BAD_REQUEST",
		"REQUIRED_QUERY", "DATA_IS_EXIT", "METHOD_NOT_FOUND"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized", "Duplicated", "ValidateError", "BadRequest", "RequiredQuery", "DataIsExit", "Method Not Found"}[r-1]
}

const (
	Male   Gender = "M"
	Female Gender = "F"
)
