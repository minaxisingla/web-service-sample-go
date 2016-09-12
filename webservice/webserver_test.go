	package webservice

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockClient struct{
	mock.Mock
}

type HandlerTestSuite struct{
	suite.Suite
	formResponseRecorder *httptest.ResponseRecorder
	viewResponseRecorder *httptest.ResponseRecorder
}

func (suite *HandlerTestSuite) SetupTest() {
	suite.formResponseRecorder = httptest.NewRecorder()
	suite.viewResponseRecorder = httptest.NewRecorder()
}

func (suite *HandlerTestSuite) TestMakeFormHandler() {
	formHandle := MakeFormHandler()
	req, _ := http.NewRequest("GET", "http://localhost:8080/hello", nil)

	formHandle.ServeHTTP(suite.formResponseRecorder, req)
	assert.Equal(suite.T(), suite.formResponseRecorder.Code, http.StatusOK, "Page should return status OK")
}

func (suite *HandlerTestSuite) TestMakeViewHandler() {
	viewHandle := MakeViewHandler()
	req, _ := http.NewRequest("GET", "http://localhost:8080/submit?name=ABC", nil)

	viewHandle.ServeHTTP(suite.viewResponseRecorder, req)
	assert.Equal(suite.T(), suite.viewResponseRecorder.Code, http.StatusOK, "Page should return status OK")
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}