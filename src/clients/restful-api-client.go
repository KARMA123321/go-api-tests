package clients

import (
	"fmt"

	"github.com/karma123321/go-api-tests/src/helpers"
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

		if err := helpers.AttachRequestDataToReport(sCtx, res); err != nil {
			(*t).Fatalf("Failed to attach request data to report: %v", err)
		}

		result = res
	})

	return result
}

func (c *RestfulApiClient) GetObjectById(t *provider.T, objectId string) *resty.Response {
	var result *resty.Response

	(*t).WithNewStep(fmt.Sprintf("Make a GET request to /object/%s", objectId), func(sCtx provider.StepCtx) {
		res, err := c.req.SetResult(&restfulapistructs.Object{}).Get(fmt.Sprintf("/objects/%s", objectId))

		if err != nil {
			(*t).Fatalf("Failed to make GET request: %v", err)
		}

		if err := helpers.AttachRequestDataToReport(sCtx, res); err != nil {
			(*t).Fatalf("Failed to attach request data to report: %v", err)
		}

		result = res
	})

	return result
}

func (c *RestfulApiClient) CreateObject(t *provider.T, object restfulapistructs.CreateObjectRequestBody) *resty.Response {
	var result *resty.Response

	(*t).WithNewStep("Make a POST request to /objects", func(sCtx provider.StepCtx) {
		res, err := c.req.SetBody(object).SetResult(&restfulapistructs.CreateObjectResponseBody{}).Post("/objects")

		if err != nil {
			(*t).Fatalf("Failed to make POST request: %v", err)
		}

		if err := helpers.AttachRequestDataToReport(sCtx, res); err != nil {
			(*t).Fatalf("Failed to attach request data to report: %v", err)
		}
		
		result = res
	})

	return result
}
