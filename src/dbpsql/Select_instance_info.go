package dbpsql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/masayoshi4649/GamePartySys/src/common"
)

// InstanceInfo ... Select_instance_info
type InstanceInfo struct {
	InstanceID         int
	RecruitType        int
	RecruitTypeName    string
	RecruitTitle       string
	RecruitPlace       string
	MaxJoin            int
	MinJoin            int
	RecruitDue         *time.Time
	JoinPassExist      bool
	RecruitCommentMust bool
	RecruitComment     sql.NullString
	PlayStyle          int
	DisplayUserName    string
	DisplayUserIcon    string
	CreateTime         *time.Time
}

// SelectInstanceInfo ... SELECT * FROM instance_info JOIN mst_recruit/d_info/user_config
func SelectInstanceInfo() []InstanceInfo {
	query, _ := common.ReadSQLFile("src/dbpsql/Select_instance_info.sql")

	conf := common.Config
	db, err := sql.Open("postgres", conf.Postgres.ConnectStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var allres []InstanceInfo
	var r InstanceInfo
	for rows.Next() {
		dberr := rows.Scan(
			&r.InstanceID,
			&r.RecruitType,
			&r.RecruitTypeName,
			&r.RecruitTitle,
			&r.RecruitPlace,
			&r.MaxJoin,
			&r.MinJoin,
			&r.RecruitDue,
			&r.JoinPassExist,
			&r.RecruitCommentMust,
			&r.RecruitComment,
			&r.PlayStyle,
			&r.DisplayUserName,
			&r.DisplayUserIcon,
			&r.CreateTime,
		)
		if dberr != nil {
			fmt.Println(dberr)
		}
		allres = append(allres, r)
	}
	return allres
}
