package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"encoding/json"
	_ "test/go-sql-driver/mysql"
)

type Product struct {
	Id	int
	Name	string
	Price	int
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello !!!  %s", r.URL.Path[1:])
}

func listHandler(w http.ResponseWriter, r *http.Request){
	db := getConnection()

	fmt.Println("listHandler")

	var products []Product

	rows,err := db.Query("SELECT id ,name,price from product")

	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		var p Product
		rows.Scan(&p.Id,&p.Name,&p.Price)
		fmt.Println(p)
		products = append(products, p)
	}

	js, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/bignibou")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
        fmt.Println("application start")
        http.HandleFunc("/World", handler)
        http.HandleFunc("/productList", listHandler)
        http.ListenAndServe(":8080", nil)
}
