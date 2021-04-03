package main

import (
  "log"
  "net/http"
  "encoding/json"
  "strconv"
  "strings"
)

type Handler struct {
  l *log.Logger
}

func NewHandler(l *log.Logger) *Handler{
  return &Handler{l}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
  if r.Method == http.MethodGet{
    url := r.URL.Path
    if url == "/"{
    // if url doesn't contain any id then get every order
     getEveryOrderHandler(w,r)
    }  else {
    // otherwise get an order info by id
      tmp := strings.Trim(url,"/")
      orderIdConv,err := strconv.ParseUint(tmp,10,64)
      if err != nil{
        http.Error(w,"failed to convert URL Path in uint",http.StatusInternalServerError)
      }
      getOrderByIdHandler(uint(orderIdConv),w,r)
      return
    }
  }
  if r.Method == http.MethodPost{
    // the method won't do anything, because I don't have to create a new order 
      http.Error(w,"there is nothing that you can do with a post method",http.StatusBadRequest)
      return
  }
  if r.Method == http.MethodPut{
    //TODO code to update data in my orders
    return
  }
  
  if r.Method == http.MethodDelete{
    //TODO code to delete an order
    return
  }
     
  
}

func getOrderByIdHandler(OrderId uint,w http.ResponseWriter, r *http.Request){
     lp,err := getInfoOrderById(OrderId)
     if err != nil{
        http.Error(w,"failed to get Order Info by Id",http.StatusInternalServerError)
     }
     lp_json,err := json.Marshal(lp)
     if err != nil{
         http.Error(w,"failed to marshal json",http.StatusInternalServerError)
      }
      w.Write(lp_json)
}
// handler to get information about every order 
func getEveryOrderHandler(w http.ResponseWriter, r *http.Request){
  lp,err := getProducts()
  if err != nil{
    http.Error(w,"Failed to get every order",http.StatusInternalServerError)
  }
  lp_json,err := json.Marshal(lp)
  if err != nil{
    http.Error(w,"failed to marshal json",http.StatusInternalServerError)
  }
  w.Write(lp_json)

}
