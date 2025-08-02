package restfulapistructs

type Data struct {
	Year         int     `json:"year"`
	Price        float32 `json:"price"`
	CPUModel     string  `json:"CPU model"`
	HardDiskSize string  `json:"Hard disk size"`
	Color        string  `json:"color"`
	Capacity     string  `json:"capacity"`
	CapacityGB   int     `json:"capacity GB"`
	Generation   string  `json:"Generation"`
	StrapColour  string  `json:"Strap Color"`
	Description  string  `json:"Description"`
	ScreenSize   float32 `json:"Screen size"`
}
