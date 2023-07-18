package response

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/edgarSucre/mye"
)

type (
	httpResponseType string
	ErrorTypeCode    struct {
		http int
	}
	ErrResponse struct {
		XMLName xml.Name `xml:"response"`
		Error   string   `json:"Error" xml:"Error"`
	}
)

const (
	charsetUTF8       = "charset=UTF-8"
	HeaderContentType = "Content-Type"

	MIMEApplicationJSON            httpResponseType = "application/json"
	MIMEApplicationJSONCharsetUTF8 httpResponseType = MIMEApplicationJSON + "; " + charsetUTF8
	MIMEApplicationXML             httpResponseType = "application/xml"
	MIMEApplicationXMLCharsetUTF8  httpResponseType = MIMEApplicationXML + "; " + charsetUTF8
)

var (
	errCode = map[mye.ErrorType]ErrorTypeCode{
		mye.Cancelation:  {http.StatusInternalServerError},
		mye.Forbiden:     {http.StatusForbidden},
		mye.Internal:     {http.StatusInternalServerError},
		mye.NotFound:     {http.StatusNotFound},
		mye.Timeout:      {http.StatusForbidden},
		mye.Unauthorized: {http.StatusUnauthorized},
		mye.Validation:   {http.StatusBadRequest},
	}
)

func statusContent(err error) (int, string) {
	uErr := mye.UnWrap(err)
	if local, ok := err.(mye.Err); ok {
		return errCode[local.T].http, uErr.Error()
	}

	return 500, uErr.Error()
}

func Http(err error, w http.ResponseWriter, content []byte) error {
	status, _ := statusContent(err)
	w.WriteHeader(status)
	_, err = w.Write(content)

	return err
}

func HttpJSON(err error, w http.ResponseWriter) error {
	status, msg := statusContent(err)
	errResponse := ErrResponse{Error: msg}

	w.Header().Set(HeaderContentType, string(MIMEApplicationJSONCharsetUTF8))
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	return encoder.Encode(errResponse)
}

func HttpXml(err error, w http.ResponseWriter) error {
	status, msg := statusContent(err)
	errResponse := ErrResponse{Error: msg}

	w.Header().Set(HeaderContentType, string(MIMEApplicationXMLCharsetUTF8))
	w.WriteHeader(status)

	encoder := xml.NewEncoder(w)
	return encoder.Encode(errResponse)
}

type JSONSerializer interface {
	JSON(code int, i interface{}) error
}

func Echo(err error, ctx JSONSerializer) error {
	status, msg := statusContent(err)
	errResponse := ErrResponse{Error: msg}

	return ctx.JSON(status, errResponse)
}
