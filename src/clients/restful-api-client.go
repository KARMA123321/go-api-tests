package clients

import (
	"fmt"

	restfulapistructs "github.com/karma123321/go-api-tests/src/structs/restful-api-structs"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"resty.dev/v3"
)

type RestfulApiClient struct {
	client *resty.Client
	req    *resty.Request
}

func (c *RestfulApiClient) Initialize() {
	c.client = resty.New().SetBaseURL("https://api.restful-api.dev")
	c.req = c.client.R()
}

func (c *RestfulApiClient) GetObjects(t *provider.T) *resty.Response {
	var result *resty.Response

	(*t).WithNewStep("Make a GET request to /objects", func(sCtx provider.StepCtx) {
		res, err := c.req.SetResult(make([]restfulapistructs.Object, 0)).Get("/objects")

		if err != nil {
			(*t).Fatalf("Failed to make GET request: %v", err)
		}

		result = res
	})

	return result
}

func (c *RestfulApiClient) GetObjectById(t *provider.T, objectId string) *resty.Response {
	var result *resty.Response

	(*t).WithNewStep("Make a GET request to /object/{id}", func(sCtx provider.StepCtx) {
		res, err := c.req.SetResult(&restfulapistructs.Object{}).Get(fmt.Sprintf("/objects/%s", objectId))

		if err != nil {
			(*t).Fatalf("Failed to make GET request: %v", err)
		}

		result = res
	})

	return result

}
