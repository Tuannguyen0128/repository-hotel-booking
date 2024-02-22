package repository

import (
	"database/sql"
	"repository-hotel-booking/internal/app/repository/account"
)

type (
	Repositories struct {
		AccountRepo *account.Repository
	}
)

func New(db *sql.DB) *Repositories {
	accountRepo := account.New(db)

	return &Repositories{
		AccountRepo: accountRepo,
	}
}

//func (r *Repositories) Close() {
//	if r.pfDb != nil {
//		_ = r.pfDb.Close()
//	}
//	if r.txbDb != nil {
//		_ = r.txbDb.Close()
//	}
//}
