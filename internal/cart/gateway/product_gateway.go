package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type ProductGateway interface {
	CheckProductExists(ctx context.Context, productId int) (int, int, error)
}

type HttpProductGateway struct {
	baseURL string
	client *http.Client
}

type ProductResponse struct {
	ID int `json:"id"`
	Stock int `json:"stock"`
}


func NewHttpProductGateway(baseUrl string, client *http.Client) *HttpProductGateway {
	return &HttpProductGateway{baseURL: baseUrl, client: client}
}

func (h *HttpProductGateway) CheckProductExists(ctx context.Context, productId int) (int, int, error) {
	url := fmt.Sprintf("%s/%d", h.baseURL, productId)

	log.Printf("url: %v", url)


	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return -1, -1, err
	}

	result, err := http.DefaultClient.Do(r)

	if err != nil {
		return -1, -1, err
	}

	var product ProductResponse

	log.Print("after request")


	log.Printf("response: %v", result.StatusCode)
	log.Printf("d: %v", http.StatusNotFound)

	if result.StatusCode == http.StatusNotFound {
		return -1, -1, errors.New("Product not found")
	}
	log.Print("after validating statuscode")

	if result.StatusCode == http.StatusInternalServerError {
		return -1, -1, errors.New("Internal server error")
	}

	log.Print("after validating statuscode 2")


	err = json.NewDecoder(result.Body).Decode(&product)

	log.Printf("after decoding body: %+v", product)

	if err != nil {
		return -1, -1, err
	}

	return product.ID, product.Stock, nil
}
