package req

type CreateProjectReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
