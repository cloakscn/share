package models

type Greet struct {
	// database model ...
}

type GreetReq struct {
	Name string `url:"name"`
}

type GreetRes struct {
	Message string `json:"msg"`
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}
