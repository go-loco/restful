package rest

import (
	"container/list"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
	"unsafe"
)

// Response ...
type Response struct {
	*http.Response
	Err             error
	byteBody        []byte
	listElement     *list.Element
	skipListElement *skipListNode
	ttl             *time.Time
	lastModified    *time.Time
	etag            string
	revalidate      bool
}

func (r *Response) size() int64 {

	size := int64(unsafe.Sizeof(*r))

	size += int64(len(r.byteBody))
	size += int64(unsafe.Sizeof(*r.listElement))
	size += int64(unsafe.Sizeof(*r.skipListElement))
	size += int64(unsafe.Sizeof(*r.ttl))
	size += int64(unsafe.Sizeof(*r.lastModified))
	size += int64(len(r.etag))

	size += int64(len(r.Response.Proto))
	size += int64(len(r.Response.Status))
	/*r.Response.Header
	r.Response.TLS
	r.Response.Trailer
	r.Response.TransferEncoding


	r.Response.Request.Header
	r.Response.Request.Host
	r.Response.Request.Method
	r.Response.Request.Proto
	r.Response.Request.RemoteAddr
	r.Response.Request.RequestURI
	*/

	return size
}

// String return the Respnse Body as a String.
func (r *Response) String() string {
	return string(r.Bytes())
}

// Bytes return the Response Body as bytes.
func (r *Response) Bytes() []byte {
	return r.byteBody
}

// FillUp set the *fill* parameter with the corresponding JSON or XML response.
// fill could be `struct` or `map[string]interface{}`
func (r *Response) FillUp(fill interface{}) error {

	ctypeJSON := "application/json"
	ctypeXML := "application/xml"

	ctype := strings.ToLower(r.Header.Get("Content-Type"))

	for i := 0; i < 2; i++ {

		switch {
		case strings.Contains(ctype, ctypeJSON):
			return json.Unmarshal(r.byteBody, fill)
		case strings.Contains(ctype, ctypeXML):
			return xml.Unmarshal(r.byteBody, fill)
		case i == 0:
			ctype = http.DetectContentType(r.byteBody)
		}

	}

	return errors.New("Response format neither JSON nor XML")

}

func (r *Response) Debug() string {

	var strReq, strResp string

	if req, err := httputil.DumpRequest(r.Request, true); err != nil {
		strReq = err.Error()
	} else {
		strReq = string(req)
	}

	if resp, err := httputil.DumpResponse(r.Response, true); err != nil {
		strResp = err.Error()
	} else {
		strResp = string(resp)
	}

	result := "REQUEST\n"
	result += "--------\n"
	result += strReq
	result += "\n\nRESPONSE\n"
	result += "--------\n"
	result += strResp

	return result
}
