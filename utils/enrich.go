package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RandomUserResponse struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Location struct {
			Street struct {
				Number int    `json:"number"`
				Name   string `json:"name"`
			} `json:"street"`
			City        string `json:"city"`
			State       string `json:"state"`
			Country     string `json:"country"`
			Postcode    int    `json:"postcode"`
			Coordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"coordinates"`
			Timezone struct {
				Offset      string `json:"offset"`
				Description string `json:"description"`
			} `json:"timezone"`
		} `json:"location"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Cell    string `json:"cell"`
		Picture struct {
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Thumbnail string `json:"thumbnail"`
		} `json:"picture"`
	} `json:"results"`
}

func FetchRandomUser(url string) (string, string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}

	var data RandomUserResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", "", "", err
	}

	randomUser := data.Results[0]

	return randomUser.Name.First,
		randomUser.Name.Last,
		fmt.Sprintf("%d %s, %s", randomUser.Location.Street.Number, randomUser.Location.Street.Name, randomUser.Location.State),
		nil
}
