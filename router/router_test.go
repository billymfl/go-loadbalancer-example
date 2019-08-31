package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/billymfl/go-loadbalancer-example/config"
)

type Expect struct {
	status int
	method string
	path   string
	value  string
}

// Add static route info here
var expected = []*Expect{
	{200, "GET", "/", fmt.Sprintf("%s %s", config.Name(), config.Version())},
	{200, "GET", "/healthcheck", `{"alive": true}`},
}

func TestStaticRoutes(t *testing.T) {
	for _, exp := range expected {

		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest(exp.method, exp.path, nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		Router.ServeHTTP(rr, req)

		// Check the status code is what we expect
		if status := rr.Code; status != exp.status {
			t.Errorf("%s returned wrong status code: got %v want %v",
				exp.path, status, exp.status)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != exp.value {
			t.Errorf("%s returned unexpected body: got %v want %v",
				exp.path, rr.Body.String(), exp.value)
		}
	}
}

func TestRegisterNoAuth(t *testing.T) {
	req, err := http.NewRequest("PUT", "/register/room_name/8/0/v1.0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}
}

func TestRegisterWithAuth(t *testing.T) {
	req, err := http.NewRequest("PUT", "/register/room_name/8/0/v1.0", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-API-Key", "mock_key")

	rr := httptest.NewRecorder()

	Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
