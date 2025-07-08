package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"perfect-numbers-api/internal/handlers"
	"perfect-numbers-api/internal/models"
	"perfect-numbers-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	service := services.NewPerfectNumberService()
	handler := handlers.NewPerfectNumberHandler(service)

	router := gin.New()
	router.POST("/perfect-numbers", handler.FindPerfectNumbers)
	router.GET("/health", handler.Health)
	router.GET("/info", handler.APIInfo)

	return router
}

func TestPerfectNumberHandler_FindPerfectNumbers_Success(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name           string
		requestBody    models.PerfectNumberRequest
		expectedStatus int
		expectedResult []int
	}{
		{
			name:           "Range 1-10 deve retornar [6]",
			requestBody:    models.PerfectNumberRequest{Start: 1, End: 10},
			expectedStatus: http.StatusOK,
			expectedResult: []int{6},
		},
		{
			name:           "Range 1-100 deve retornar [6, 28]",
			requestBody:    models.PerfectNumberRequest{Start: 1, End: 100},
			expectedStatus: http.StatusOK,
			expectedResult: []int{6, 28},
		},
		{
			name:           "Range 1-10000 deve retornar [6, 28, 496, 8128]",
			requestBody:    models.PerfectNumberRequest{Start: 1, End: 10000},
			expectedStatus: http.StatusOK,
			expectedResult: []int{6, 28, 496, 8128},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/perfect-numbers", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response models.PerfectNumberResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResult, response.PerfectNumbers)
			assert.Equal(t, len(tt.expectedResult), response.Count)
			assert.NotEmpty(t, response.ProcessingTime)
		})
	}
}

func TestPerfectNumberHandler_FindPerfectNumbers_ValidationErrors(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Start maior que End deve retornar erro",
			requestBody:    models.PerfectNumberRequest{Start: 10, End: 5},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "start deve ser menor ou igual a end",
		},
		{
			name:           "Start negativo deve retornar erro",
			requestBody:    models.PerfectNumberRequest{Start: -1, End: 10},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Field validation for 'Start' failed on the 'min' tag",
		},
		{
			name:           "End muito grande deve retornar erro",
			requestBody:    models.PerfectNumberRequest{Start: 1, End: 2000000},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "end não pode ser maior que 1.000.000 para evitar timeout",
		},
		{
			name:           "JSON inválido deve retornar erro",
			requestBody:    `{"start": "invalid", "end": 10}`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "JSON inválido",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jsonBody []byte
			if str, ok := tt.requestBody.(string); ok {
				jsonBody = []byte(str)
			} else {
				jsonBody, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("POST", "/perfect-numbers", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response models.ErrorResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Contains(t, response.Error, tt.expectedError)
		})
	}
}

func TestPerfectNumberHandler_Health(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "OK", response.Status)
	assert.Equal(t, "Perfect Numbers API está funcionando", response.Message)
	assert.Equal(t, "2.0.0", response.Version)
}

func TestPerfectNumberHandler_APIInfo(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/info", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.APIInfoResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Perfect Numbers API", response.Name)
	assert.Equal(t, "2.0.0", response.Version)
	assert.NotEmpty(t, response.Endpoints)
}
