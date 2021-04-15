package main

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "net/http"
  "log"
  "strings"
  "strconv"
  "encoding/json"
)
func connectToDataBase()(db *gorm.DB,err error){
  dsn := "host=localhost user=postgres password=123 dbname=warehouse port=5432 sslmode=disable"
  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil,err
  }
  return db,nil
}

type Product struct {
  ID     uint       `gorm:"primary_key auto_increment:true;column:id" json:"id"`
  Amount uint       `gorm:"column:amount" json:"amount"`

}
func getProductAmount(orderId uint)(amount []byte,err error){

  return amount,nil
}

func productHandler(w http.ResponseWriter, r *http.Request){
  url := r.URL.Path
  if url == "/"{
    http.Error(w,"You didn't write any product's id",http.StatusBadRequest)
  }

      tmp := strings.Trim(url,"/")
      orderIdConv,err := strconv.ParseUint(tmp,10,64)
      if err != nil{
        http.Error(w,"failed to convert URL Path in uint",http.StatusInternalServerError)
      }
      amountBytes,err := getProductAmount(uint(orderIdConv))
      w.Write(amountBytes)
    

}
func main (){
  http.HandleFunc("/",productHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

  
}


