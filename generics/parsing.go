package generics

import (
	"encoding/json"
	"net/http"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type apiResponse[T any] struct {
	Data   []T `json:"data"`
	Paging struct {
		Total int `json:"total"`
		Limit int `json:"limit"`
	} `json:"paging"`
}

func getApiData[T any](url string) ([]T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var result apiResponse[T]
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
