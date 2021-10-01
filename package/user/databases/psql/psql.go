package psql

import (
	"askUs/v1/package/user"
	"context"

	"gorm.io/gorm"
)

type UserDb struct{}

func (udb *UserDb) FindOrCreateDoctor(ctx context.Context, email string) (*user.Doctor, error) {
	// db := ctx.Value("pgClient").(*gorm.DB)
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Doctor{Info: user.Info{Email: email}, Connections: &user.Connections{
		ConnectedWith: "doctor",
	}}
	tx := db.Where("email=?", email).
		FirstOrCreate(doc)
	return doc, tx.Error
}
func (udb *UserDb) FindOrCreatePatient(ctx context.Context, email string) (*user.Patient, error) {
	// db := ctx.Value("pgClient").(*gorm.DB)
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Patient{Info: user.Info{Email: email}, Connections: &user.Connections{ConnectedWith: "patient"}}
	tx := db.Where("email=?", email).
		FirstOrCreate(doc)
	return doc, tx.Error
}
func (udb *UserDb) GetDoctorByID(ctx context.Context, id float64) (*user.Doctor, error) {
	// db := ctx.Value("pgClient").(*gorm.DB)
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Doctor{}
	tx := db.First(doc, "id=?", id)
	return doc, tx.Error
}
func (udb *UserDb) GetDoctorByName(ctx context.Context, name string) (*user.Doctor, error) {
	// db := ctx.Value("pgClient").(*gorm.DB)
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Doctor{}
	tx := db.First(doc, "name=?", name)
	return doc, tx.Error
}
func (udb *UserDb) GetPatientByID(ctx context.Context, id float64) (*user.Patient, error) {
	// db := ctx.Value("pgClient").(*gorm.DB)
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Patient{}
	tx := db.First(doc, "id=?", id)
	return doc, tx.Error
}
func (udb *UserDb) UpdatePatientByID(ctx context.Context, pt *user.Patient, id float64) (*user.Patient, error) {
	// db := ctx.Value("pgClient").(*gorm.DB)
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Model(pt).Where("id=?", id).Updates(pt)
	return pt, tx.Error
}
func (udb *UserDb) UpdateDoctorByID(ctx context.Context, doc *user.Doctor, id float64) (*user.Doctor, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Model(doc).Where("id=?", id).Updates(doc)
	return doc, tx.Error
}

func (udb *UserDb) CreateConnection(ctx context.Context, conn *user.Connections) (*user.Connections, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(conn)
	return conn, tx.Error
}
