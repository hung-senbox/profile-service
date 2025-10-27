package request

type GenerateOwnerCodeRequest struct {
	OwnerID      string `json:"owner_id" binding:"required"`
	CreatedIndex int    `json:"created_index" binding:"required"`
}
