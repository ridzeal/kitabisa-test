package entity

// BundleRequest is parameter required during bundle request
type BundleRequest struct {
	Cakes int `json:"cakes"`
	Apples int `json:"apples"`
}

// BundleResponse returns when bundle API called
type BundleResponse struct {
	TotalBoxes int `json:"total_boxes"`
	Cakes int `json:"cakes"`
	Apples int `json:"apples"`
	ErrorMessage string `json:"error_message"`
}