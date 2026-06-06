package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetPokemon(query string) (*Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, query)

	r, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		if r.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("pokemon '%s' not found", query)
		}
		return nil, fmt.Errorf("pokeapi returned unexpected status: %s", r.Status)
	}

	var p Pokemon

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return nil, fmt.Errorf("failed to decode responde: %w", err)
	}

	return &p, nil
}
