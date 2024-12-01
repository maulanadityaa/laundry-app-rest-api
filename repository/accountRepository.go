package repository

import "maulanadityaa/laundry-app-rest-api/model/entity"

type AccountRepository interface {
	AddAccount(account entity.Account) (entity.Account, error)
	UpdateAccount(account entity.Account) (entity.Account, error)
	GetAccountByID(accountID string) (entity.Account, error)
	GetAccountByEmail(email string) (entity.Account, error)
}
