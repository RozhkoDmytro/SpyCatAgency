package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type Breed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	breedMap = make(map[string]struct{}) // used to store all breeds!
	mu       sync.RWMutex
)

func LoadBreeds() error {
	apiKey := os.Getenv("CAT_API_KEY") // now nnot used
	url := "https://api.thecatapi.com/v1/breeds"

	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to fetch breeds")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var breeds []Breed
	if err := json.Unmarshal(body, &breeds); err != nil {
		return err
	}

	breedMap = make(map[string]struct{})
	for _, b := range breeds {
		breedMap[b.Name] = struct{}{}
	}

	return nil
}

func IsValidBreed(breedName string) bool {
	mu.RLock()
	defer mu.RUnlock()
	_, exists := breedMap[breedName]
	return exists
}
