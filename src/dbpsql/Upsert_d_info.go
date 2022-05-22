package dbpsql

import (
	"database/sql"
	"log"

	"github.com/masayoshi4649/GamePartySys/src/common"
)

// UpsertDInfo ... INSERT d_info / UPDATE d_info
func UpsertDInfo(
	id string,
	username string,
	avatar string,
	discriminator string,
	publicFlags int,
	flags int,
	email string,
	verified bool,
	locale string,
	mfaEnabled bool) {
	query, err := common.ReadSQLFile("src/dbpsql/Upsert_d_info.sql")
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
		username,
		avatar,
		discriminator,
		publicFlags,
		flags,
		email,
		verified,
		locale,
		mfaEnabled,
	)

	if dberr != nil {
		panic(dberr)
	}
}
