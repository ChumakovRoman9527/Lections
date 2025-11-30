package stat

type GetStatRequest struct {
	FromD string `json:"fromd" validate:"required"`
	ToD   string `json:"tod" validate:"required"`
	By    string `json:"by" validate:"required,oneof=day month"`
}

type GetStatResponce struct {
	Period string `json:"period"`
	Sum    int    `json:"summa"`
}
