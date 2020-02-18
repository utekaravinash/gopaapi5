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

type request struct {
	Operation Operation
	Payload   map[string]interface{}
	client    *Client
	httpReq   *http.Request
	path      string
	dateTime  time.Time
}

func (r *request) build() error {
	endpoint := fmt.Sprintf("https://%s/%s", r.client.host, r.path)
	r.dateTime = time.Now().UTC()
	amzDate := formatDate(r.dateTime)

	r.Payload["PartnerType"] = r.client.partnerType
	r.Payload["PartnerTag"] = r.client.AssociateTag
	r.Payload["Marketplace"] = r.client.marketplace

	jsonBody, err := json.Marshal(r.Payload)
	if err != nil {
		return err
	}

	r.httpReq, err = http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	r.httpReq.Header.Set("host", r.client.host)
	r.httpReq.Header.Set("content-type", "application/json; charset=utf-8")
	r.httpReq.Header.Set("content-encoding", "amz-1.0")
	r.httpReq.Header.Set("x-amz-date", amzDate)
	r.httpReq.Header.Set("x-amz-target", fmt.Sprintf("com.amazon.paapi5.v1.ProductAdvertisingAPIv1.%s", r.Operation))

	return nil

}

func (r *request) sign() error {
	// mac := hmac.New(sha256.New, []byte(c.SecretKey))

	// authorization := "NewAuthorization"
	// r.httpReq.Header.Set("Authorization", authorization)

	headers := []string{}
	canonicalHeaders := ""
	signedHeaders := ""
	signedHeaderVals := http.Header{}

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

	stripExcessSpaces(headerValues)
	canonicalHeaders = strings.Join(headerValues, "\n")

	r.httpReq.URL.RawQuery = strings.Replace(r.httpReq.URL.Query().Encode(), "+", "%20", -1)

	uri := getURIPath(r.httpReq.URL)

	body, err := json.Marshal(r.Payload)
	if err != nil {
		return err
	}

	bodyDigest := hex.EncodeToString(hashSHA256(body))

	canonicalString := strings.Join([]string{
		r.httpReq.Method,
		uri,
		r.httpReq.URL.RawQuery,
		canonicalHeaders + "\n",
		signedHeaders,
		bodyDigest,
	}, "\n")

	canonicalStringHash := hex.EncodeToString(hashSHA256([]byte(canonicalString)))

	credentialScope := strings.Join([]string{
		formatShortDate(r.dateTime),
		r.client.region,
		serviceName,
		terminationString,
	}, "/")

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

	signature := hex.EncodeToString(hmacSHA256(signingKey, []byte(stringToSign)))

	parts := []string{
		fmt.Sprintf("%s Credential=%s/%s", hashingAlgorithm, r.client.AccessKey, credentialScope),
		fmt.Sprintf("SignedHeaders=%s", signedHeaders),
		fmt.Sprintf("Signature=%s", signature),
	}

	r.httpReq.Header.Set(authorizationHeader, strings.Join(parts, " "))

	return nil
}

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

func stripExcessSpaces(vals []string) {
	var j, k, l, m, spaces int
	for i, str := range vals {
		// Trim trailing spaces
		for j = len(str) - 1; j >= 0 && str[j] == ' '; j-- {
		}

		// Trim leading spaces
		for k = 0; k < j && str[k] == ' '; k++ {
		}
		str = str[k : j+1]

		// Strip multiple spaces.
		j = strings.Index(str, doubleSpace)
		if j < 0 {
			vals[i] = str
			continue
		}

		buf := []byte(str)
		for k, m, l = j, j, len(buf); k < l; k++ {
			if buf[k] == ' ' {
				if spaces == 0 {
					// First space.
					buf[m] = buf[k]
					m++
				}
				spaces++
			} else {
				// End of multiple spaces.
				spaces = 0
				buf[m] = buf[k]
				m++
			}
		}

		vals[i] = string(buf[:m])
	}
}

func hmacSHA256(key []byte, data []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

func hashSHA256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func formatShortDate(dt time.Time) string {
	return dt.UTC().Format(shortDateFormat)
}

func formatDate(dt time.Time) string {
	return dt.UTC().Format(dateFormat)
}
