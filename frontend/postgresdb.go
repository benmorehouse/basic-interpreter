package main

import(
	"database/sql"
	_ "github.com/lib/pq" // driver for postgresql database
)

func (a *App) ValidateUserLogin(username, password string )(bool, error){
	// this password needs to go through a hashing first
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func HashPassword(password string)(string, error){

}
