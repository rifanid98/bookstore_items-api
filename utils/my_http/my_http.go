package my_http

import (
	"encoding/json"
	"net/http"

	resp "github.com/rifanid98/bookstore_helper-go/response"
)

type IMyHttp interface {
	ToJsonRestErr(w http.ResponseWriter, body *resp.RestErr)
	ToJsonRestResp(w http.ResponseWriter, body *resp.RestResp)
	ToJsonRest(w http.ResponseWriter, body interface{}, statusCode int)
}

type myHttp struct{}

var MyHttp IMyHttp = &myHttp{}

func (myHttp *myHttp) ToJsonRestErr(w http.ResponseWriter, body *resp.RestErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(body.StatusCode))
	json.NewEncoder(w).Encode(body)
}

func (myHttp *myHttp) ToJsonRestResp(w http.ResponseWriter, body *resp.RestResp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(body.StatusCode))
	json.NewEncoder(w).Encode(body)
}

func (myHttp *myHttp) ToJsonRest(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}
