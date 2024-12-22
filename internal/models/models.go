package models

type CalculationRequest struct {
	Expression string `json:"expression"`
}

type CalculationResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}
