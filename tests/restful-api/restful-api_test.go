package tests

import (
	"fmt"
	"testing"

	"github.com/karma123321/go-api-tests/src/constants"
	"github.com/karma123321/go-api-tests/src/services"
	restfulapistructs "github.com/karma123321/go-api-tests/src/structs/restful-api-structs"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type RestfulApiSuite struct {
	suite.Suite
	service                   *services.RestfulApiService
	ParamGetObjectById        []*restfulapistructs.Object
	ParamGetNonExistentObject []string
}

func (s *RestfulApiSuite) BeforeAll(t provider.T) {
	s.service = &services.RestfulApiService{}
	s.service.Initialize()

	s.ParamGetObjectById = []*restfulapistructs.Object{
		{
			Id:   "1",
			Name: "Google Pixel 6 Pro",
			Data: restfulapistructs.Data{
				Color:    "Cloudy White",
				Capacity: "128 GB",
			},
		},
		{
			Id:   "2",
			Name: "Apple iPhone 12 Mini, 256GB, Blue",
			Data: restfulapistructs.Data{},
		},
	}

	s.ParamGetNonExistentObject = []string{"1000", "-1"}
}

func (s *RestfulApiSuite) TestGetObjects(t provider.T) {
	t.Skip("Skipped because response contains fields with wrong types")
	t.Title("GET /objects returns a list of 13 objects")
	t.Tags(constants.TagPositive, constants.TagGetObjectsMethod)
	t.Severity(allure.CRITICAL)

	_, list := s.service.GetObjects(&t)

	t.Require().Equal(13, len(list))
}

func (s *RestfulApiSuite) TableTestGetObjectById(t provider.T, expected *restfulapistructs.Object) {
	t.Titlef("GET /object/%s returns a specific object", expected.Id)
	t.Tags(constants.TagPositive, constants.TagGetObjectByIdMethod)
	t.Severity(allure.CRITICAL)

	_, actual := s.service.GetObjectById(&t, expected.Id)

	t.Require().Equal(*expected, *actual, "Objects should be equal")
}

func (s *RestfulApiSuite) TableTestGetNonExistentObject(t provider.T, objectId string) {
	t.Titlef("GET /object/%s returns the 404 error because of non-existent object", objectId)
	t.Tags(constants.TagNegative, constants.TagGetObjectByIdMethod)
	t.Severity(allure.CRITICAL)

	expectedError := fmt.Sprintf("{\"error\":\"Oject with id=%s was not found.\"}", objectId)

	response := s.service.Client.GetObjectById(&t, objectId)

	t.Assert().Equal(404, response.StatusCode())
	t.Require().Equal(expectedError, response.String())
}

func TestSuiteRunner(t *testing.T) {
	t.Parallel()

	suite.RunSuite(t, new(RestfulApiSuite))
}
