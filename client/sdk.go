package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const stagingServer = "https://caregiver-gateway.staging.affinity-project.org/api"

// SDK provides a simple interface to access all the functionality
// provided by the Affinity gateway service.
type SDK struct {
	c           *http.Client
	userAgent   string
	apiEndpoint string

	// Service handlers
	DID *didService
	VC  *vcService
}

// Options available when configuring an SDK client instance.
type Options struct {
	// Time to wait for requests, in seconds
	Timeout uint

	// Time to maintain open the connection with the service, in seconds
	KeepAlive uint

	// Maximum network connections to keep open with the service
	MaxConnections uint

	// User agent value to report to the service
	UserAgent string
}

// Return sane default configuration values
func defaultOptions() *Options {
	return &Options{
		Timeout:        30,
		KeepAlive:      600,
		MaxConnections: 100,
		UserAgent:      "affinityctl/0.1.0",
	}
}

// New SDK instance to access the Affinity gateway service. If no
// configuration options are provided (i.e., nil), sane default
// values are used.
func New(opts *Options) (*SDK, error) {
	// Default settings
	if opts == nil {
		opts = defaultOptions()
	}

	// Configure base HTTP transport
	t := &http.Transport{
		MaxIdleConns:        int(opts.MaxConnections),
		MaxIdleConnsPerHost: int(opts.MaxConnections),
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(opts.Timeout) * time.Second,
			KeepAlive: time.Duration(opts.KeepAlive) * time.Second,
		}).DialContext,
	}

	// Setup main client
	client := &SDK{
		userAgent: opts.UserAgent,
		c: &http.Client{
			Transport: t,
			Timeout:   time.Duration(opts.Timeout) * time.Second,
		},
	}

	// Set client endpoint and services
	client.apiEndpoint = stagingServer
	client.DID = &didService{sdk: client}
	client.VC = &vcService{sdk: client}
	return client, nil
}

// Dispatch a network request to the service
func (i *SDK) request(method, endpoint string, data interface{}, pl map[string]interface{}) error {
	// Get request endpoint
	url := fmt.Sprintf("%s%s", i.apiEndpoint, endpoint)

	// Encode data
	var rd []byte
	if data != nil {
		rd, _ = json.Marshal(data)
	}

	// Build request
	req, _ := http.NewRequest(method, url, bytes.NewReader(rd))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if i.userAgent != "" {
		req.Header.Add("User-Agent", i.userAgent)
	}

	// Execute request
	res, err := i.c.Do(req)
	if err != nil {
		return err
	}

	// Check status
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		return errors.New("internal server error")
	}

	// Get response contents and decode expected payload
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if pl == nil {
		return nil
	}
	return json.Unmarshal(content, &pl)
}
