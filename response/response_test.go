package response_test

import (
	"encoding/xml"
	"net/http/httptest"
	"testing"

	"github.com/edgarSucre/mye"
	"github.com/edgarSucre/mye/response"
)

func TestHttpJSON(t *testing.T) {

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	err := mye.Forbiden.New("test")
	nerr := response.HttpXml(err, rr)

	if nerr != nil {
		t.FailNow()
	}

	if rr.Code != 403 {
		t.FailNow()
	}

	decoder := xml.NewDecoder(rr.Body)
	msg := &response.ErrResponse{}
	decoder.Decode(msg)

	if msg.Error != "test" {
		t.FailNow()
	}
}
