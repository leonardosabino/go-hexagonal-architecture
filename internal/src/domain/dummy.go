package domain

type Dummy struct {
	Id          *string `json:"id"`
	Description *string `json:"description" validate:"required"`
}
