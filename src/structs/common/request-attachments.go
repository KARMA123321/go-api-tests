package common

type Request struct {
	Url     string
	Method  string
	Headers map[string][]string
	Body    any
}

type Response struct {
	StatusCode int
	Status     string
	Headers    map[string][]string
	Body       any
}
