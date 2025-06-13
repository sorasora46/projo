package dtos

type FailRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type FailValidationRes struct {
	Success bool              `json:"success"`
	Errors  []ValidationError `json:"errors"`
}

type SuccessRes struct {
	Success bool `json:"success"`
	Result  any  `json:"result"`
}

type Response struct {
	Code   int
	Result any
	Error  error
}
