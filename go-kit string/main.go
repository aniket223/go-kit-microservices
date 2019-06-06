package main

import(
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)
var ErrEmpty = errors.New("empty string") // ErrEmpty is returned when an input string is empty.
type StringService interface{
	Uppercase(string)(string,error)
	Count(string)(int)
}

//busines layer
// stringService is a concrete implementation of StringService
type stringService struct{}

func(stringService)Uppercase(s string)(string,error){
	if s==""{
		return"",ErrEmpty
	}
	return strings.ToUpper(s),nil
}

func(stringService)Count(s string)(int){
	return len(s)
}
// endpoint layer


// for each method, we define request and response structs
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint{
	return func(_ context.Context, request interface{}) (interface{}, error) {

		req:= request.(uppercaseRequest)
		v,err:= svc.Uppercase(req.S)
		if err!=nil{
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
		}
	}

	func makeCountEndpoint(svc StringService) endpoint.Endpoint{
		return func(_ context.Context, request interface{}) (interface{},error){
			req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
		}
	}

	//transport layer

	func decodeUppercaseRequest(_ context.Context, req*http.Request) (interface{},error){
		var request uppercaseRequest
		if err:= json.NewDecoder(req.Body).Decode(&request); err!=nil{
			return nil, err
		}
		return request,nil
	}

	func decodeCounteRequest(_ context.Context, req*http.Request) (interface{},error){
		var request uppercaseRequest
		if err:= json.NewDecoder(req.Body).Decode(&request); err!=nil{
			return nil, err
		}
		return request,nil
	}

	func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
		return json.NewEncoder(w).Encode(response)
	}

//Entry Point

func main(){
	svc:= stringService()

	uppercaseHandler:=httptransport.NewServer{
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	}

	countHandler:=httptransport.NewServer{
		makeCountEndpoint(svc),
		decodeCounteRequest,
		encodeResponse,
	}

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}






