package main //required of any executable program

import ( //our Go packages for this project
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// where our program begins
func main() {
	//connect to the database
	var (
		server   = "oxypocserver"
		user     = "oxypocserver"
		password = "weofcv*498"
		port     = 1433
		database = "gaslift"
	)

	connString := fmt.Sprintf("server=%s.database.windows.net;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}

	//execute the query for ftp paths
	rows, err := db.Query("select id, processdatetime from Input_Output")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
	defer rows.Close()

	var (
		id              string
		processdatetime string
	)

	for rows.Next() {
		if err := rows.Scan(&id, &processdatetime); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, processdatetime)
	}

}
