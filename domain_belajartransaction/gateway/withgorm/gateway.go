package withgorm

import (
	"context"
	"gogen-transaction/domain_belajartransaction/model/entity"
	"gogen-transaction/shared/gogen"
	"gogen-transaction/shared/infrastructure/config"
	"gogen-transaction/shared/infrastructure/database"
	"gogen-transaction/shared/infrastructure/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type gateway struct {
	*database.GormWithTransaction
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	//db      *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(entity.Product{}, entity.Order{})
	if err != nil {
		panic(err)
	}

	return &gateway{
		GormWithTransaction: database.NewGormWithTransaction(db, log),
		log:                 log,
		appData:             appData,
		config:              cfg,
		//db:                  db,
	}
}

func (r *gateway) SaveProduct(ctx context.Context, obj *entity.Product) error {
	r.log.Info(ctx, "called")

	err := r.ExtractDB(ctx).Save(obj).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) SaveOrder(ctx context.Context, obj *entity.Order) error {
	r.log.Info(ctx, "called")

	err := r.ExtractDB(ctx).Save(obj).Error
	if err != nil {
		return err
	}

	return nil
}
