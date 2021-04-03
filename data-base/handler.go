package main

import (
  "log"
  "net/http"
  "encoding/json"
)

type Handler struct {
  l *log.Logger
}

func NewHandler(l *log.Logger) *Handler{
  return &Handler{l}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
  if r.Method == http.MethodGet{
    getHandler(w,r)
    return
  }
  if r.Method == http.MethodPost{

    //TODO code to add new data in my orders

  }
  if r.Method == http.MethodPut{
    //TODO code to update data in my orders
  }
  
}
func getHandler(w http.ResponseWriter, r *http.Request){
  lp,err := getProducts()
  if err != nil{
    http.Error(w,"an error",http.StatusInternalServerError)
  }
  lp_json,err := json.Marshal(lp)
  if err != nil{
    http.Error(w,"failed to marshal json",http.StatusInternalServerError)
  }
  w.Write(lp_json)

}
