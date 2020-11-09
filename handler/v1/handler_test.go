package v1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	v1 "github.com/mtrang1263/tdd/handler/v1"
)

func TestGetCurrentDate(t *testing.T) {
	rr := httptest.NewRecorder()

	// Testing successful request
	reqGet, err := http.NewRequest("GET", "http://localhost/", nil)
	if err != nil {
		t.Errorf("Could not create request %s", err.Error())
	}

	v1.GetCurrentDate(rr, reqGet)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 got %d", rr.Code)
	}

	body := rr.Body.String()
	date, err := time.Parse(time.RFC3339, body)
	if err != nil {
		t.Errorf("Could not parse body %s %s", body, err.Error())
	}
	if date.IsZero() {
		t.Errorf("Time should not be zero!")
	}

	// Testing bad request
	rr = httptest.NewRecorder()
	reqPost, err := http.NewRequest("POST", "http://localhost/", nil)
	if err != nil {
		t.Errorf("Could not create request %s", err.Error())
	}
	v1.GetCurrentDate(rr, reqPost)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 got %d", rr.Code)
	}
}

func TestGetCurrentDateValid(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost/", nil)
	if err != nil {
		t.Errorf("Could not create request %s", err.Error())
	}

	v1.GetCurrentDate(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 got %d", rr.Code)
	}

	body := rr.Body.String()
	date, err := time.Parse(time.RFC3339, body)
	if err != nil {
		t.Errorf("Could not parse body %s %s", body, err.Error())
	}
	if date.IsZero() {
		t.Errorf("Time should not be zero!")
	}
}
func TestGetCurrentDateInvalid(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "http://localhost/", nil)
	if err != nil {
		t.Errorf("Could not create request %s", err.Error())
	}
	v1.GetCurrentDate(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 got %d", rr.Code)
	}
}

type HandlerTestSuite struct {
	suite.Suite

	rr *httptest.ResponseRecorder
}

func (s *HandlerTestSuite) SetupTest() {
	s.rr = httptest.NewRecorder()
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (s *HandlerTestSuite) TestGetCurrentDateValid() {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost/", nil)
	s.NoError(err, "Could not create request")

	v1.GetCurrentDate(rr, req)
	s.Equal(http.StatusOK, 200)

	body := rr.Body.String()
	date, err := time.Parse(time.RFC3339, body)
	s.NoError(err, "Could not parse body %s", body)
	s.False(date.IsZero())
}
func (s *HandlerTestSuite) TestGetCurrentDateInvalid() {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "http://localhost/", nil)
	s.NoError(err, "Could not create request")

	v1.GetCurrentDate(rr, req)
	s.Equal(http.StatusBadRequest, rr.Code)
}

func TestGetCurrentDateTDD(t *testing.T) {

	tests := []struct {
		name      string
		method    string
		url       string
		respCode  int
		expFormat string
	}{
		{"Successful: No format", "GET", "http://localhost/", http.StatusOK, time.RFC3339},
		{"Successful: RFC3339", "GET", "http://localhost/?format=" + time.RFC3339Nano, http.StatusOK, time.RFC3339},
		{"Successful: RFC1123", "GET", "http://localhost/?format=" + time.RFC1123, http.StatusOK, time.RFC1123},
		{"Successful: Custom Format", "GET", "http://localhost/?format=MARTIN_COOL_FORMAT", http.StatusOK, "02-2006-01"},
		{"Bad Request: POST", "POST", "http://localhost/", http.StatusBadRequest, ""},
		{"Bad Request: PUT", "PUT", "http://localhost/", http.StatusBadRequest, ""},
		{"Bad Request: DELETE", "DELETE", "http://localhost/", http.StatusBadRequest, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Errorf("Could not create request %s", err.Error())
			}
			v1.GetCurrentDate(rr, req)
			if tt.respCode != rr.Code {
				t.Errorf("Expected %d received %d", tt.respCode, rr.Code)
			}
			if tt.respCode != http.StatusOK {
				body := rr.Body.String()
				if body != "Only GET requests are supported" {
					t.Errorf("Unexpected body %s returned.", body)
				}
				return
			}

			body := rr.Body.String()
			date, err := time.Parse(tt.expFormat, body)
			if err != nil {
				t.Errorf("Could not parse body %s %s", body, err.Error())
			}
			if date.IsZero() {
				t.Errorf("Time should not be zero!")
			}
		})
	}
}
