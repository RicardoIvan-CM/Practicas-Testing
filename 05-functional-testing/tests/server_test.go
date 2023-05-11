package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreyEscaped(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		server := createServer()

		//Arrange
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `{
			"success":	true
		}`

		//Configure Prey

		//Act
		request, respRecorder := createRequestTest(http.MethodPut, "/v1/prey", `{
			"speed" : 120
		}`)
		server.ServeHTTP(respRecorder, request)

		//Assert
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		assert.Equal(t, expectedHeaders, respRecorder.Header())
		assert.JSONEq(t, expectedBody, respRecorder.Body.String())

		//Configure Shark

		//Act
		request, respRecorder = createRequestTest(http.MethodPut, "/v1/shark", `{
			"x_position" : 200,
			"y_position" : 200,
			"speed" : 100
		}`)
		server.ServeHTTP(respRecorder, request)

		//Assert
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		assert.Equal(t, expectedHeaders, respRecorder.Header())
		assert.JSONEq(t, expectedBody, respRecorder.Body.String())

		//Simulation

		//Arrange
		expectedBody = `{
			"success":	false,
			"message": "could not catch it",
			"time": 0
		}`

		//Act
		request, respRecorder = createRequestTest(http.MethodPost, "/v1/simulate", ``)
		server.ServeHTTP(respRecorder, request)

		//Assert
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		assert.Equal(t, expectedHeaders, respRecorder.Header())
		assert.JSONEq(t, expectedBody, respRecorder.Body.String())
	})

}

func TestPreyTooFar(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		server := createServer()

		//Arrange
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `{
			"success":	true
		}`

		//Configure Prey

		//Act
		request, respRecorder := createRequestTest(http.MethodPut, "/v1/prey", `{
			"speed" : 10
		}`)
		server.ServeHTTP(respRecorder, request)

		//Assert
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		assert.Equal(t, expectedHeaders, respRecorder.Header())
		assert.JSONEq(t, expectedBody, respRecorder.Body.String())

		//Configure Shark

		//Act
		request, respRecorder = createRequestTest(http.MethodPut, "/v1/shark", `{
			"x_position" : 1000,
			"y_position" : 1000,
			"speed" : 30
		}`)
		server.ServeHTTP(respRecorder, request)

		//Assert
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		assert.Equal(t, expectedHeaders, respRecorder.Header())
		assert.JSONEq(t, expectedBody, respRecorder.Body.String())

		//Simulation

		//Arrange
		expectedBody = `{
			"success":	false,
			"message": "could not catch it",
			"time": 0
		}`

		//Act
		request, respRecorder = createRequestTest(http.MethodPost, "/v1/simulate", ``)
		server.ServeHTTP(respRecorder, request)

		//Assert
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		assert.Equal(t, expectedHeaders, respRecorder.Header())
		assert.JSONEq(t, expectedBody, respRecorder.Body.String())
	})
}

func TestCatchAfter24Secs(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		server := createServer()

		//Arrange
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `{
			"success":	true
		}`

		//Configure Prey

		//Act
		request, respRecorder := createRequestTest(http.MethodPut, "/v1/prey", `{
			"speed" : 2
		}`)
		server.ServeHTTP(respRecorder, request)

		//Assert
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		assert.Equal(t, expectedHeaders, respRecorder.Header())
		assert.JSONEq(t, expectedBody, respRecorder.Body.String())

		//Configure Shark

		//Act
		request, respRecorder = createRequestTest(http.MethodPut, "/v1/shark", `{
			"x_position" : 2,
			"y_position" : 3.4641,
			"speed" : 2.1666
		}`)
		server.ServeHTTP(respRecorder, request)

		//Assert
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		assert.Equal(t, expectedHeaders, respRecorder.Header())
		assert.JSONEq(t, expectedBody, respRecorder.Body.String())

		//Simulation

		//Arrange
		expectedBody = `{
			"success":	true,
			"message": "could catch it",
			"time": 24
		}`

		//Act
		request, respRecorder = createRequestTest(http.MethodPost, "/v1/simulate", ``)
		server.ServeHTTP(respRecorder, request)

		//Assert
		assert.Equal(t, http.StatusOK, respRecorder.Code)
		assert.Equal(t, expectedHeaders, respRecorder.Header())
		assert.JSONEq(t, expectedBody, respRecorder.Body.String())
	})
}

func TestFailConfigurePrey(t *testing.T) {
	server := createServer()

	//Arrange
	expectedHeaders := http.Header{
		"Content-Type": []string{"application/json; charset=utf-8"},
	}
	expectedBody := `{
		"success":	false
	}`

	//Act
	request, respRecorder := createRequestTest(http.MethodPut, "/v1/prey", `{
		"speed" : a
	}`)
	server.ServeHTTP(respRecorder, request)

	//Assert
	assert.Equal(t, http.StatusBadRequest, respRecorder.Code)
	assert.Equal(t, expectedHeaders, respRecorder.Header())
	assert.JSONEq(t, expectedBody, respRecorder.Body.String())
}

func TestFailConfigureShark(t *testing.T) {
	server := createServer()

	//Arrange
	expectedHeaders := http.Header{
		"Content-Type": []string{"application/json; charset=utf-8"},
	}
	expectedBody := `{
		"success":	false
	}`

	//Act
	request, respRecorder := createRequestTest(http.MethodPut, "/v1/shark", `{
		"x_position" : "2",
		"y_position" : true,
		"speed" : a
	}`)
	server.ServeHTTP(respRecorder, request)

	//Assert
	assert.Equal(t, http.StatusBadRequest, respRecorder.Code)
	assert.Equal(t, expectedHeaders, respRecorder.Header())
	assert.JSONEq(t, expectedBody, respRecorder.Body.String())
}
