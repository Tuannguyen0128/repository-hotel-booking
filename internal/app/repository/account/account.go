package account

import (
	"database/sql"
	"fmt"
	"log"
	"repository-hotel-booking/internal/app/model"
	"repository-hotel-booking/internal/app/repository/id_info"
	"repository-hotel-booking/internal/app/util"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAccounts(q *model.AccountQuery) ([]model.Account, *model.ErrInfo) {
	result := []model.Account{}
	query := `SELECT * FROM ACCOUNT WHERE DELETED_AT IS NULL`
	if q.ID != "" {
		query += " AND "
		query += fmt.Sprintf(`ID = '%s'`, q.ID)
	}
	if q.StaffID != "" {
		query += " AND "
		query += fmt.Sprintf(`STAFF_ID = '%s'`, q.StaffID)
	}
	if q.Username != "" {
		query += " AND "
		query += fmt.Sprintf(`USERNAME = '%s'`, q.Username)
	}

	if q.Page != 0 {
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", q.Size, (q.Page-1)*q.Size)
	}

	fmt.Println(query)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, util.BuildErrInfo("E01", err.Error())
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}(stmt)
	rows, err := stmt.Query()
	if err != nil {
		return nil, util.BuildErrInfo("E01", err.Error())
	}
	for rows.Next() {
		acc := model.Account{}
		err := rows.Scan(&acc.ID, &acc.StaffID, &acc.Username, &acc.Password,
			&acc.UserRoleID, &acc.CreatedAt, &acc.UpdatedAt, &acc.DeletedAt,
			&acc.LastLoginAt)
		if err != nil {
			return nil, util.BuildErrInfo("E01", err.Error())
		}
		result = append(result, acc)
	}
	return result, util.BuildErrInfo("", "")
}

func (r *Repository) InsertAccount(a *model.Account) (string, *model.ErrInfo) {
	IDInfo := id_info.GetIDInfo(r.db, "ACCOUNT")
	newID := id_info.GetNewID(IDInfo)
	query := "INSERT INTO ACCOUNT(`ID`,`STAFF_ID`,`USERNAME`,`PASSWORD`,`USER_ROLE_ID`) VALUES " +
		"(?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return newID, util.BuildErrInfo("E01", err.Error())
	}
	_, err = stmt.Exec(newID, a.StaffID, a.Username, a.Password, a.UserRoleID)
	if err != nil {
		fmt.Println(err.Error())
		return newID, util.BuildErrInfo("E01", err.Error())
	}

	err = id_info.IncreaseID(r.db, "ACCOUNT")
	if err != nil {
		return "", util.BuildErrInfo("E01", err.Error())
	}

	return newID, util.BuildErrInfo("", "")
}
