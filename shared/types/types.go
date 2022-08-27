package types

// OrderServiceResponse model info
// @Description types.OrderServiceResponse is a response model for the endpoint of a client
type OrderServiceResponse struct {
	Status string `json:"status" validate:"required"`
}

// SuccessResponseModel model info
// @Description types.SuccessResponseModel is a success response model for PUT&POST ops
type SuccessResponseModel struct {
	ID string `json:"id"`
}

func GetSRM(id string) *SuccessResponseModel {
	srm := &SuccessResponseModel{
		ID: id,
	}
	return srm
}
