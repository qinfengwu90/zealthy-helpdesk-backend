package dao

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"zealthy-helpdesk-backend/utility"
)

var DB *sqlx.DB

func DbInit(dbConfig *utility.PostgresInfo) {
	DBPortString := dbConfig.Port
	DBPort, _ := strconv.Atoi(DBPortString)

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		DBPort,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Dbname)

	db, err := sqlx.Connect("postgres", psqlConn)

	if err != nil {
		psqlVpcConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.VpcPrivateHost,
			DBPort,
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Dbname)

		db, err := sqlx.Connect("postgres", psqlVpcConn)
		if err != nil {
			log.Fatal(err)
		}
		DB = db
	} else {
		DB = db
	}
}
