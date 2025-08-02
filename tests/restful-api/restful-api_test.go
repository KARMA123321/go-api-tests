package tests

import (
	"fmt"
	"strconv"
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

// Positive tests

func (s *RestfulApiSuite) TestGetObjects(t provider.T) {
	t.Skip("Skipped because response contains fields with wrong types")
	t.Title("GET /objects returns a list of 13 objects")
	t.Tags(constants.TagPositive, constants.TagGetObjectsMethod)
	t.Severity(allure.CRITICAL)

	response, list := s.service.GetObjects(&t)

	t.Assert().Equal(200, response.StatusCode())
	t.Assert().Equal(13, len(list))
}

func (s *RestfulApiSuite) TableTestGetObjectById(t provider.T, expected *restfulapistructs.Object) {
	t.Titlef("GET /object/%s returns a specific object", expected.Id)
	t.Tags(constants.TagPositive, constants.TagGetObjectByIdMethod)
	t.Severity(allure.CRITICAL)

	response, actual := s.service.GetObjectById(&t, expected.Id)

	t.Assert().Equal(200, response.StatusCode())
	t.Assert().Equal(*expected, *actual, "Objects should be equal")
}

func (s *RestfulApiSuite) TestCreateObject(t provider.T) {
	t.Skip("Skipped because response's Id field is not a valid integer and CreatedAt field is not a valid time")
	t.Title("POST /object/%s creates a new object with provided data")
	t.Tags(constants.TagPositive, constants.TagCreateObjectMethod)
	t.Severity(allure.CRITICAL)

	newObject := restfulapistructs.CreateObjectRequestBody{
		Name: "IPhone 17 Pro Max",
		Data: restfulapistructs.Data{
			Year:         2025,
			Color:        "Black",
			Capacity:     "1 TB",
			Price:        1000.00,
			CPUModel:     "Apple A17 Pro",
			HardDiskSize: "1 TB",
			CapacityGB:   1024,
			Generation:   "17",
			StrapColour:  "Black",
			Description:  "Latest iPhone with advanced features",
			ScreenSize:   6.7,
		},
	}

	response, body := s.service.CreateObject(&t, newObject)
	_, err := strconv.Atoi(body.Id)

	// We could consider 200 as a wrong status code here, but we won't
	t.Assert().Equal(200, response.StatusCode())
	t.Assert().NoError(err, "Id should be a valid integer")
	t.Assert().Equal(newObject.Name, body.Name)
	t.Assert().Equal(newObject.Data, body.Data)
	t.Assert().Regexp(constants.TimeRegexp, body.CreatedAt)
}

// Negative tests

func (s *RestfulApiSuite) TableTestGetNonExistentObject(t provider.T, objectId string) {
	t.Titlef("GET /object/%s returns the 404 error because of non-existent object", objectId)
	t.Tags(constants.TagNegative, constants.TagGetObjectByIdMethod)
	t.Severity(allure.MINOR)

	expectedError := fmt.Sprintf("{\"error\":\"Oject with id=%s was not found.\"}", objectId)

	response := s.service.Client.GetObjectById(&t, objectId)

	t.Assert().Equal(404, response.StatusCode())
	t.Require().Equal(expectedError, response.String())
}

func TestSuiteRunner(t *testing.T) {
	t.Parallel()

	suite.RunSuite(t, new(RestfulApiSuite))
}
