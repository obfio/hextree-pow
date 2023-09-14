package hextree

import (
	"encoding/json"
	"io"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

// GetConfig gets the POW config
func (c *Client) GetConfig() (*Config, error) {
	req, err := http.NewRequest("GET", "https://app.hextree.io/api/account/reset", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("Authorization", "undefined")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	f := &Config{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// TestHeader tests the generated X-POW header
func (c *Client) TestHeader(POW string) (string, error) {
	req, err := http.NewRequest("POST", "https://app.hextree.io/api/account/reset", strings.NewReader("{\"email\":\"admin@antibot.blog\"}"))
	if err != nil {
		return "", err
	}
	req.Header.Set("X-POW", POW)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
