package models

type DeleteCartItemResponse struct {
	 Status  bool `json:"status,omitempty"`
	 Err            error  `json:"error,omitempty"`
}