package main 

import (
  "github.com/squeezesky/restless-program/data-base/handlers"
  "log"
  "net/http"
  "os"
)

func main(){
  l := log.New(os.Stdout,"product-api",log.LstdFlags)

  hh := handlers.NewHello(l)
  
  sm := http.NewServeMux()
  sm.Handle("/",hh)

  http.ListenAndServe(":8080",nil)

}
