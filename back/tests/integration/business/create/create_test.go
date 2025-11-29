package create

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alexinator1/sumb/back/internal/domain/business/api/v1/businessgenerated"
	"github.com/alexinator1/sumb/back/tests/integration"
	"github.com/alexinator1/sumb/back/tests/integration/helpers"
	"github.com/stretchr/testify/suite"
)

type BusinessTestSuite struct {
	integration.IntegrationTestSuite
	dbVerifier *DbVerifier
}

func (bs *BusinessTestSuite) SetupTest() {
	bs.dbVerifier = NewDbVerifier(bs.App.DB(), bs.Require())
	bs.ClearTables("business", "employee")
}

func (bs *BusinessTestSuite) TestCreateBusiness() {
	t := bs.T()
	require := bs.Require()

	httpRequestFile := getTestDataPath("create_request.http")
	
	req, err := helpers.ParseHttpRequestFromFile(httpRequestFile)
	if err != nil {
		t.Fatalf("Не удалось распарсить http запрос: %v", err)
	}

	w := httptest.NewRecorder()

	bs.App.Router().ServeHTTP(w, req)

	// Assert response
	require.Equal(http.StatusCreated, w.Code, w.Body.Bytes())

	var actualResponse businessgenerated.CreateBusinessResponse
	if err := json.Unmarshal(w.Body.Bytes(), &actualResponse); err != nil {
		t.Fatalf("Не удалось анмаршалировать ответ: %v. Body: %s", err, w.Body.String())
	}

	expectedResponseBody, err := os.ReadFile(getTestDataPath("expected_create_response.json"))
	if err != nil {
		t.Fatalf("Не удалось прочитать файл с ожидаемым ответом: %v", err)
	}

	var expectedResponse businessgenerated.CreateBusinessResponse
	if err := json.Unmarshal(expectedResponseBody, &expectedResponse); err != nil {
		t.Fatalf("Не удалось анмаршалировать ожидаемый ответ: %v", err)
	}

	require.Equal(expectedResponse.Data.Id, actualResponse.Data.Id)
	require.Equal(expectedResponse.Data.Name, actualResponse.Data.Name)
	require.Equal(expectedResponse.Data.Description, actualResponse.Data.Description)
	require.Equal(expectedResponse.Data.OwnerEmail, actualResponse.Data.OwnerEmail)
	require.Equal(expectedResponse.Data.OwnerFirstName, actualResponse.Data.OwnerFirstName)
	require.Equal(expectedResponse.Data.OwnerLastName, actualResponse.Data.OwnerLastName)
	require.Equal(expectedResponse.Data.OwnerMiddleName, actualResponse.Data.OwnerMiddleName)
	require.Equal(expectedResponse.Data.OwnerPhone, actualResponse.Data.OwnerPhone)
	require.Equal(expectedResponse.Data.IsWorking, actualResponse.Data.IsWorking)
	require.Equal(expectedResponse.Data.LogoId, actualResponse.Data.LogoId)
	require.Equal(expectedResponse.Data.DeletedAt, actualResponse.Data.DeletedAt)
	require.Equal(expectedResponse.Message, actualResponse.Message)

	require.NoError(bs.dbVerifier.VerifyBusinessAndOwner(expectedResponse), "Не удалось проверить бизнес и владельца в базе данных")
}

func (bs *BusinessTestSuite) TestCreateBusinessError() {
	require := bs.Require()

	// Названия файлов тест кейсов для невалидных запросов и ожидаемых ответов
	testCases := []struct {
		reqFile  string
		respFile string
	}{
		{"validation/missing_required_fields.http", "validation/missing_required_fields_response.json"},
		{"validation/invalid_email.http", "validation/invalid_email_response.json"},
		{"validation/passwords_mismatch.http", "validation/password_mismatch_response.json"},
		{"error/empty.http", "error/empty_response.json"},
	}

	for _, tc := range testCases {
		httpRequestFile := getTestDataPath(tc.reqFile)

		req, err := helpers.ParseHttpRequestFromFile(httpRequestFile)
		require.NoErrorf(err, "Failed to parse http request from file data/%s: %v", tc.reqFile, err)

		expectedResponseBody, err := os.ReadFile(getTestDataPath(tc.respFile))
		require.NoErrorf(err, "Expected response file data/%s is missing: %v. Please create it manually with the expected validation error response.", tc.respFile, err)

		w := httptest.NewRecorder()

		bs.App.Router().ServeHTTP(w, req)

		require.Equal(http.StatusBadRequest, w.Code, "Expected status code 400 for validation error")

		var actualValidationError businessgenerated.ValidationErrorResponse
		err = json.Unmarshal(w.Body.Bytes(), &actualValidationError)
		require.NoError(err, "Failed to unmarshal error response")

		var expectedValidationError businessgenerated.ValidationErrorResponse
		err = json.Unmarshal(expectedResponseBody, &expectedValidationError)
		require.NoError(err, "Failed to unmarshal expected response")

		require.Equalf(expectedValidationError.Error, actualValidationError.Error, "Expected error field to match for %s", tc.reqFile)
		require.Equalf(expectedValidationError.Message, actualValidationError.Message, "Expected message field to match for %s", tc.reqFile)
		require.ElementsMatchf(expectedValidationError.Details, actualValidationError.Details, "Validation error details do not match for %s", tc.reqFile)

		bs.dbVerifier.VerifyEmptyTable("business")
		bs.dbVerifier.VerifyEmptyTable("employee")
	}
}
func TestBusinessSuite(t *testing.T) {
	suite.Run(t, new(BusinessTestSuite))
}

// getTestDataPath возвращает абсолютный путь к файлу данных теста
func getTestDataPath(filename string) string {
	_, testFile, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(testFile)
	return filepath.Join(testDir, "data", filename)
}