package pokesdk

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/"

// client: used to perform REST requests to PokeAPI
type client struct {
	httpClient *http.Client
}

// newClient: creates a new REST client
func newClient() *client {
	return &client{
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

// setTimeout: sets the timeout on the REST client
func (c *client) setTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

// getPokemon: gets pokenmon details using the given name or ID
func (c *client) getPokemon(parameter string) (*apiPokemon, error) {
	var pokemon apiPokemon

	err := c.callEndpoint("pokemon", parameter, &pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

// getNature: gets nature details using the given name or ID
func (c *client) getNature(parameter string) (*apiNature, error) {
	var nature apiNature

	err := c.callEndpoint("nature", parameter, &nature)
	if err != nil {
		return nil, err
	}

	return &nature, nil
}

// getStat: gets stat details using the given name or ID
func (c *client) getStat(parameter string) (*apiStat, error) {
	var stat apiStat

	err := c.callEndpoint("stat", parameter, &stat)
	if err != nil {
		return nil, err
	}

	return &stat, nil
}

// callEndpoint: makes external calls to the given path, with the given paramter
// attempts to marshall the returned payload into the given target
func (c *client) callEndpoint(path, parameter string, target interface{}) error {
	// Call the endpoint
	resp, err := c.httpClient.Get(baseUrl + path + "/" + parameter)
	if err != nil {
		return ErrReadFailed
	}

	// Check the response status
	switch resp.StatusCode {
	case 404:
		return ErrNotFound
		// TODO: add more status codes

	}

	// Read the body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ErrReadFailed
	}

	// Marshall to target struct
	err = json.Unmarshal(body, target)
	if err != nil {
		return ErrUnexpectedResponse
	}
	return nil
}
