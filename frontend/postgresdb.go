package main

import(
	"database/sql"
	"errors"
	_ "github.com/lib/pq" // driver for postgresql database
)

func (a *App) ValidateUserLogin(email, password string)(bool, error){
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
	fmt.Println("Now we are going to create the table for the users")

	if &a.Config.UserTable == nil || a.Config.UserTable == ""{
		return false, errors.New("Table in configuation unexpectadly found to be blank")
	}

	emailQuery := "select email from " + a.Config.UserTable
	emailQuery += " where email=" + username

	result, err := db.Exec(userQuery)
	if err != nil{
		return false, err
	}else if results != email{
		return false, err
	}

	passwordQuery := "select password from " + a.Config.UserTable
	passwordQuery := " where user=" + username

	result, err = db.Exec(passwordQuery);
	if err != nil{
		return false, err
	}else if result != password{
		return false, nil
	}

	return true, nil
}

func HashPassword(password string)(string, error){

}
