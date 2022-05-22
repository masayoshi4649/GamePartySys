package dbpsql

import (
	"database/sql"
	"log"

	"github.com/masayoshi4649/GamePartySys/src/common"
)

// UpsertDToken ... INSERT d_token / UPDATE d_token
func UpsertDToken(
	id string,
	accessToken string,
	expiresIn int,
	refreshToken string,
	scope string,
	tokenType string) {
	query, err := common.ReadSQLFile("./src/dbpsql/Upsert_d_token.sql")
	if err != nil {
		log.Fatal(err)
	}

	conf := common.Config
	db, err := sql.Open("postgres", conf.Postgres.ConnectStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, dberr := db.Exec(query,
		id,
		accessToken,
		expiresIn,
		refreshToken,
		scope,
		tokenType,
	)
	if dberr != nil {
		panic(dberr)
	}
}
