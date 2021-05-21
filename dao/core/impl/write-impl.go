package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/config"
	"gitlab.com/umi7/DezHab_user_backend/dao/core"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

var Write core.Writer

type WriterImpl struct {
	master *gorm.DB
}

const (
	postgresDialect = "postgres"
	ErrOnDB         = "some error occurred on database call"
	ErrRowsAffected = "unexpected number of rows affected"
	plsqlConfFormat = "host=%s port=%s dbname=%s sslmode=disable user=%s password=%s"
	CREATE          = "create"
	UPDATE          = "update"
	DELETE          = "delete"
	ErrOnFilters    = "no filters specified"
)

func getDb(host, port, database, user, password string, idleConnections, openConnections int) *gorm.DB {
	plsqlConf := fmt.Sprintf(plsqlConfFormat, host, port, database, user, password)
	db, err := gorm.Open(postgresDialect, plsqlConf)
	if err != nil {
		logger.Error(context.Background(), err.Error())
		panic(err)
	}
	db.DB().SetMaxIdleConns(idleConnections)
	db.DB().SetMaxOpenConns(openConnections)
	db.LogMode(true)
	db.SetLogger(logger.GetDBLogger())
	return db
}

func init() {
	writerImpl := WriterImpl{
		master: getDb(config.AppConfig.Database.Master.Host,
			config.AppConfig.Database.Master.Port,
			config.AppConfig.Database.Master.Name,
			config.AppConfig.Database.Master.Username,
			config.AppConfig.Database.Master.Password,
			config.AppConfig.Database.Master.IdleConnections,
			config.AppConfig.Database.Master.OpenConnections),
	}
	Write = writerImpl
	if config.AppConfig.Database.Master.NeedsMigration {
		//TODO : find a way to fetch the model from config
		writerImpl.autoMigrate(&models.ApiClient{}, models.BookKeeping{})
	}
}

// Create is used to create record in DB
//
// Based on object provided it will create record in DB
func (w WriterImpl) Create(object interface{}) error {
	return w.writeUtil(object, CREATE)
}

// Update is used to update records in DB
//
// Based on filters provided it will update corresponding records
func (w WriterImpl) Update(object interface{}, filters ...interface{}) error {
	return w.writeUtil(object, UPDATE, filters...)
}

// Ping function is used to ping the database
//
// It is used to check the health status of DB
func (w WriterImpl) Ping() error {
	return w.master.DB().Ping()
}

// Delete is used to delete records in DB
//
// Based on filters provided it will delete corresponding records
func (w WriterImpl) Delete(object interface{}) error {
	return w.writeUtil(object, DELETE)
}

// autoMigrate takes the pointer to the model and creates the table with
// the name in pluralised form
func (w WriterImpl) autoMigrate(values ...interface{}) {
	for _, val := range values {
		if w.master.HasTable(val) {
			continue
		}
		if db := w.master.AutoMigrate(val); db != nil {
			if err := db.Error; err != nil {
				panic(err)
			}
		}
	}
}

func (w WriterImpl) writeUtil(object interface{}, operation string, filters ...interface{}) (err error) {
	// gorm handles the transaction for single record write operation
	// so, we are owning the custom txn, if the record is an array.
	if val, ok := object.([]interface{}); ok {
		if len(val) == 1 {
			object = val[0]
			goto singleRecord
		}
		return w.withTxn(val, operation)
	}
	//singleRecord is used to insert a singe record
singleRecord:
	var db *gorm.DB
	if operation == CREATE {
		db = w.master.Create(object)
	} else if operation == UPDATE {
		//TODO : Find a generic way to handle
		if len(filters) < 3 {
			err = errors.New(ErrOnFilters)
			return
		}
		updateQuery := filters[0]
		updateValues := filters[1]
		values := filters[2]
		db = w.master.Model(object).Where(updateQuery, updateValues).Updates(values)
	} else if operation == DELETE {
		db = w.master.Delete(object)
	}
	if db != nil {
		if db.Error != nil {
			err = db.Error
			return
		}
		if db.RowsAffected != 1 {
			err = errors.New(ErrRowsAffected)
		}
		return
	}
	return errors.New(ErrOnDB)
}

func (w WriterImpl) withTxn(object []interface{}, operation string) (err error) {
	txn := w.master.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			txn.Rollback()
		}
	}()
	var db *gorm.DB
	if operation == CREATE {
		db = txn.Create(object)
	} else if operation == UPDATE {
		db = txn.Updates(object)
	}
	if db != nil {
		if db.Error != nil {
			err = db.Error
			return
		}
		if db.RowsAffected != int64(len(object)) {
			err = errors.New(ErrRowsAffected)
			return
		}
		err = txn.Commit().Error
		return
	}
	return errors.New(ErrOnDB)
}
