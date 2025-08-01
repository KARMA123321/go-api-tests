package services

import (
	"github.com/karma123321/go-api-tests/src/clients"
	restfulapistructs "github.com/karma123321/go-api-tests/src/structs/restful-api-structs"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"resty.dev/v3"
)

type RestfulApiService struct {
	Client *clients.RestfulApiClient
}

func (s *RestfulApiService) Initialize() {
	s.Client = &clients.RestfulApiClient{}

	s.Client.Initialize()
}

func (s *RestfulApiService) GetObjects(t *provider.T) (*resty.Response, []*restfulapistructs.Object) {
	response := s.Client.GetObjects(t)

	if response.IsError() {
		(*t).Fatalf("Failed to get objects: %v", response.Error())
	}

	return response, response.Result().([]*restfulapistructs.Object)
}

func (s *RestfulApiService) GetObjectById(t *provider.T, objectId string) (*resty.Response, *restfulapistructs.Object) {
	response := s.Client.GetObjectById(t, objectId)

	if response.IsError() {
		(*t).Fatalf("Failed to get object with ID '%s': %v", objectId, response.Error())
	}

	return response, response.Result().(*restfulapistructs.Object)
}
