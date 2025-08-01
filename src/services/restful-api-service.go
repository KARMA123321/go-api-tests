package services

import (
	"github.com/karma123321/go-api-tests/src/clients"
	restfulapistructs "github.com/karma123321/go-api-tests/src/structs/restful-api-structs"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"resty.dev/v3"
)

type RestfulApiService struct {
	client *clients.RestfulApiClient
}

func (s *RestfulApiService) Initialize() {
	s.client = &clients.RestfulApiClient{}

	s.client.Initialize()
}

func (s *RestfulApiService) GetObjects(t *provider.T) (*resty.Response, []*restfulapistructs.Object) {
	response := s.client.GetObjects(t)

	return response, response.Result().([]*restfulapistructs.Object)
}

func (s *RestfulApiService) GetObjectById(t *provider.T, objectId string) (*resty.Response, *restfulapistructs.Object) {
	response := s.client.GetObjectById(t, objectId)

	return response, response.Result().(*restfulapistructs.Object)
}
