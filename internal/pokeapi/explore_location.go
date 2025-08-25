package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Explore(locationArea string) (Encounters, error) {
	url := baseURL + "/location-area/" + locationArea

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Encounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Encounters{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Encounters{}, err
	}

	encounters := Encounters{}
	err = json.Unmarshal(data, &encounters)
	if err != nil {
		return Encounters{}, err
	}

	return encounters, nil
}
