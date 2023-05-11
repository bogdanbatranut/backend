package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	dbmodels "github.com/neoxelox/microservice-template/pkg/tenant_dbmodels"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getInvoicesHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func getInvoices(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var invoices []dbmodels.InvoiceMaster
		db.Preload("Customer").Preload("Details.Article").Find(&invoices)
		// db.Preload(clause.Associations).Find(&invoices)
		err := json.NewEncoder(w).Encode(invoices)
		if err != nil {
			panic(err)
		}
	}
}

func main() {

	log.Println("Started HTTP service")

	router := mux.NewRouter().StrictSlash(true)

	db := initGORM()

	router.HandleFunc("/invoices", getInvoices(db)).Methods("GET")

	httpPort := 8090
	log.Println("Listening on port ", httpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router)
	if err != nil {
		panic(err)
	}
}

func initGORM() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/15076503?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Created db connection...")
	return db
}
