package dto

type SetDummyRequest struct {
	Id                  string `json:"id"`
	Method              string `json:"method"`
	Path                string `json:"path"`
	RequestBody         string `json:"request_body"`
	ResponseContentType string `json:"response_conten_type"`
	ResponseBody        string `json:"response_body"`
	ResponseCode        int    `json:"response_code"`
}
