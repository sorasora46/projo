package dtos

type ValidationError struct {
	Field         string `json:"field"`
	Message       string `json:"message"`
	ValueProvided any    `json:"valueProvided"`
}
