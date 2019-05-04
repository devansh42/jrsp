//@author Devansh Gupta
package jrsp

type ErrorResponse map[string]interface{}

//NewErrorResponse created new error response
func NewErrorResponse(errorcode uint, msg string) ErrorResponse {
	return map[string]interface{}{
		"code":   errorcode,
		"error":  true,
		"reason": msg}
}

//Response used to api request
type JsonResponse map[string]interface{}

//SetResponse sets response of the request
func (j JsonResponse) SetResponse(v interface{}) {
	j["resp"] = v
}

func (j JsonResponse) SetCode(c uint) {
	j["code"] = c
}
func (j JsonResponse) GetCode() int {
	return j["code"].(int)
}
func (j JsonResponse) SetErr(e ErrorResponse) {
	j["err"] = e
}

//Ok  Sets json reponse code to http.StatusOk
func (j JsonResponse) Ok() {
	j["code"] = 200
}
func (j JsonResponse) Free() {
	delete(j, "code")
	delete(j, "resp")
	delete(j, "err")
}
