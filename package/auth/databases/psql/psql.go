package psql

import (
	"askUs/v1/package/auth"
	"context"

	"gorm.io/gorm"
)

type AuthDb struct{}

func New() auth.DB {
	return &AuthDb{}
}
func (atb AuthDb) FindOrInsert(ctx context.Context, atr *auth.AuthRequest) (*auth.AuthRequest, error) {
	db := ctx.Value("pgClient").(*gorm.DB)
	tx := db.Where(&auth.AuthRequest{Email: atr.Email}).FirstOrCreate(atr)
	return atr, tx.Error
}

func (atb AuthDb) Update(ctx context.Context, atr *auth.AuthRequest) (*auth.AuthRequest, error) {
	// db := ctx.Value("pgClient").(*gorm.DB)
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Updates(atr)
	return atr, tx.Error
}

func (atb AuthDb) GetRequest(ctx context.Context, id int) (*auth.AuthRequest, error) {
	// db := ctx.Value("pgClient").(*gorm.DB)
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)

	var authReq = &auth.AuthRequest{
		ID: id,
	}
	tx := db.First(authReq)
	return authReq, tx.Error
}
