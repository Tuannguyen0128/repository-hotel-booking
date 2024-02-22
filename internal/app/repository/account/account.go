package account

import (
	"database/sql"
	"fmt"
	"repository-hotel-booking/internal/app/model"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAccounts(q *model.AccountQuery) ([]model.Account, error) {
	result := []model.Account{}
	query := `SELECT * FROM ACCOUNT WHERE `
	notEmpty := false
	if q.ID != "" {
		query += fmt.Sprintf(`ID = '%s'`, q.ID)
		notEmpty = true
	}
	if q.StaffID != "" {
		if notEmpty {
			query += fmt.Sprintf(" AND ")
		}
		query += fmt.Sprintf(`STAFF_ID = '%s'`, q.StaffID)
		notEmpty = true
	}
	if q.Username != "" {
		if notEmpty {
			query += fmt.Sprintf(" AND ")
		}
		query += fmt.Sprintf(`USERNAME = '%s'`, q.Username)
		notEmpty = true
	}

	if !notEmpty {
		query += fmt.Sprintf(" TRUE")
	}

	if q.Page != 0 {
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", q.Size, (q.Page-1)*q.Size)
	}

	fmt.Println(query)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		acc := model.Account{}
		err := rows.Scan(&acc.ID, &acc.StaffID, &acc.Username, &acc.Password,
			&acc.UserRoleID, &acc.CreatedAt, &acc.UpdatedAt, &acc.DeletedAt,
			&acc.LastLoginAt)
		if err != nil {
			return nil, err
		}
		result = append(result, acc)
	}
	return result, nil
}
