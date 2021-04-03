package handlers 

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

type Hello struct {
  l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
  return &Hello{l}
}

func (h *Hello) ServerHttp(w http.ResponseWriter, r *http.Request){
  
  log.Println("hello")
  d, err := ioutil.ReadAll(r.Body)
  if err != nil{
    http.Error(w,"oops",http.StatusBadRequest)
  }
  fmt.Println(w,d)
}