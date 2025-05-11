package wildsapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type QueryOptions struct {
	Page        int
	FilterQuery map[string]any
}

type Client struct {
	BaseUrl      string
	LanguageCode string
	PageSize     int
}

func GetDefaultClient() *Client {
	return &Client{
		BaseUrl:      "https://wilds.mhdb.io/",
		LanguageCode: "en",
		PageSize:     20,
	}
}

func (c *Client) getFetchUrl(path string, query QueryOptions) (*url.URL, error) {
	url, err := url.Parse(c.BaseUrl)
	if err != nil {
		return nil, err
	}
	url = url.JoinPath(c.LanguageCode, path)

	q := url.Query()

	// Pagination
	if c.PageSize > 0 {
		q.Set("limit", strconv.Itoa(c.PageSize))
	}
	if query.Page > 0 {
		offset := (query.Page - 1) * c.PageSize
		q.Set("offset", strconv.Itoa(offset))
	}

	// Filter
	if query.FilterQuery != nil {
		queryJSON, err := json.Marshal(query.FilterQuery)
		if err != nil {
			return nil, err
		}
		q.Set("q", string(queryJSON))
	}

	url.RawQuery = q.Encode()
	return url, nil
}

func (c *Client) getFetchByIdUrl(path string, id int) (*url.URL, error) {
	url, err := url.Parse(c.BaseUrl)
	if err != nil {
		return nil, err
	}

	url = url.JoinPath(c.LanguageCode, path, strconv.Itoa(id))

	return url, nil
}

func (c *Client) fetchById(path string, id int) (*http.Response, error) {
	url, err := c.getFetchByIdUrl(path, id)
	if err != nil {
		return nil, err
	}
	return http.Get(url.String())
}

func (c *Client) fetch(path string, q QueryOptions) (*http.Response, error) {
	url, err := c.getFetchUrl(path, q)
	if err != nil {
		return nil, err
	}

	return http.Get(url.String())
}
