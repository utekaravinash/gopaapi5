package gopaapi5

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/utekaravinash/gopaapi5/api"
)

const (
	doubleSpace         = "  "
	dateFormat          = "20060102T150405Z"
	shortDateFormat     = "20060102"
	hashingAlgorithm    = "AWS4-HMAC-SHA256"
	serviceName         = "ProductAdvertisingAPI"
	terminationString   = "aws4_request"
	authorizationHeader = "Authorization"
)

// request holds information specific to the operation and http request
type request struct {
	operation api.Operation
	payload   map[string]interface{}
	client    *Client
	httpReq   *http.Request
	path      string
	dateTime  time.Time
}

// build builds http request and returns error if any
func (r *request) build() error {
	// Construct API endpoint
	endpoint := fmt.Sprintf("%s://%s/%s", r.scheme(), r.client.host, r.path)
	amzDate := formatDate(r.dateTime)

	// Set common payload values
	r.payload["PartnerType"] = r.client.partnerType
	r.payload["PartnerTag"] = r.client.AssociateTag
	r.payload["Marketplace"] = r.client.marketplace

	// Convert payload to json format
	jsonBody, err := json.Marshal(r.payload)
	if err != nil {
		return err
	}

	// Construct HTTP request
	r.httpReq, err = http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	// Set HTTP request headers
	r.httpReq.Header.Set("host", r.client.host)
	r.httpReq.Header.Set("content-type", "application/json; charset=utf-8")
	r.httpReq.Header.Set("content-encoding", "amz-1.0")
	r.httpReq.Header.Set("x-amz-date", amzDate)
	r.httpReq.Header.Set("x-amz-target", fmt.Sprintf("com.amazon.paapi5.v1.ProductAdvertisingAPIv1.%s", r.operation))

	return nil

}

// sign calculates the request signature and sets Authorization header value
func (r *request) sign() error {
	headers := []string{}
	canonicalHeaders := ""
	signedHeaders := ""
	signedHeaderVals := http.Header{}

	// Set signed header values
	for k, v := range r.httpReq.Header {
		lowerCaseKey := strings.ToLower(k)

		if _, ok := signedHeaderVals[lowerCaseKey]; ok {
			signedHeaderVals[lowerCaseKey] = append(signedHeaderVals[lowerCaseKey], v...)
			continue
		}

		headers = append(headers, lowerCaseKey)
		signedHeaderVals[lowerCaseKey] = v
	}

	sort.Strings(headers)
	signedHeaders = strings.Join(headers, ";")

	headerValues := make([]string, len(headers))
	for i, k := range headers {
		headerValues[i] = k + ":" + strings.Join(signedHeaderVals[k], ",")
	}

	// strip excess space from header values
	stripExcessSpaces(headerValues)
	canonicalHeaders = strings.Join(headerValues, "\n")

	r.httpReq.URL.RawQuery = strings.Replace(r.httpReq.URL.Query().Encode(), "+", "%20", -1)

	uri := getURIPath(r.httpReq.URL)

	body, err := json.Marshal(r.payload)
	if err != nil {
		return err
	}

	// Calculate hash of the request body
	bodyDigest := hex.EncodeToString(hashSHA256(body))

	// Construct canonical string
	canonicalString := strings.Join([]string{
		r.httpReq.Method,
		uri,
		r.httpReq.URL.RawQuery,
		canonicalHeaders + "\n",
		signedHeaders,
		bodyDigest,
	}, "\n")

	// Calculate hash of canonical string
	canonicalStringHash := hex.EncodeToString(hashSHA256([]byte(canonicalString)))

	// Build credential scope
	credentialScope := strings.Join([]string{
		formatShortDate(r.dateTime),
		r.client.region,
		serviceName,
		terminationString,
	}, "/")

	// Build string to sign
	stringToSign := strings.Join([]string{
		hashingAlgorithm,
		formatDate(r.dateTime),
		credentialScope,
		canonicalStringHash,
	}, "\n")

	kDate := hmacSHA256([]byte("AWS4"+r.client.SecretKey), []byte(formatShortDate(r.dateTime)))
	kRegion := hmacSHA256(kDate, []byte(r.client.region))
	kService := hmacSHA256(kRegion, []byte(serviceName))
	signingKey := hmacSHA256(kService, []byte(terminationString))

	// Calculate request signature
	signature := hex.EncodeToString(hmacSHA256(signingKey, []byte(stringToSign)))

	parts := []string{
		fmt.Sprintf("%s Credential=%s/%s", hashingAlgorithm, r.client.AccessKey, credentialScope),
		fmt.Sprintf("SignedHeaders=%s", signedHeaders),
		fmt.Sprintf("Signature=%s", signature),
	}

	// Set Authorization header
	r.httpReq.Header.Set(authorizationHeader, strings.Join(parts, " "))

	return nil
}

// getURIPath gets URI path from the URL
func getURIPath(u *url.URL) string {
	var uri string

	if len(u.Opaque) > 0 {
		uri = "/" + strings.Join(strings.Split(u.Opaque, "/")[3:], "/")
	} else {
		uri = u.EscapedPath()
	}

	if len(uri) == 0 {
		uri = "/"
	}

	return uri
}

// scheme returns the request's scheme
func (r *request) scheme() string {
	if !r.client.testing {
		return "https"
	}

	return "http"
}

// stripExcessSpaces stips excess space at both ends of each value in a slice
func stripExcessSpaces(vals []string) {
	for i, str := range vals {
		vals[i] = strings.TrimSpace(str)
	}
}

// hmacSHA256 returns SHA256 hash of data using a key
func hmacSHA256(key []byte, data []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

// hashSHA256 return SHA25 hash of data without a key
func hashSHA256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

// formatShortDate returns short form of a date
func formatShortDate(dt time.Time) string {
	return dt.UTC().Format(shortDateFormat)
}

// formatDate returns long form of a date
func formatDate(dt time.Time) string {
	return dt.UTC().Format(dateFormat)
}
