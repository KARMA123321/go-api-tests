package helpers

import (
	"encoding/json"

	"github.com/karma123321/go-api-tests/src/structs/common"
	"github.com/ozontech/allure-go/pkg/allure"
	"resty.dev/v3"
)

type IAllureContext interface {
	WithNewAttachment(name string, mimeType allure.MimeType, content []byte)
}

func AttachRequestDataToReport(ctx IAllureContext, response *resty.Response) error {
	reqAttachment := common.Request{
		Url:     response.Request.URL,
		Method:  response.Request.Method,
		Headers: response.Request.Header,
		Body:    response.Request.Body,
	}

	if reqAttachmentJson, err := json.Marshal(reqAttachment); err != nil {
		return err
	} else {
		ctx.WithNewAttachment("Request:", allure.JSON, reqAttachmentJson)
	}

	resAttachment := common.Response{
		StatusCode: response.StatusCode(),
		Status:     response.Status(),
		Headers:    response.Header(),
		Body:       response.Result(),
	}

	if resAttachmentJson, err := json.Marshal(resAttachment); err != nil {
		return err
	} else {
		ctx.WithNewAttachment("Response:", allure.JSON, resAttachmentJson)
	}

	return nil
}
