package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type item struct {
	Name  string
	Price float32
}

var items = [3]item{
	{"OnePlus 10 Pro", 4500},
	{"Asus X513EA", 2000},
	{"Samsung UE55TU7022K", 1900},
}

func itemFunc(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/item/"):]) // pobranie ID przedmiotu

	// Sprawdzenie Id
	if id < 0 || id >= len(items) {
		http.Error(w, "Bledne ID", http.StatusBadRequest)
		return
	}

	// wyświetlenie informacji o produkcie na stronie
	tmpl, _ := template.ParseFiles("pages/items.html")
	tmpl.Execute(w, &item{items[id].Name, items[id].Price})

	// zapisywanie informacji do pliku logs.txt
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE, 0600)
	file.WriteString(fmt.Sprint(id) + "\t" + time.Now().String() + "\n")
	file.Close()
}

func main() {
	http.HandleFunc("/item/", itemFunc)
	http.ListenAndServe("localhost:8081", nil)
}
