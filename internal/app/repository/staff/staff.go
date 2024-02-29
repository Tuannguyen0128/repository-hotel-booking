package staff

import (
	"database/sql"
	"fmt"
	"log"
	"repository-hotel-booking/internal/app/model"
	"repository-hotel-booking/internal/app/repository/id_info"
	"repository-hotel-booking/internal/app/util"
	"time"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetStaffs(q *model.StaffQuery) ([]model.Staff, *model.ErrInfo) {
	var result []model.Staff
	query := `SELECT * FROM STAFF WHERE DELETED_AT IS NULL`
	if q.ID != "" {
		query += " AND "
		query += fmt.Sprintf(`ID = '%s'`, q.ID)
	}
	if q.Position != "" {
		query += " AND "
		query += fmt.Sprintf(`POSITION = '%s'`, q.Position)
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
		staff := model.Staff{}
		err := rows.Scan(&staff.ID, &staff.FirstName, &staff.LastName, &staff.Position,
			&staff.Salary, &staff.DateOfBirth, &staff.Phone, &staff.Email,
			&staff.StartDate, &staff.DeletedAt)
		if err != nil {
			return nil, util.BuildErrInfo("E01", err.Error())
		}
		result = append(result, staff)
	}
	return result, util.BuildErrInfo("", "")
}

func (r *Repository) InsertStaff(s *model.Staff) (string, *model.ErrInfo) {
	log.Println(*s)
	IDInfo := id_info.GetIDInfo(r.db, "STAFF")
	newID := id_info.GetNewID(IDInfo)
	query := "INSERT INTO STAFF(`ID`,`FIRST_NAME`,`LAST_NAME`,`POSITION`,`SALARY`,`DATE_OF_BIRTH`,`PHONE`,`EMAIL`,`START_DATE`) VALUES " +
		"(?,?,?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", util.BuildErrInfo("E01", err.Error())
	}
	_, err = stmt.Exec(newID, s.FirstName, s.LastName, s.Position, s.Salary, s.DateOfBirth, s.Phone, s.Email, s.StartDate)
	if err != nil {
		return "", util.BuildErrInfo("E01", err.Error())
	}

	err = id_info.IncreaseID(r.db, "STAFF")
	if err != nil {
		return "", util.BuildErrInfo("E01", err.Error())
	}

	return newID, util.BuildErrInfo("", "")
}

func (r *Repository) DeleteStaff(id string) (string, *model.ErrInfo) {
	query := "DELETE from STAFF where ID=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", util.BuildErrInfo("E01", err.Error())
	}
	result, err := stmt.Exec(id)
	log.Println(query)
	log.Println(result)
	if err != nil {
		return "", util.BuildErrInfo("E01", err.Error())
	}

	return "Deleted", util.BuildErrInfo("", "")
}

func (r *Repository) UpdateStaff(s *model.Staff) (*model.Staff, *model.ErrInfo) {
	log.Println(*s)
	query := "UPDATE STAFF SET FIRST_NAME=?, LAST_NAME=?, POSITION=?, SALARY=?, DATE_OF_BIRTH=?, PHONE=?, EMAIL=?, START_DATE=? WHERE ID=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, util.BuildErrInfo("E01", err.Error())
	}
	time.LoadLocation("Asia/Ho_Chi_Minh")
	_, err = stmt.Exec(s.FirstName, s.LastName, s.Position, s.Salary, s.DateOfBirth, s.Phone, s.Email, s.StartDate, s.ID)
	if err != nil {

		return nil, util.BuildErrInfo("E01", err.Error())
	}

	staff, errI := r.GetStaffs(&model.StaffQuery{ID: s.ID})
	if errI.Code != "" {
		return nil, util.BuildErrInfo("E01", err.Error())
	}
	return &staff[0], util.BuildErrInfo("", "")
}
