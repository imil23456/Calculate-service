package handler

import (
	"calculator/internal/calculator"
	"calculator/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func HandleCalculation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req models.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := models.CalculationResponse{Error: "Invalid JSON format"}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := calculator.Calc(req.Expression)
	if err != nil {
		response := models.CalculationResponse{Error: err.Error()}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.CalculationResponse{Result: strconv.FormatFloat(result, 'f', -1, 64)}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
