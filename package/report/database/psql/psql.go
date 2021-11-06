package psql

import (
	"askUs/v1/package/report"
	"askUs/v1/util"
	"context"

	"gorm.io/gorm"
)

type ReportDB struct {
}

func Init() report.DB {
	return &ReportDB{}
}

func (rdb ReportDB) Create(ctx context.Context, ur *report.UserReport) (*report.UserReport, error) {
	db := util.GetFromServiceCtx(ctx, "pgClient").(*gorm.DB)
	tx := db.Create(ur)
	return ur, tx.Error
}
func (rdb ReportDB) Update(ctx context.Context,
	ur *report.UserReport) (
	*report.UserReport, error) {
	db := util.GetFromServiceCtx(ctx, "pgClient").(*gorm.DB)
	tx := db.Save(ur)
	return ur, tx.Error
}
func (rdb ReportDB) Delete(ctx context.Context,
	ur *report.UserReport) (
	*report.UserReport, error) {
	db := util.GetFromServiceCtx(ctx, "pgClient").(*gorm.DB)
	tx := db.Delete(ur)
	return ur, tx.Error
}
func (rdb ReportDB) GetAll(context.Context,
	int) (*[]report.UserReport, error) {
	return nil, nil
}

func (rdb ReportDB) Get(ctx context.Context, reportID string, patientID float64) (*report.UserReport, error) {
	db := util.GetFromServiceCtx(ctx, "pgClient").(*gorm.DB)
	ur := &report.UserReport{}
	tx := db.Where("id=? AND patient_id=?", reportID, patientID).Find(ur)
	return ur, tx.Error
}
