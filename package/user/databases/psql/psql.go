package psql

import (
	"askUs/v1/package/user"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDb struct{}

func (udb *UserDb) FindOrCreateDoctor(ctx context.Context, email string) (*user.Doctor, error) {
	// db := ctx.Value("pgClient").(*gorm.D
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Doctor{Info: user.Info{Email: email}}
	tx := db.Where("email=?", email).
		FirstOrCreate(doc)
	return doc, tx.Error
}
func (udb *UserDb) FindOrCreatePatient(ctx context.Context, email string) (*user.Patient, error) {
	// db := ctx.Value("pgClient").(*gorm.DB)
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Patient{Info: user.Info{Email: email}}
	tx := db.Where("email=?", email).
		FirstOrCreate(doc)
	return doc, tx.Error
}
func (udb *UserDb) GetDoctorByID(ctx context.Context, id float64) (*user.Doctor, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	doc := &user.Doctor{}
	model := db.Model(doc)
	tx := model.Preload(clause.Associations).Preload("FollowedByPatients.User").First(doc, "id=?", id)
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

func (udb *UserDb) CreateReq(ctx context.Context, req *user.Request) (*user.Request, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(req)
	return req, tx.Error
}
func (udb *UserDb) GetReqsWithSenderID(ctx context.Context, senderId float64) (*[]user.Request, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	reqs := &[]user.Request{}
	tx := db.Find(reqs, "sender_id=?", senderId)
	return reqs, tx.Error
}

func (udb *UserDb) GetReqsWithDocID(ctx context.Context, senderId float64) (*[]user.Request, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	reqs := &[]user.Request{}
	tx := db.Find(reqs, "doctor_id=?", senderId)
	return reqs, tx.Error
}
func (udb *UserDb) GetReq(ctx context.Context, Id string) (*user.Request, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	reqs := &user.Request{}
	tx := db.First(reqs, "id=?", Id)
	return reqs, tx.Error
}
func (udb *UserDb) GetReqWithSenderandDocID(ctx context.Context, senderId float64, docId string, typeOfClient string) (*user.Request, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	reqs := &user.Request{}
	tx := db.First(reqs, "sender_id=? AND doctor_id=? AND genrated_by=?", senderId, docId, typeOfClient)
	return reqs, tx.Error
}
func (udb *UserDb) UpdateReq(ctx context.Context, req *user.Request) (*user.Request, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Where("id=?", req.ID).Updates(req)
	return req, tx.Error
}

func (udb *UserDb) CreateFBD(ctx context.Context, req *user.FollowedByDoctor) (*user.FollowedByDoctor, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(req)
	return req, tx.Error
}
func (udb *UserDb) CreateFD(ctx context.Context, req *user.FollowingDoctor) (*user.FollowingDoctor, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(req)
	return req, tx.Error
}
func (udb *UserDb) CreateFBP(ctx context.Context, req *user.FollowedByPatient) (*user.FollowedByPatient, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(req)
	return req, tx.Error
}
func (udb *UserDb) CreateFDBP(ctx context.Context, req *user.FollowedDoctorsByPatient) (*user.FollowedDoctorsByPatient, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(req)
	return req, tx.Error
}

func (udb *UserDb) GetFBD(ctx context.Context, id float64) (*[]user.FollowedByDoctor, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	fbds := &[]user.FollowedByDoctor{}
	tx := db.Find(fbds, "doctor_id=?", id)
	return fbds, tx.Error
}
func (udb *UserDb) GetFD(ctx context.Context, id float64) (*[]user.FollowingDoctor, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	fbds := &[]user.FollowingDoctor{}
	tx := db.Find(fbds, "doctor_id=?", id)
	return fbds, tx.Error
}
func (udb *UserDb) GetFBP(ctx context.Context, id float64) (*[]user.FollowedByPatient, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	fbps := &[]user.FollowedByPatient{}
	// doctor := &[]user.Doctor{}
	model := db.Model(fbps)
	err := model.Preload("Doctor").Preload("User").Find(fbps, "doctor_id=?", id) //, "id=?", id)
	// model.Association("Doctor").Find(doctor)

	return fbps, err.Error
}
func (udb *UserDb) GetFDBP(ctx context.Context, id float64) (*[]user.FollowedDoctorsByPatient, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	fbds := &[]user.FollowedDoctorsByPatient{}
	tx := db.Find(fbds, "patient_id=?", id)
	return fbds, tx.Error
}
