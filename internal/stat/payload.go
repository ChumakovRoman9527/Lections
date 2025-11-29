package stat

type GetStatRequest struct {
	FromD string `json:"fromd" validate:"required"`
	ToD   string `json:"tod" validate:"required"`
	By    string `json:"by" validate:"required,oneof=day month"`
}
