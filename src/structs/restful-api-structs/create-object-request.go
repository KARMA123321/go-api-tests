package restfulapistructs

type CreateObjectRequestBody struct {
	Name string `json:"name"`
	Data Data   `json:"data"`
}

type CreateObjectResponseBody struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Data      Data   `json:"data"`
	CreatedAt string `json:"createdAt"`
}
