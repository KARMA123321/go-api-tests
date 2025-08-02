package restfulapistructs

type Object struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Data Data   `json:"data"`
}
