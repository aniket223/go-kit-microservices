package main
import(
	"encoding/json"
    "errors"
    "fmt"
    "net/http"
    "strings"
)
var ErrEmptyString = errors.New("Empty String") 

type stringRequest struct { 
	S string `json:"s"` // it will represent json as {s:""}
}

type countResponse struct {
	C int `json:"c"`
}

type uppercaseResponse struct {
	S string `json:"s"`
}

func main(){
	http.HandleFunc("/count",countHandler)
	http.HandleFunc("/uppercase",uppercaseHandler)
	http.ListenAndServe(":8081",nil)
}

func count(s string) int {  //for length of string
	return len(s)
}

func uppercase(s string) string { //for converting to uppercase 
	return strings.ToUpper(s)
}

func countHandler(w http.ResponseWriter, req *http.Request){
	var request stringRequest

	if err:=json.NewDecoder(req.Body).Decode(&request); err!=nil{
		fmt.Println(err)
	}

	if request.S == ""{
		json.NewEncoder(w).Encode(ErrEmptyString)
	}
	response:= &countResponse{
		C:count(request.S),
	}
json.NewEncoder(w).Encode(&response)
}

func uppercaseHandler(w http.ResponseWriter, req *http.Request){
	var request stringRequest

	if err:=json.NewDecoder(req.Body).Decode(&request); err!=nil{
		fmt.Println(err)
	}

	if request.S == ""{
		json.NewEncoder(w).Encode(ErrEmptyString)
	}
	response:= &uppercaseResponse{
		S:uppercase(request.S),
	}
json.NewEncoder(w).Encode(&response)
}
