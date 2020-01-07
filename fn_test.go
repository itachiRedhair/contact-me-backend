package contactmebackend

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSendEmail(t *testing.T) {
	tests := []struct {
		body       string
		wantStatus int
	}{
		{body: `{"name": "Itachi Redhair", "emailAddress": "itachi.redhair@gmail.com", "message": "Would love to meet over a beer!"}`, wantStatus: http.StatusOK},
		{body: `{"name": "Itachi Redhair", "message": "Would love to meet over a beer!"}`, wantStatus: http.StatusInternalServerError},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.body))
		req.Header.Add("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		SendEmail(rr, req)

		if got := rr.Result().StatusCode; got != test.wantStatus {
			t.Errorf("SendEmail(%q) = %q, want %q", test.body, got, test.wantStatus)
		}
	}
}
