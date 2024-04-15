package integration_test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"

)

const (
	// Attempts connection
	host       = "localhost:8080"
	healthPath = "http://" + host + "/healthz"
	attempts   = 20

	// HTTP REST
	basePath = "http://" + host + "/v1"

)

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

// HTTP POST: /users/create.
func TestHTTPcreateUser(t *testing.T) {
	body := `{
		"full_name": "Doston Shernazarov",
		"username": "doston",
		"email": "dostonshernazarov2001@gmail.com",
		"password": "1234abcd",
		"bio": "Life is unpredictable",
		"website": "test.com"
	}`
	Test(t,
		Description("Create user Success"),
		Post(basePath+"/users/create"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".bio").Equal("Life is unpredictable"),
	)

	body = `{
		"full_name": "",
		"username": "doston",
		"email": "dostonshernazarov2001@gmail.com",
		"password": "1234abcd",
		"bio": "Life is unpredictable",
		"website": "test.com"
	}`
	Test(t,
		Description("Create user Fail"),
		Post(basePath+"/users/create"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").Equal("invalid request body"),
	)
}

// HTTP GET: /users/:page/:limit.
func TestHTTPHistory(t *testing.T) {
	Test(t,
		Description("History Success"),
		Get(basePath+"/users/1/10"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Contains(`{"users":[{`),
	)
}