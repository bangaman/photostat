package user




import (
	"fmt"
	// "net/http"
	// "html/template"
	"db"
	// "log"
	// "os"
)


func allusers() string{
	track := []string{}
	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("DROP TABLE notification")
	//"CREATE TABLE IF NOT EXISTS users (usersid BIGSERIAL NOT NULL PRIMARY KEY,  username VARCHAR(50) NOT NULL, email VARCHAR(50) NOT NULL,password VARCHAR(100) NOT NULL, date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)")
	
	if err != nil {
		panic(err.Error())
		track = append(track, "error")
		
	}else{
		track = append(track, " Dropper deeni")
	}


	fmt.Println(result)

	return track[0]

}


func status() string{
	track := []string{}
	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("CREATE TABLE IF NOT EXISTS status ( statusid BIGSERIAL NOT NULL PRIMARY KEY, usersid TEXT NOT NULL, image TEXT NOT NULL, tag1 TEXT NOT NULL, text1 TEXT NOT NULL, tag2 TEXT NOT NULL, text2 TEXT NOT NULL, tag3 TEXT NOT NULL, text3 TEXT NOT NULL, statusvalue VARCHAR(50) NOT NULL, statusdate TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)")
	//"CREATE TABLE IF NOT EXISTS users (usersid BIGSERIAL NOT NULL PRIMARY KEY,  username VARCHAR(50) NOT NULL, email VARCHAR(50) NOT NULL,password VARCHAR(100) NOT NULL, date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)")
	
	if err != nil {
		panic(err.Error())
		track = append(track, "error")
		
	}else{
		track = append(track, " Dropper deeni")
	}


	fmt.Println(result)

	return track[0]

}



func comments() string{
	track := []string{}
	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("CREATE TABLE IF NOT EXISTS comments ( id BIGSERIAL NOT NULL PRIMARY KEY, usersid TEXT NOT NULL, statusid INT, comment TEXT, commentdate TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)")
	//"CREATE TABLE IF NOT EXISTS users (usersid BIGSERIAL NOT NULL PRIMARY KEY,  username VARCHAR(50) NOT NULL, email VARCHAR(50) NOT NULL,password VARCHAR(100) NOT NULL, date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)")
	
	if err != nil {
		panic(err.Error())
		track = append(track, "error")
		
	}else{
		track = append(track, " Dropper deeni")
	}


	fmt.Println(result)

	return track[0]

}



func like() string{
	track := []string{}
	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("CREATE TABLE IF NOT EXISTS statuslike (id BIGSERIAL NOT NULL PRIMARY KEY, mainuser TEXT NOT NULL, statusid INT)")
	//"CREATE TABLE IF NOT EXISTS users (usersid BIGSERIAL NOT NULL PRIMARY KEY,  username VARCHAR(50) NOT NULL, email VARCHAR(50) NOT NULL,password VARCHAR(100) NOT NULL, date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)")
	
	if err != nil {
		panic(err.Error())
		track = append(track, "error")
		
	}else{
		track = append(track, " Dropper deeni")
	}


	fmt.Println(result)

	return track[0]

}





func notify() string{
	track := []string{}
	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	//08166054078
	defer db.Close()

	result, err := db.Query("CREATE TABLE IF NOT EXISTS notification ( notifyid BIGSERIAL NOT NULL PRIMARY KEY, usersid TEXT NOT NULL, notifyvalue TEXT NOT NULL, checked VARCHAR(50) NOT NULL, statusdate TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)")
	//"CREATE TABLE IF NOT EXISTS users (usersid BIGSERIAL NOT NULL PRIMARY KEY,  username VARCHAR(50) NOT NULL, email VARCHAR(50) NOT NULL,password VARCHAR(100) NOT NULL, date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)")
	
	if err != nil {
		panic(err.Error())
		track = append(track, "error")
		
	}else{
		track = append(track, " Dropper deeni")
	}


	fmt.Println(result)

	return track[0]

}