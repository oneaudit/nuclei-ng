package proxy

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/elazarl/goproxy"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/useragent"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

var cache sync.Map
var defaultUserAgent = useragent.PickRandom().String()

func CreateProxy(options *types.Options) *goproxy.ProxyHttpServer {
	proxy := goproxy.NewProxyHttpServer()
	if options.ProxyHost != "" {
		// Manipulate requests/responses
		proxy.OnRequest().DoFunc(handleRequest)
		proxy.OnResponse().DoFunc(handleResponse)
		// Forward requests
		proxy.Tr.Proxy = http.ProxyURL(&url.URL{
			Scheme: "http",
			Host:   options.ProxyHost,
		})
	}
	return proxy
}

func handleRequest(req *http.Request, _ *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	if strings.HasSuffix(req.UserAgent(), "(proxy)") {
		req.Header.Set("User-Agent", defaultUserAgent)

		hash, err := computeRequestHash(req)
		if err != nil {
			gologger.Warning().Msgf("Error computing request hash: %s", err.Error())
		} else {
			if value, ok := cache.Load(hash); ok {
				if value == nil {
					value, err = waitForResponse(hash)
					if err != nil {
						gologger.Warning().Msgf("Error waiting for cached response: %s", err.Error())
						return req, nil
					}
				}

				if resp, ok := value.(cachedResponse); ok {
					gologger.Debug().Msgf("Using cache for hash: %s", hash)
					return req, resp.withRequest(req)
				}
			} else {
				cache.Store(hash, nil)
			}
		}

		// perform the request and add it to the cache
		// [snip]

		// This is a dummy internal header
		req.Header.Add("X-Request-Cache", hash)
	}
	return req, nil
}

func waitForResponse(key string) (*cachedResponse, error) {
	responseChan := make(chan *cachedResponse)

	go func() {
		for {
			if value, ok := cache.Load(key); ok {
				if resp, ok := value.(*cachedResponse); ok && resp != nil {
					responseChan <- resp
					return
				}
			}
			time.Sleep(100 * time.Millisecond) // Polling interval
		}
	}()

	return <-responseChan, nil
}

type cachedResponse struct {
	header http.Header
	status int
	body   string
}

func (c *cachedResponse) withRequest(req *http.Request) *http.Response {
	resp := &http.Response{}
	resp.Request = req
	resp.TransferEncoding = req.TransferEncoding
	resp.Header = c.header
	resp.StatusCode = c.status
	resp.Status = http.StatusText(c.status)
	buf := bytes.NewBufferString(c.body)
	resp.ContentLength = int64(buf.Len())
	resp.Body = io.NopCloser(buf)
	return resp
}

func handleResponse(resp *http.Response, _ *goproxy.ProxyCtx) *http.Response {
	var body string
	if resp.Body != nil {
		raw, _ := io.ReadAll(resp.Body)
		body = string(raw)
	}
	cachedResp := &cachedResponse{
		header: resp.Header,
		status: resp.StatusCode,
		body:   body,
	}

	if _, ok := resp.Request.Header["X-Request-Cache"]; ok {
		hash := resp.Request.Header["X-Request-Cache"][0]
		cache.Store(hash, cachedResp)
	}

	return cachedResp.withRequest(resp.Request)
}

func computeRequestHash(req *http.Request) (string, error) {
	hasher := sha256.New()

	// Method + URL
	_, err := hasher.Write([]byte(req.Method + req.URL.String()))
	if err != nil {
		return "", err
	}

	// Headers
	var keys []string
	for key := range req.Header {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if key == "User-Agent" {
			continue
		}

		values := req.Header[key]
		for _, value := range values {
			_, err := hasher.Write([]byte(key + ": " + value + "\n"))
			if err != nil {
				return "", err
			}
		}
	}

	// Request Body
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return "", err
		}
		req.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))

		_, err = hasher.Write(bodyBytes)
		if err != nil {
			return "", err
		}
	}

	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash), nil
}
