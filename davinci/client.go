// Davinci Admin API GO Client
//
// This package is go client to be used for interacting with PingOne DaVinci Administrative APIs.
// Use cases include:
// - Creating Connections
// - Importing Flows
package davinci

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

// baseURL
var baseURL = url.URL{
	Scheme: "https",
	Host:   "orchestrate-api.pingone.com",
	Path:   "/v1",
}

var dvApiHost = map[string]string{
	"NorthAmerica": "orchestrate-api.pingone.com",
	"Europe":       "orchestrate-api.pingone.eu",
	"AsiaPacific":  "orchestrate-api.pingone.asia",
	"Canada":       "orchestrate-api.pingone.ca",
}

var defaultUserAgent = "PingOne-DaVinci-GOLANG-SDK"

// const HostURL string = "https://api.singularkey.com/v1"

func (args Params) QueryParams() url.Values {
	q := make(url.Values)

	if args.Page != "" {
		q.Add("page", args.Page)
	}

	if args.Limit != "" {
		q.Add("limit", args.Limit)
	}
	for i, v := range args.ExtraParams {
		q.Add(i, v)
	}

	return q
}

func NewClient(inputs *ClientInput) (*APIClient, error) {
	// adjust host according to received region
	if inputs.PingOneRegion == "" {
		return nil, fmt.Errorf("PingOneRegion must be set")
	} else {
		if dvApiHost[inputs.PingOneRegion] == "" {
			return nil, fmt.Errorf("Invalid region: %v", inputs.PingOneRegion)
		}
		baseURL.Host = dvApiHost[inputs.PingOneRegion]
	}

	hostUrl := baseURL.ResolveReference(&url.URL{}).String()

	if inputs.HostURL != "" {
		hostUrl = inputs.HostURL
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("Got error while creating cookie jar %s", err.Error())
	}
	c := APIClient{
		HTTPClient: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 300 * time.Second,
			Jar:     jar},
		HostURL: hostUrl,
	}

	if inputs.AccessToken != "" {
		c.Token = inputs.AccessToken
		c.CompanyID = inputs.PingOneSSOEnvId
		return &c, nil
	}

	if inputs.Username == "" || inputs.Password == "" {
		// return nil, fmt.Errorf("User or Password not found")
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: inputs.Username,
		Password: inputs.Password,
	}

	c.UserAgent = defaultUserAgent

	if inputs.UserAgent != "" {
		c.UserAgent = inputs.UserAgent
	}

	// Use P1SSO if available
	if inputs.PingOneSSOEnvId != "" {
		c.PingOneSSOEnvId = inputs.PingOneSSOEnvId
	}
	err = c.DoSignIn(nil)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &c, nil
}

func (c *APIClient) DoSignIn(targetCompanyId *string) error {
	if c.PingOneSSOEnvId != "" {
		ar, _, err := c.SignInSSOWithResponse(targetCompanyId)
		if err != nil {
			return err
		}

		c.Token = ar.AccessToken
		return nil
	}

	return fmt.Errorf("Sign in failed. Not using SSO")
}

func (c *APIClient) doRequestVerbose(req *http.Request, authToken *string, args *Params) (*DvHttpResponse, *http.Response, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
		var bearer = "Bearer " + token
		req.Header.Add("Authorization", bearer)
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if args != nil {
		req.URL.RawQuery = args.QueryParams().Encode()
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()

	rbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusFound {
		return nil, res, fmt.Errorf("status: %d, body: %s", res.StatusCode, rbody)
	}
	resp := DvHttpResponse{
		Body:       rbody,
		Headers:    res.Header,
		StatusCode: res.StatusCode,
	}

	if res.StatusCode == http.StatusFound && res.Header["Location"] != nil {
		resp.Location, _ = url.Parse(res.Header["Location"][0])
		resp.LocationParams, _ = url.ParseQuery(resp.Location.RawQuery)
		// Handle wepbage hash value strangeness
		if resp.Location.Fragment != "" {
			_, v, ok := strings.Cut(resp.Location.Fragment, "?")
			if ok {
				resp.LocationParams, _ = url.ParseQuery(v)
			}
		}
	}
	if res.Header["Set-Cookie"] != nil {
		c.HTTPClient.Jar.SetCookies(req.URL, res.Cookies())
	}

	return &resp, res, err
}

func (c *APIClient) doRequest(req *http.Request, args *Params) ([]byte, *http.Response, error) {

	if c.Token != "" {
		req.Header.Del("Authorization")

		var bearer = "Bearer " + c.Token
		req.Header.Add("Authorization", bearer)
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if args != nil {
		req.URL.RawQuery = args.QueryParams().Encode()
	}
	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, nil, err
	}
	log.Printf("\n%s\n", string(dump))
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	dump, err = httputil.DumpResponse(res, true)
	if err != nil {
		return nil, nil, err
	}
	log.Printf("\n%s\n", string(dump))

	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode >= 300 {
		var errObject ErrorResponse

		if err := json.Unmarshal(body, &errObject); err != nil {
			return nil, res, err
		}

		b, err := json.Marshal(errObject)
		if err != nil {
			return nil, res, err
		}

		if string(b) != "{}" {
			log.Printf("Error response handled: %s", string(b))
			return nil, res, errObject
		} else {
			log.Printf("Error response unhandled: %d", res.StatusCode)
		}
	}

	return body, res, err
}

func (c *APIClient) doRequestRetryable(companyId *string, req DvHttpRequest, args *Params) ([]byte, *http.Response, error) {

	// This API action isn't thread safe - the environment may be switched by another thread.  We need to lock it
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// handle environment switching
	if companyId != nil && *companyId != c.CompanyID {
		_, res, err := c.SetEnvironmentWithResponse(companyId)
		if err != nil {
			return nil, res, err
		}

		if c.CompanyID != *companyId {
			return nil, nil, fmt.Errorf("Failed to set environment to %s after successful switch", *companyId)
		}
	}

	reqInit, err := http.NewRequest(req.Method, req.Url, req.Body)
	if err != nil {
		return nil, nil, err
	}

	body, res, err := exponentialBackOffRetry(func() (any, *http.Response, error) {
		return c.doRequest(reqInit, args)
	})
	if err != nil {
		return nil, res, err
	}

	return body.([]byte), res, nil
}

type SDKInterfaceFunc func() (any, *http.Response, error)

var (
	maxRetries               = 5
	maximumRetryAfterBackoff = 30
)

func exponentialBackOffRetry(f SDKInterfaceFunc) (interface{}, *http.Response, error) {
	var obj interface{}
	var resp *http.Response
	var err error
	backOffTime := time.Second
	var isRetryable bool

	for i := 0; i < maxRetries; i++ {
		obj, resp, err = f()

		backOffTime, isRetryable = testForRetryable(resp, err, backOffTime)

		if isRetryable {
			log.Printf("Attempt %d failed: %v, backing off by %s.", i+1, err, backOffTime.String())
			time.Sleep(backOffTime)
			continue
		}

		return obj, resp, err
	}

	log.Printf("Request failed after %d attempts: %s", maxRetries, err)

	return obj, resp, err // output the final error
}

func testForRetryable(r *http.Response, err error, currentBackoff time.Duration) (time.Duration, bool) {

	backoff := currentBackoff

	if r != nil {
		if r.StatusCode == http.StatusNotImplemented || r.StatusCode == http.StatusServiceUnavailable || r.StatusCode == http.StatusTooManyRequests {
			retryAfter, err := parseRetryAfterHeader(r)
			if err != nil {
				log.Printf("Cannot parse the expected \"Retry-After\" header: %s", err)
				backoff = currentBackoff * 2
			}

			if retryAfter <= time.Duration(maximumRetryAfterBackoff) {
				backoff += time.Duration(maximumRetryAfterBackoff)
			} else {
				backoff += retryAfter
			}
		} else {
			backoff = currentBackoff
		}

		retryAbleCodes := []int{
			http.StatusTooManyRequests,
			http.StatusInternalServerError,
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
		}

		if slices.Contains(retryAbleCodes, r.StatusCode) {
			log.Printf("HTTP status code %d detected, available for retry", r.StatusCode)
			return backoff, true
		}
	}

	if err != nil {
		if res1, matchErr := regexp.MatchString(`^http: ContentLength=[0-9]+ with Body length [0-9]+$`, err.Error()); matchErr == nil && res1 {
			log.Printf("HTTP content error detected, available for retry: %v", err)
			backoff += (2 * time.Second)
			return backoff, true
		}
	}

	return backoff, false
}

func parseRetryAfterHeader(resp *http.Response) (time.Duration, error) {
	retryAfterHeader := resp.Header.Get("Retry-After")

	if retryAfterHeader == "" {
		return 0, fmt.Errorf("Retry-After header not found")
	}

	retryAfterSeconds, err := strconv.Atoi(retryAfterHeader)

	if err == nil {
		return time.Duration(retryAfterSeconds) * time.Second, nil
	}

	retryAfterTime, err := http.ParseTime(retryAfterHeader)

	if err != nil {
		return 0, fmt.Errorf("Unable to parse Retry-After header value: %v", err)
	}

	return time.Until(retryAfterTime), nil
}
