package db

import (
	"fmt"
	"database/sql"
	// _ "go_test/set"
	// _ "github.com/lib/pg"
	// _ "github.com/jackc/pgerrcode"
	_ "github.com/lib/pq"
)

func init() {
	fmt.Println(" DB")
}

type setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgBase     string
}

var db *sql.DB

var cfg setting

func Connect()  error {
	var e error

	const dbstr = `host=127.0.0.1 port=5432 user=test password=qwerty dbname=test_test`

	// const dbstr = `host=127.0.0.1 port=5432 user=test password=qwerty dbname=test_test sslmode=disable TABLESPACE=pg_default`

	db, e = sql.Open("postgres", dbstr)
	// db, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgBase))

	if e != nil {
		return  e
	}

	fmt.Println( "DB Ready ")
		
	// _, err := db.Exec(`
	// INSERT INTO "spr_wh_hi" (name)  VALUES ( 'go_test') ;
	// INSERT INTO "spr_wh_lo" (name,parent_id)  VALUES ( 'go_test',2 ) ;
	// INSERT INTO "spr_wh_hi2" (name)  VALUES ( 'go_test') ;
	// INSERT INTO "spr_wh_lo2" (name)  VALUES ( 'go_test') ;
	// INSERT INTO categories(name) VALUES ('Beverages');
	// `)

	// if err != nil {
	// 	fmt.Println(err) //
	// 	return e
	// }
	// fmt.Println(" Saved into DB")

	return nil

}


func CreateNew() error {
	var e error

	const dbstr = `host=127.0.0.1 port=5432 user=test password=qwerty dbname=test_test`

	db, e = sql.Open("postgres", dbstr)
	//  db, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgBase))

	if e != nil {
		return e
	}

	fmt.Println( "DB Ready ")
		
	_, err := db.Exec(`
	INSERT INTO "spr_wh_hi" (name)  VALUES ( 'go_test') ;
	INSERT INTO "spr_wh_lo" (name,parent_id)  VALUES ( 'go_test',2 ) ;
	INSERT INTO "spr_wh_hi2" (name)  VALUES ( 'go_test') ;
	INSERT INTO "spr_wh_lo2" (name)  VALUES ( 'go_test') ;
	INSERT INTO categories(name) VALUES ('Beverages');
	`)

	if err != nil {
		fmt.Println(err) //
		return e
	}
	fmt.Println(" Saved into DB")

	return nil

}

