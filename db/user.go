package db

import (
	mydb "cloud-storage/db/mysql"
	"fmt"
	"log"
)

type User struct {
	Username   string
	Email      string
	Phone      string
	SignupAt   string
	LastActive string
	Status     int
}

func GetTokenFromDB(username string) (string, error) {
	stmt, err := mydb.DBConn().Prepare(
		"SELECT user_token FROM tbl_user_token WHERE user_name=?;",
	)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer stmt.Close()

	r, err := stmt.Query(username)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	} else if r == nil {
		log.Fatalf("Username not found: %s", err.Error())
		return "", err
	}
	pRows := mydb.ParseRows(r)
	if len(pRows) > 0 {
		return string(pRows[0]["user_token"].([]byte)), nil
	}
	return "", err
}

func UserSignUp(username string, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare("INSERT IGNORE INTO tbl_user(`user_name`, `user_pwd`) VALUES(?, ?)")
	if err != nil {
		fmt.Println("Failed to prepare statement, err:", err)
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to execute statement, err:", err)
		return false
	}

	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("Failed to get rows affected, err:", err)
		return false
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)

	if rowsAffected > 0 {
		fmt.Println("User registered successfully")
		return true
	} else {
		fmt.Println("No rows were affected. Possibly a duplicate entry or other issue.")
		return false
	}
}

// verify user password
func UserSignIn(username string, encpwd string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"SELECT * FROM tbl_user WHERE user_name=? LIMIT 1",
	)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Fatalf(err.Error())
		return false
	} else if rows == nil {
		log.Fatalf("Username not found: %s", err.Error())
		return false
	}
	pRows := mydb.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	return false
}

// update user token in the databse
func UpdateToken(username string, token string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"REPLACE INTO tbl_user_token(`user_name`, `user_token`) VALUES(?, ?)",
	)
	if err != nil {
		log.Fatalln(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		log.Fatalf(err.Error())
		return false
	}
	return true

}

func GetUserInfo(username string) (User, error) {
	user := User{}

	stmt, err := mydb.DBConn().Prepare(
		"SELECT user_name, signup_at FROM tbl_user WHERE user_name=? LIMIT 1",
	)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
