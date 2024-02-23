package id_info

import (
	"database/sql"
	"fmt"
	"log"
	"repository-hotel-booking/internal/app/model"
	"strconv"
	"strings"
)

func GetIDInfo(db *sql.DB, name string) *model.IDInfo {
	result := &model.IDInfo{}
	lockQuery := "LOCK TABLES ID_INFO WRITE, ACCOUNT WRITE"
	_, err := db.Exec(lockQuery)
	if err != nil {
		return nil
	}

	query := "SELECT * FROM ID_INFO WHERE NAME = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	rows, err := stmt.Query(name)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	//unlockQuery := "UNLOCK TABLES ID_INFO"
	//_, err = db.Exec(unlockQuery)
	//if err != nil {
	//	return nil
	//}
	for rows.Next() {
		err = rows.Scan(&result.Table, &result.Prefix, &result.Length, &result.Current)
		if err != nil {
			return nil
		}
	}

	return result

}

func GetNewID(i *model.IDInfo) string {
	next := i.Current + 1
	nextLen := len(strconv.FormatInt(next, 10))
	var sb strings.Builder
	sb.Grow(i.Length)
	for k := 0; k < i.Length; k++ {
		if k == 0 {
			sb.Write([]byte(i.Prefix))
		} else if k+nextLen >= i.Length {
			sb.Write([]byte(strconv.FormatInt(next, 10)))
			break
		} else {
			sb.Write([]byte("0"))
		}
	}
	return sb.String()
}

func IncreaseID(db *sql.DB, name string) error {

	query := "UPDATE ID_INFO SET CURRENT = CURRENT + 1 WHERE NAME = ? "
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	_, err = stmt.Exec(name)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	unlockWrite := "UNLOCK TABLES"
	_, err = db.Exec(unlockWrite)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return nil

}

//func UnlockTables
