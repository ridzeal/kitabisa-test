package handler

import (
	"net/http"
	"fmt"
	"kitabisa-test/entity"
	"kitabisa-test/controller"
)

// Bundle handles request to bundle cakes and apples
func Bundle(w http.ResponseWriter, r *http.Request) {
	var bundleRequest *entity.BundleRequest

	if err := parseBody(r.Body, r.Header, &bundleRequest); err != nil {
		fmt.Printf("Request body parsing error: %v", err.Error())
		
		respBody := &entity.BundleResponse{
			TotalBoxes: 0,
			Cakes: 0,
			Apples: 0,
			ErrorMessage: "Missing parameter",
		}

		renderResponse(w, r, http.StatusBadRequest, respBody)
		return
	}

	// Zero validation
	if bundleRequest.Apples <= 0 || bundleRequest.Cakes <= 0 {
		respBody := &entity.BundleResponse{
			TotalBoxes: 0,
			Cakes: 0,
			Apples: 0,
			ErrorMessage: "Zero quantity",
		}

		renderResponse(w, r, http.StatusBadRequest, respBody)
		return
	}

	total, apple, cake := controller.BundleIt(bundleRequest.Apples, bundleRequest.Cakes)

	bundleResponse := entity.BundleResponse{
		TotalBoxes: total,
		Cakes: cake,
		Apples: apple,
	}
	renderResponse(w, r, http.StatusOK, bundleResponse)
}