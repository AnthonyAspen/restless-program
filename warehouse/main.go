package main

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "net/http"
  "log"
  "strings"
  "strconv"
  "encoding/json"
  "time"
)
func connectToDataBase()(db *gorm.DB,err error){
  dsn := "host=localhost user=postgres password=123 dbname=warehouse port=5432 sslmode=disable"
  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil,err
  }
  return db,nil
}

  type Model struct {
 ID        uint       `gorm:"primary_key auto_increment:true;column:id" json:"-"`
 CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
 UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
 DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}
type Product struct {
  Model    // changed gorm.model
  Amount uint       `gorm:"column:amount" json:"amount"`

}
// a method to get product amount 
func getProductAmount(orderId uint)(product []*Product,err error){

 db,err := connectToDataBase()
  if err != nil {
    return nil,err
  }
  db.Model(&Product{}).Select("products.amount").Where("id=?",orderId).Scan(&product)
  if db.Error != nil {
    return nil, db.Error
  }
  return product,nil
}

func productHandler(w http.ResponseWriter, r *http.Request){
  url := r.URL.Path
  if url == "/"{
    http.Error(w,"You didn't write any product's id",http.StatusBadRequest)
    return 
  }

      tmp := strings.Trim(url,"/")
      orderIdConv,err := strconv.ParseUint(tmp,10,64)
      if err != nil{
        http.Error(w,"failed to convert URL Path in uint",http.StatusInternalServerError)
      }
      amountBytes,err := getProductAmount(uint(orderIdConv))
      amountBytesJson,err := json.Marshal(amountBytes)
    
      w.Write(amountBytesJson)
      return 
    

}
func main (){
  http.HandleFunc("/",productHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}


