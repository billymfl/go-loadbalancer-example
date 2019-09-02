package router

import (
	"fmt"
	"io"
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
	secure bool
}

// Add static route info here
var expected = []*Expect{
	{200, "GET", "/", fmt.Sprintf(`{"app": "%s", "version": "%s"}`, config.Name(), config.Version()), false},
	{200, "GET", "/healthcheck", `{"alive": true}`, false},
}

func TestStaticRoutes(t *testing.T) {
	for _, exp := range expected {

		rr, req, err := request(exp.method, exp.path, nil, exp.secure)
		if err != nil {
			t.Fatal(err)
		}

		Router.ServeHTTP(rr, req)

		// Check the status code is what we expect
		checkstatus(rr.Code, exp.status, t)
		// if status := rr.Code; status != exp.status {
		// 	t.Errorf("%s returned wrong status code: got %v want %v",
		// 		exp.path, status, exp.status)
		// }

		// Check the response body is what we expect.
		if rr.Body.String() != exp.value {
			t.Errorf("%s returned unexpected body: got %v want %v",
				exp.path, rr.Body.String(), exp.value)
		}
	}
}

func TestRegisterWithoutAuth(t *testing.T) {
	//set secure to false so api key header is not set
	rr, req, err := request("PUT", "/register/room_name/8/0/v1.0", nil, false)
	if err != nil {
		t.Fatal(err)
	}
	Router.ServeHTTP(rr, req)
	checkstatus(rr.Code, http.StatusForbidden, t)
}

func TestRegisterWithAuth(t *testing.T) {
	rr, req, err := request("PUT", "/register/room_name/8/0/v1.0", nil, true)
	if err != nil {
		t.Fatal(err)
	}
	Router.ServeHTTP(rr, req)
	checkstatus(rr.Code, http.StatusOK, t)
}

func TestUnregister(t *testing.T) {
	rr, req, err := request("DELETE", "/register/room_name", nil, true)
	if err != nil {
		t.Fatal(err)
	}
	Router.ServeHTTP(rr, req)
	checkstatus(rr.Code, http.StatusOK, t)

	expected := `{"message": "unregistered"}`
	if rr.Body.String() != expected {
		t.Errorf("Unregister returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// func TestList(t *testing.T) {
// 	rr, req, err = request("GET", "/list", nil, true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	Router.ServeHTTP(rr, req)
// 	checkstatus(rr.Code, http.StatusOK, t)

// 	expected := ""
// 	if rr.Body.String() != expected {
// 		t.Errorf("Unregister returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }

func request(method string, url string, body io.Reader, secure bool) (*httptest.ResponseRecorder, *http.Request, error) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	if secure {
		req.Header.Set("X-API-Key", "mock_key")
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	return rr, req, nil
}

//checkstatus checks that the expstatus == code, else fails the test
func checkstatus(code int, expstatus int, t *testing.T) {
	if status := code; status != expstatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expstatus)
	}
}
