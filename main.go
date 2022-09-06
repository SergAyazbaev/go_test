package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "go_test/db"

	"database/sql"
	_ "github.com/lib/pq"
)


///
func main() {

	handleFunc()
}

func handleFunc() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	
	fmt.Println(t)
	

	t.ExecuteTemplate(w, "create", nil)
}


var db *sql.DB


func save_article(w http.ResponseWriter, r *http.Request) {
	var e error

	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	fmt.Println(title)
	fmt.Println(anons)
	fmt.Println(full_text)


	// Create client
	const dbstr = `host=127.0.0.1 port=5432 user=test password=qwerty dbname=test_test`
	// const dbstr = `host=127.0.0.1 port=5432 user=test password=qwerty dbname=test_test sslmode=disable TABLESPACE=pg_default`
	db, e = sql.Open("postgres", dbstr)
	if e != nil {
		fmt.Println(e)
		return 
	}
	fmt.Println( "DB Ready ")
			
	_, err := db.Exec(`
		INSERT INTO "spr_wh_hi" (name)  VALUES ('`+ fmt.Sprintf(title) +`') ;
		INSERT INTO "spr_wh_hi2" (name)  VALUES ('`+  fmt.Sprintf(anons) +`') ;
		INSERT INTO "spr_wh_lo2" (name)  VALUES ('`+  fmt.Sprintf(full_text) +`') ;
		INSERT INTO categories(name) VALUES ('Beverages');
	`)

	if err != nil {
		fmt.Println(err) //		
		// http.Redirect(w, r, "/create", http.StatusSeeOther)
		// return 
	}
	fmt.Println(" Saved into DB")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
