package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HttpAdapter struct {
	Client  *http.Client
	Headers map[string]string
	BaseUrl string
}

func NewHttpAdapter(baseUrl string) *HttpAdapter {
	client := &http.Client{}

	return &HttpAdapter{
		Client:  client,
		Headers: make(map[string]string),
		BaseUrl: baseUrl,
	}
}

func (r *HttpAdapter) AddHeader(key string, value string) {
	r.Headers[key] = value
}

func (r *HttpAdapter) Get(url string) ([]byte, error) {
	urlComplete := r.BaseUrl + url

	req, err := http.NewRequest("GET", urlComplete, nil)
	if err != nil {
		return nil, errors.New("Failed to create request: " + err.Error())
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, errors.New("Failed to send request: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("GET request failed, status code: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	responseBody, err := r.convertBodyToByte(resp.Body)

	return responseBody, err
}

func (r *HttpAdapter) Post(url string, body io.Reader) ([]byte, error) {
	urlComplete := r.BaseUrl + url

	req, err := http.NewRequest("POST", urlComplete, body)
	if err != nil {
		return nil, errors.New("Failed to create request: " + err.Error())
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, errors.New("Failed to send request: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("POST request failed, status code: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	responseBody, err := r.convertBodyToByte(resp.Body)

	return responseBody, err
}

func (r *HttpAdapter) Put(url string, body io.Reader) ([]byte, error) {
	urlComplete := r.BaseUrl + url

	req, err := http.NewRequest("PUT", urlComplete, body)
	if err != nil {
		return nil, errors.New("Failed to create request: " + err.Error())
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, errors.New("Failed to send request: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("PUT request failed, status code: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	responseBody, err := r.convertBodyToByte(resp.Body)

	return responseBody, err
}

func (r *HttpAdapter) Delete(url string) ([]byte, error) {
	urlComplete := r.BaseUrl + url

	req, err := http.NewRequest("DELETE", urlComplete, nil)
	if err != nil {
		return nil, errors.New("Failed to create request: " + err.Error())
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, errors.New("Failed to send request: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("DELETE request failed, status code: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	responseBody, err := r.convertBodyToByte(resp.Body)

	return responseBody, err
}

func (r *HttpAdapter) convertBodyToByte(body io.ReadCloser) ([]byte, error) {
	responseBody, err := io.ReadAll(body)
	if err != nil {
		return nil, errors.New("Failed to read response body: " + err.Error())
	}

	return responseBody, nil
}
