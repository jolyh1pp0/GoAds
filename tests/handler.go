package tests

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func AuthPrepare(t *testing.T, email string, password string, e *echo.Echo) (string, string) {
	rec := httptest.NewRecorder()
	var req *http.Request

	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var buf bytes.Buffer

	user := User {
		Email: email,
		Password: password,
	}

	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		assert.NotEmpty(t, err)
	}

	req = httptest.NewRequest(http.MethodGet, "/login", &buf)

	e.ServeHTTP(rec, req)

	bodyString := make(map[string]string)
	json.Unmarshal(rec.Body.Bytes(), &bodyString)

	return bodyString["access_token"], ""
}

func DoRequest(t *testing.T, input interface{}, path, method string, e *echo.Echo, access string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	var req *http.Request

	if input != nil {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		if err != nil {
			assert.NotEmpty(t, err)
		}

		req = httptest.NewRequest(method, path, &buf)
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	if access != "" {
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+access)
	}

	e.ServeHTTP(rec, req)

	return rec
}
