package tests

import (
	"GoAds/domain"
	"GoAds/tests/test_models"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

var path = "/users"

func TestGetUsers(t *testing.T) {
	access, _ := AuthPrepare(t, "vasya.pupkin@gmail.com", "SuPeRpassWORD124214", EchoServer)

	respStruct := []test_models.TestGetUserResp{
		{
			ID:        "a2c4a3db-2397-494b-93ef-3da38be95a70",
			Email:     "vasya.pupkin@gmail.com",
			FirstName: "Vasya",
			LastName:  "Pupkin",
		},
		{
			ID:        "a2c4a3db-2397-494b-93ef-3da38be95a71",
			Email:     "petya.pupkin@gmail.com",
			FirstName: "Petya",
			LastName:  "Pupkin",
		},
	}

	samples := []struct {
		access string
		err    *echo.HTTPError
	}{
		{
			access: access,
			err:    nil,
		},
		{
			access: "",
			err: domain.ErrInvalidAccessToken,
		},
	}

	for _, want := range samples {
		rec := DoRequest(t, nil, path, http.MethodGet, EchoServer, want.access)

		t.Logf("rec.Code: %d", rec.Code)
		if want.err == nil {
			if !assert.Equal(t, http.StatusOK, rec.Code) {
				t.Errorf("Body: %s", string(rec.Body.Bytes()))
			}

			var data []test_models.TestGetUserResp
			json.Unmarshal(rec.Body.Bytes(), &data)
			assert.Equal(t, respStruct, data)
		} else {
			assert.Equal(t, want.err.Code, rec.Code)
			errResp, err := json.Marshal(want.err.Message)
			assert.Nil(t, err)

			type Err struct {
				Message string `json:"message"`
			}

			var errMsg Err
			err = json.Unmarshal(rec.Body.Bytes(), &errMsg)
			if err != nil {
				log.Print(err)
			}
			message, _ := json.Marshal(errMsg.Message)

			assert.Equal(t, string(errResp), string(message))
		}
	}
}

func TestGetOneUser(t *testing.T) {
	access, _ := AuthPrepare(t, "vasya.pupkin@gmail.com", "SuPeRpassWORD124214", EchoServer)

	respStruct := []test_models.TestGetUserResp{
		{
			ID:        "a2c4a3db-2397-494b-93ef-3da38be95a71",
			Email:     "petya.pupkin@gmail.com",
			FirstName: "Petya",
			LastName:  "Pupkin",
		},
	}

	samples := []struct {
		access string
		userID string
		err    *echo.HTTPError
	}{
		{
			access: access,
			userID: "a2c4a3db-2397-494b-93ef-3da38be95a71",
			err:    nil,
		},
		{
			access: "",
			err: domain.ErrInvalidAccessToken,
		},
	}

	for _, want := range samples {
		rec := DoRequest(t, nil, path + "/" + want.userID, http.MethodGet, EchoServer, want.access)

		t.Logf("rec.Code: %d", rec.Code)
		if want.err == nil {
			if !assert.Equal(t, http.StatusOK, rec.Code) {
				t.Errorf("Body: %s", string(rec.Body.Bytes()))
			}

			var data []test_models.TestGetUserResp
			json.Unmarshal(rec.Body.Bytes(), &data)
			assert.Equal(t, respStruct, data)
		} else {
			assert.Equal(t, want.err.Code, rec.Code)
			errResp, err := json.Marshal(want.err.Message)
			assert.Nil(t, err)

			type Err struct {
				Message string `json:"message"`
			}

			var errMsg Err
			err = json.Unmarshal(rec.Body.Bytes(), &errMsg)
			if err != nil {
				log.Print(err)
			}
			message, _ := json.Marshal(errMsg.Message)

			assert.Equal(t, string(errResp), string(message))
		}
	}
}

func TestCreateUser(t *testing.T) {
	samples := []struct {
		input test_models.TestCreateUser
		err   *echo.HTTPError
	}{
		{
			input: test_models.TestCreateUser{
				Email: "example@gmail.com",
				Password: "eXaMplE_passWORD123",
				FirstName: "Example",
				LastName: "Example",
				Phone: "+1234567890",
				VerifiedType: "email",
			},
			err:    nil,
		},
		{
			input: test_models.TestCreateUser{
				Email: "",
				Password: "eXaMplE_passWORD123",
				FirstName: "Example",
				LastName: "Example",
				Phone: "+1234",
				VerifiedType: "email",
			},
			err: domain.ErrInvalidEmailAddress,
		},
		{
			input: test_models.TestCreateUser{
				Email: "example@gmail.com",
				Password: "eXaMplE_passWORD123",
				FirstName: "Example",
				LastName: "Example",
				Phone: "+1234",
				VerifiedType: "email",
			},
			err: domain.ErrUserEmailAlreadyExists,
		},
		{
			input: test_models.TestCreateUser{
				Email: "test@gmail.com",
				Password: "eXaMplE_passWORD123",
				FirstName: "Example",
				LastName: "Example",
				Phone: "+123",
				VerifiedType: "email",
			},
			err: domain.ErrUserPhoneAlreadyExists,
		},
		{
			input: test_models.TestCreateUser{
				Email: "test@gmail.com",
				Password: "pass",
				FirstName: "Example",
				LastName: "Example",
				Phone: "+12345",
				VerifiedType: "email",
			},
			err: domain.ErrInsecurePassword,
		},
	}

	for _, want := range samples {
		rec := DoRequest(t, want.input, "/register", http.MethodPost, EchoServer, "")

		t.Logf("rec.Code: %d", rec.Code)
		if want.err == nil {
			if !assert.Equal(t, http.StatusCreated, rec.Code) {
				t.Errorf("Body: %s", string(rec.Body.Bytes()))
			}

			var data test_models.TestCreateUserResponse
			json.Unmarshal(rec.Body.Bytes(), &data)

			respStruct := test_models.TestCreateUserResponse{
				ID:		      data.ID,
				Email:        want.input.Email,
				Password:     data.Password,
				FirstName:    want.input.FirstName,
				LastName:     want.input.LastName,
				Phone:        want.input.Phone,
				VerifiedType: want.input.VerifiedType,
				CreatedAt: 	  data.CreatedAt,
				UpdatedAt: 	  data.UpdatedAt,
			}

			assert.Equal(t, respStruct, data)
		} else {
			assert.Equal(t, want.err.Code, rec.Code)
			errResp, err := json.Marshal(want.err.Message)
			assert.Nil(t, err)

			type Err struct {
				Message string `json:"message"`
			}

			var errMsg Err
			err = json.Unmarshal(rec.Body.Bytes(), &errMsg)
			if err != nil {
				log.Print(err)
			}
			message, _ := json.Marshal(errMsg.Message)

			assert.Equal(t, string(errResp), string(message))
		}
	}
}

func TestUpdateUser(t *testing.T) {
	access, _ := AuthPrepare(t, "vasya.pupkin@gmail.com", "SuPeRpassWORD124214", EchoServer)

	respStruct := test_models.TestGetUserUpdateResp{
		FirstName: "Petr",
	}

	samples := []struct {
		access    string
		userID    string
		input test_models.TestSetUserUpdate
		err       *echo.HTTPError
	}{
		{
			access: access,
			userID: "a2c4a3db-2397-494b-93ef-3da38be95a71",
			input: test_models.TestSetUserUpdate{
				FirstName: "Petr",
			},
			err:    nil,
		},
		{
			access: "",
			err: domain.ErrInvalidAccessToken,
		},
	}

	for _, want := range samples {
		rec := DoRequest(t, want.input, path + "/" + want.userID, http.MethodPut, EchoServer, want.access)

		t.Logf("rec.Code: %d", rec.Code)
		if want.err == nil {
			if !assert.Equal(t, http.StatusCreated, rec.Code) {
				t.Errorf("Body: %s", string(rec.Body.Bytes()))
			}

			var data test_models.TestGetUserUpdateResp
			json.Unmarshal(rec.Body.Bytes(), &data)
			assert.Equal(t, respStruct, data)
		} else {
			assert.Equal(t, want.err.Code, rec.Code)
			errResp, err := json.Marshal(want.err.Message)
			assert.Nil(t, err)

			type Err struct {
				Message string `json:"message"`
			}

			var errMsg Err
			err = json.Unmarshal(rec.Body.Bytes(), &errMsg)
			if err != nil {
				log.Print(err)
			}
			message, _ := json.Marshal(errMsg.Message)

			assert.Equal(t, string(errResp), string(message))
		}
	}
}

func TestDeleteUser(t *testing.T) {
	access, _ := AuthPrepare(t, "vasya.pupkin@gmail.com", "SuPeRpassWORD124214", EchoServer)

	samples := []struct {
		access string
		userID string
		err    *echo.HTTPError
	}{
		{
			access: access,
			userID: "a2c4a3db-2397-494b-93ef-3da38be95a71",
			err:    nil,
		},
		{
			access: "",
			err: domain.ErrInvalidAccessToken,
		},
	}

	for _, want := range samples {
		respMessage := "User " + want.userID + " deleted"
		rec := DoRequest(t, nil, path + "/" + want.userID, http.MethodDelete, EchoServer, want.access)

		t.Logf("rec.Code: %d", rec.Code)
		if want.err == nil {
			if !assert.Equal(t, http.StatusOK, rec.Code) {
				t.Errorf("Body: %s", string(rec.Body.Bytes()))
			}

			var data string
			json.Unmarshal(rec.Body.Bytes(), &data)
			assert.Equal(t, respMessage, data)
		} else {
			assert.Equal(t, want.err.Code, rec.Code)
			errResp, err := json.Marshal(want.err.Message)
			assert.Nil(t, err)

			type Err struct {
				Message string `json:"message"`
			}

			var errMsg Err
			err = json.Unmarshal(rec.Body.Bytes(), &errMsg)
			if err != nil {
				log.Print(err)
			}
			message, _ := json.Marshal(errMsg.Message)

			assert.Equal(t, string(errResp), string(message))
		}
	}
}
