package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Operation struct {
	Number1 float64
	OperationSign string
	Number2 float64
	Result float64
}

var history []*Operation

func sum(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-TYpe", "application/json")

	parameters1 := mux.Vars(request)["number1"]
	parameters2 := mux.Vars(request)["number2"]

	number1, err := strconv.ParseFloat(parameters1, 64)
	number2, err := strconv.ParseFloat(parameters2, 64)

	if err == nil{
		fmt.Print(err)
	}

	result := number1 + number2

	sum := new(Operation)
	sum.Number1 = number1
	sum.OperationSign = "+"
	sum.Number2 = number2
	sum.Result = result

	sumJson, err := json.Marshal(sum)

	if err != nil {
		fmt.Println(err)
	}

	history = append(history, sum)

	response.Write(sumJson)
}

func sub(response http.ResponseWriter, request *http.Request) {
	parameters1 := mux.Vars(request)["number1"]
	parameters2 := mux.Vars(request)["number2"]

	number1, err := strconv.ParseFloat(parameters1, 64)
	number2, err := strconv.ParseFloat(parameters2, 64)

	if err == nil{
		fmt.Print(err)
	}

	result := number1 - number2

	sub := new(Operation)
	sub.Number1 = number1
	sub.OperationSign = "-"
	sub.Number2 = number2
	sub.Result = result

	subJson, err := json.Marshal(sub)
	if err != nil {
		fmt.Println(err)
	}

	history = append(history, sub)

	response.Write(subJson)
}

func mul(response http.ResponseWriter, request *http.Request) {
	parameters1 := mux.Vars(request)["number1"]
	parameters2 := mux.Vars(request)["number2"]

	number1, err := strconv.ParseFloat(parameters1, 64)
	number2, err := strconv.ParseFloat(parameters2, 64)

	if err == nil{
		fmt.Print(err)
	}

	result := number1 * number2

	mul := new(Operation)
	mul.Number1 = number1
	mul.OperationSign = "*"
	mul.Number2 = number2
	mul.Result = result

	mulJson, err := json.Marshal(mul)
	if err != nil {
		fmt.Println(err)
	}

	history = append(history, mul)

	response.Write(mulJson)
}

func div(response http.ResponseWriter, request *http.Request) {
	parameters1 := mux.Vars(request)["number1"]
	parameters2 := mux.Vars(request)["number2"]

	number1, err := strconv.ParseFloat(parameters1, 64)
	number2, err := strconv.ParseFloat(parameters2, 64)

	result := number1 / number2

	if number2 == 0 {
		response.WriteHeader(http.StatusExpectationFailed)
	}

	div := new(Operation)
	div.Number1 = number1
	div.OperationSign = "/"
	div.Number2 = number2
	div.Result = result

	divJson, err := json.Marshal(div)
	if err != nil {
		fmt.Println(err)
	}

	history = append(history, div)

	response.Write(divJson)
}

func historic(response http.ResponseWriter, request *http.Request){
	jsonHistory, err := json.Marshal(history)
	if err == nil{
		fmt.Print(err)
	}
	response.Write(jsonHistory)
}

func main(){
	fmt.Print("Serving at http://0.0.0.0:5000")
	r := mux.NewRouter()
	calcApi := r.PathPrefix("/calc").Subrouter()
	calcApi.HandleFunc("/sum/{number1}/{number2}", sum)
	calcApi.HandleFunc("/sub/{number1}/{number2}", sub)
	calcApi.HandleFunc("/mul/{number1}/{number2}", mul)
	calcApi.HandleFunc("/div/{number1}/{number2}", div)
	calcApi.HandleFunc("/history", historic)
	http.ListenAndServe(":5000", r)
}