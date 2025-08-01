package restfulapistructs

type Object struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Data Data   `json:"data"`
}

type Data struct {
	Year         int     `json:"year"`
	Price        float32 `json:"price"`
	CPUModel     string  `json:"CPU model"`
	HardDiskSize string  `json:"Hard disk size"`
	Color        string  `json:"color"`
	Capacity     string  `json:"capacity"`
}
