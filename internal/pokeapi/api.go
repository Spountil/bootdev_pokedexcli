package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Spountil/bootdev_pokedexcli/internal/pokecache"
)

func FetchApi(conf *Config, url string) ([]byte, error) {
	var data []byte
	result, ok := conf.Cache.CacheMap[url]

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return []byte{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return []byte{}, err
		}
		conf.Cache.CacheMap[url] = pokecache.CacheEntry{
			CreatedAt: time.Now(),
			Val:       data,
		}

		if res.StatusCode > 299 {
			return []byte{}, fmt.Errorf("response's status not 200. Status: %d", res.StatusCode)
		}
	} else {
		data = result.Val
	}

	return data, nil
}

func MapResponse(data []byte) (LocationResponse, error) {

	var locResp LocationResponse

	err := json.Unmarshal(data, &locResp)
	if err != nil {
		return locResp, err
	}

	return locResp, nil
}

func ExploreResponse(data []byte) (LocationAreaDetails, error) {

	var locResp LocationAreaDetails

	err := json.Unmarshal(data, &locResp)
	if err != nil {
		return locResp, err
	}

	return locResp, nil
}

func PokemonResponse(data []byte) (PokemonDetails, error) {

	var locResp PokemonDetails

	err := json.Unmarshal(data, &locResp)
	if err != nil {
		return locResp, err
	}

	return locResp, nil
}
