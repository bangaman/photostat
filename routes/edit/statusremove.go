package edit




import(
	// "fmt"
	"db"
	_ "encoding/json"
	"log"
	"context"
	"time"
)

func DeleteStatus(statusid string) error {
	db, err := db.Conn()

	query := "DELETE FROM status WHERE statusid = $1"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	//defer the close till after all connections are cloed
	defer stmt.Close()

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, statusid)

	if err != nil {
		log.Printf("%s Error when deleting data !!", err)
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Printf("%s When finding rows affected !!", err)
		return err
	}

	log.Printf("%d Products deleted ", rows)

	return nil

}


func GetAllLikesRelatedToDeletedPost(id string) []string {
	var bat []string

	type GetUser struct{
		id string `json:"id"`
		mainuser string `json:"mainuser"`
		statusid string `json:"statuid"`
	}

	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("SELECT * FROM statuslike")

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag  GetUser

		err = result.Scan(&tag.id, &tag.mainuser, &tag.statusid)

		if err != nil {
			panic(err.Error())
		}


		if tag.statusid == id{
			bat = append(bat, tag.id)
		}
		
	}

	return bat
}





func GetAllCommentsRelatedToDeletedPost(id string) []string {
	var bat []string

	type GetUser struct{
		id string `json:"id"`
		statusid string `json:"statuid"`
	}

	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("SELECT id, statusid FROM comments")

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag  GetUser

		err = result.Scan(&tag.id, &tag.statusid)

		if err != nil {
			panic(err.Error())
		}


		if tag.statusid == id{
			bat = append(bat, tag.id)
		}
		
	}

	return bat
}



func DeleteLikes(statusid string) error {
	db, err := db.Conn()

	query := "DELETE FROM statuslike WHERE id = $1"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	//defer the close till after all connections are cloed
	defer stmt.Close()

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, statusid)

	if err != nil {
		log.Printf("%s Error when deleting data !!", err)
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Printf("%s When finding rows affected !!", err)
		return err
	}

	log.Printf("%d Products deleted ", rows)

	return nil

}


func DeleteComments(statusid string) error {
	db, err := db.Conn()

	query := "DELETE FROM comments WHERE id = $1"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	//defer the close till after all connections are cloed
	defer stmt.Close()

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, statusid)

	if err != nil {
		log.Printf("%s Error when deleting data !!", err)
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Printf("%s When finding rows affected !!", err)
		return err
	}

	log.Printf("%d Products deleted ", rows)

	return nil

}

