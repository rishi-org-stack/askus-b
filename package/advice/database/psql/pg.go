package psql

import (
	"askUs/v1/package/advice"
	"context"

	"gorm.io/gorm"
)

type AdviceData struct {
}

func (adb AdviceData) CreateAdvice(ctx context.Context, adv *advice.Advice) (*advice.Advice, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(adv)
	return adv, tx.Error
}
func (adb AdviceData) CreateLike(ctx context.Context, adv *advice.Like) (*advice.Like, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	tx := db.Create(adv)
	return adv, tx.Error
}
func (adb AdviceData) GetGlobalAdviceByID(ctx context.Context, id float64) (*advice.Advice, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	adv := &advice.Advice{}
	tx := db.First(adv, "id=? AND type=?", id, advice.GLOBAL)
	return adv, tx.Error
}

func (adb AdviceData) GetPersonelAdviceByID(ctx context.Context, id float64) (*advice.Advice, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	adv := &advice.Advice{}
	tx := db.First(adv, "id=? AND type=?", id, advice.PERSONEL)
	return adv, tx.Error
}
func (adb AdviceData) GetAdviceByPatientID(ctx context.Context, id float64) (*advice.Advice, error) {
	// postedFor
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	adv := &advice.Advice{}
	tx := db.First(adv, "posted_for=?", id)
	return adv, tx.Error
}
func (adb AdviceData) GetAdviceByDocID(ctx context.Context, id float64) (*advice.Advice, error) {
	// postedFor
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	adv := &advice.Advice{}
	tx := db.First(adv, "posted_by=?", id)
	return adv, tx.Error
}
func (adb AdviceData) GetAllPersonelAdvices(ctx context.Context, id float64) (*[]advice.Advice, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	adv := &[]advice.Advice{}
	tx := db.Find(adv, "posted_for=?", id)
	return adv, tx.Error
}

func (adb AdviceData) GetAllPersonelAdvicesPostedByDoc(ctx context.Context, id float64) (*[]advice.Advice, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	adv := &[]advice.Advice{}
	tx := db.Preload("Patient").Find(adv, "posted_by=? and type=?", id, advice.PERSONEL)
	// ok := db.Exec("select  *  from advice.advices  join usr.patients on advice.advices.patient_id=usr.patients.id  where type='PERSONEL';")
	// fmt.Println("ok")
	// fmt.Println("err ck :->", ok.Error)
	// ck := ok.Scan(adv)
	// fmt.Println(adv)
	// fmt.Println("err ck :->", ck.Error)
	return adv, tx.Error
}
func (adb AdviceData) GetAllDocAdvices(ctx context.Context, id float64) (*[]advice.Advice, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	adv := &[]advice.Advice{}
	tx := db.Find(adv, "posted_by=?", id)
	return adv, tx.Error
}

func (adb AdviceData) GetAdviceWithPIDAndDID(ctx context.Context, did float64, pid string) (*[]advice.Advice, error) {
	db := ctx.Value("surround").(map[string]interface{})["pgClient"].(*gorm.DB)
	adv := &[]advice.Advice{}
	tx := db.Find(adv, "posted_by=? and posted_for=?", did, pid)
	return adv, tx.Error
}
