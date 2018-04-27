package response

import "encoding/json"

type Error struct {
	Message string
}

func (r *Error) Error() string {
	return r.Message
}

type Response struct {
	Error Error       `json:"error"`
	Data  interface{} `json:"data"`
}

func Create(data interface{}, errorMsg string) *Response {
	return &Response{Data: data, Error: Error{Message: errorMsg}}
}

func Json(data interface{}, errorMsg string) ([]byte, error) {
	r := Create(data, errorMsg)

	j, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	return j, err
}
