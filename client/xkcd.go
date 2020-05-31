package client

import (
	"encoding/json"
	"fmt"
	"github.com/Comic-Ghost/model"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

type ComicNumber int

// Constants
const (
	BaseURL              string        = "https://xkcd.com"
	DefaultClientTimeout time.Duration = 30 * time.Second
	LatestComic          ComicNumber   = 0
)

// XKCD client
type XKCDClient struct {
	client  *http.Client
	baseURL string
}

// Creates new XKCD client
func NewXKCDClient() *XKCDClient {
	return &XKCDClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL: BaseURL,
	}
}

// Adding four methods for XKCD client: SetTimeout, Fetch, SaveToDisk, BuildURL
func (x *XKCDClient) SetTimeout(timeout time.Duration) {
	x.client.Timeout = timeout
}

func (x *XKCDClient) Fetch(n ComicNumber, save bool) (model.Comic, error) {
	response, err := x.client.Get(x.BuildURL(n))
	if err != nil {
		return model.Comic{}, err
	}

	defer response.Body.Close()

	var comicResponse model.ComicResponse
	if err := json.NewDecoder(response.Body).Decode(&comicResponse); err != nil {
		return model.Comic{}, err
	}

	if save {
		if err := x.SaveToDisk(comicResponse.Img, "."); err != nil {
			fmt.Println("Failed to save image!")
		}
	}

	return comicResponse.Comic(), nil
}

func (x *XKCDClient) SaveToDisk(url string, savePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	absSavePath, _ := filepath.Abs(savePath)
	filePath := fmt.Sprintf("%s/%s", absSavePath, path.Base(url))

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	// No errors
	return nil
}

func (x *XKCDClient) BuildURL(n ComicNumber) string {
	var finalURL string

	if n == LatestComic {
		finalURL = fmt.Sprintf("%s/info.0.json", x.baseURL)
	} else {
		finalURL = fmt.Sprintf("%s/%d/info.0.json", x.baseURL, n)
	}

	return finalURL
}
