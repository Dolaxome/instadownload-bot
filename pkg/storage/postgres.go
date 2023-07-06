package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type value struct {
	chatID   int64
	dsuserid string
}

const updatevalds = "UPDATE values SET dsuserid=$2 WHERE chatid=$1"
const updatevalsess = "UPDATE values SET sessionid=$2 WHERE chatid=$1"
const updatevalcsrf = "UPDATE values SET csrftoken=$2 WHERE chatid=$1"

const initURow = "INSERT INTO values (chatid) VALUES ($1)"

const getRowds = "SELECT (dsuserid) FROM values WHERE chatid=$1"
const getRowsess = "SELECT (sessionid) FROM values WHERE chatid=$1"
const getRowcsrf = "SELECT (csrftoken) FROM values WHERE chatid=$1"

func DbAddDsID(db *sql.DB, chatID int64, dsuserid string) error {
	_, err := db.Exec(updatevalds, chatID, dsuserid)

	fmt.Println(chatID, dsuserid)
	if err != nil {
		return err
	}
	return err
}
func DbAddCsrf(db *sql.DB, chatID int64, csrftoken string) error {
	_, err := db.Exec(updatevalcsrf, chatID, csrftoken)

	fmt.Println(chatID, csrftoken)
	if err != nil {
		return err
	}
	return err
}
func DbAddSession(db *sql.DB, chatID int64, sessionid string) error {
	_, err := db.Exec(updatevalsess, chatID, sessionid)

	fmt.Println(chatID, sessionid)
	if err != nil {
		return err
	}
	return err
}

func GetDS(db *sql.DB, chatID int64) (string, error) {
	var isNIL int
	err := db.QueryRow("SELECT COUNT(dsuserid) FROM values WHERE chatid=$1", chatID).Scan(&isNIL)
	if err != nil {
		return "", err
	}

	if isNIL != 0 {
		var ds sql.NullString
		err := db.QueryRow(getRowds, chatID).Scan(&ds)
		if err != nil {
			return "", err
		}

		name := " "
		if ds.Valid {
			name = ds.String
		}
		return name, err
	}

	return "ds_user_id is missing", err
}

func GetSessID(db *sql.DB, chatID int64) (string, error) {
	var isNIL int
	err := db.QueryRow("SELECT COUNT(sessionid) FROM values WHERE chatid=$1", chatID).Scan(&isNIL)
	if err != nil {
		return "", err
	}

	if isNIL != 0 {
		var sess sql.NullString
		err := db.QueryRow(getRowsess, chatID).Scan(&sess)
		if err != nil {
			return "", err
		}

		name := " "
		if sess.Valid {
			name = sess.String
		}
		return name, err
	}

	return "sessionid is missing", err
}

func GetCSRF(db *sql.DB, chatID int64) (string, error) {
	var isNIL int
	err := db.QueryRow("SELECT COUNT(csrftoken) FROM values WHERE chatid=$1", chatID).Scan(&isNIL)
	if err != nil {
		return "", err
	}

	if isNIL != 0 {
		var csrf sql.NullString
		err := db.QueryRow(getRowcsrf, chatID).Scan(&csrf)
		if err != nil {
			return "", err
		}

		name := " "
		if csrf.Valid {
			name = csrf.String
		}
		return name, err
	}

	return "csrftoken is missing", err
}

func InitUserRow(db *sql.DB, chatid int64) error {
	_, err := db.Exec(initURow, chatid)
	if err != nil {
		return err
	}
	return err
}
