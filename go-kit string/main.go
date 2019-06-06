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
var ErrEmpty = errors.New("empty string")
type StringService interface{
	Uppercase(string)(string,error)
	Count(string)(int)
}

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

