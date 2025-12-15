package gateway

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type UserGateway interface {
	CheckUserExists(ctx context.Context, userId int) error
}

type HttpUserGateway struct {
	baseUrl string
	client *http.Client
}

func NewHttpUserGateway(url string, c *http.Client) *HttpUserGateway {
	return &HttpUserGateway{baseUrl: url, client: c}
}

func (u *HttpUserGateway) CheckUserExists(ctx context.Context, userId int) error {
	log.Println("asdadasdsa")
	url := fmt.Sprintf("%s/exists/%d", u.baseUrl, userId)

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)


	if err != nil {
		return err
	}

	result, err := http.DefaultClient.Do(r)

	if err != nil {
		return err
	}

	if result.StatusCode == http.StatusNotFound {
		return errors.New("User not found")
	}

	if result.StatusCode == http.StatusInternalServerError {
		return errors.New("Something went wrong")
	}

	return nil
}
