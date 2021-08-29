package db


import (
	"database/sql"
	_"github.com/lib/pq"
	"fmt"
	"regexp"
	"strings"
)


const (
	host = "******************************************************************"
	port = "5432"
	user = "******************************************************************"
	database = "******************************************************************"
	password = "******************************************************************"
	
)

var (

	connStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require", host, port, user, database, password)
)

func Conn() (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}



func Router(name, match string) []string {
	res1, _ := regexp.MatchString(match, name)

	final := []string{}

	if res1 == true {
		ma := strings.Split(name, match)
		
		for i, _ := range ma {
			if ma[i] != ""{
				final = append(final, strings.ToLower(strings.TrimSpace(ma[i])))
		
			}
	
		}

	}else{
		final = append(final, strings.ToLower(strings.TrimSpace(name)))

	}

	return final
	
}
