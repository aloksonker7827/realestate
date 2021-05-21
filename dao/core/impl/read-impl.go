package impl

import (
	"errors"
	"github.com/jinzhu/gorm"
	"gitlab.com/umi7/DezHab_user_backend/config"
	"gitlab.com/umi7/DezHab_user_backend/dao/core"
	"reflect"
)

var Read core.Reader

type ReaderImpl struct {
	slave *gorm.DB
}

func init() {
	Read = ReaderImpl{
		slave: getDb(config.AppConfig.Database.Master.Host,
			config.AppConfig.Database.Master.Port,
			config.AppConfig.Database.Master.Name,
			config.AppConfig.Database.Master.Username,
			config.AppConfig.Database.Master.Password,
			config.AppConfig.Database.Slave.IdleConnections,
			config.AppConfig.Database.Slave.OpenConnections),
	}
}

// First gets the first record given the struct with filters.
// In a given struct,
// First is ordered by primary key by default.
// Use this function if you want to take the first record
func (r ReaderImpl) First(out interface{}, filter interface{}) (err error) {
	if db := r.slave.First(out, filter); db != nil {
		if db.Error != nil {
			err = db.Error
		}
		return
	}
	err = errors.New(ErrOnDB)
	return
}

func (r ReaderImpl) Last(out interface{}, filter interface{}) (err error) {
	if db := r.slave.Last(out, filter); db != nil {
		if db.Error != nil {
			err = db.Error
		}
		return
	}
	err = errors.New(ErrOnDB)
	return
}

func (r ReaderImpl) Where(outType interface{}, query string, params ...interface{}) (out interface{}, err error) {
	elemType := reflect.ValueOf(outType).Type()
	sliceValue := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0)
	results := reflect.New(sliceValue.Type())
	iresults := results.Interface()

	if db := r.slave.Where(query, params...).Find(iresults); db != nil {
		out = results.Elem().Interface()
		if db.Error != nil {
			err = db.Error
		}
		return
	}
	err = errors.New(ErrOnDB)
	return
}

// Use this function if you want to fetch all records
func (r ReaderImpl) Find(out interface{}, filter interface{}, exclude interface{}) (err error) {
	if db := r.slave.Not(exclude).Find(out, filter); db != nil {
		if db.Error != nil {
			err = db.Error
		}
		return
	}
	err = errors.New(ErrOnDB)
	return
}

// Use this function if you want to fetch all records
func (r ReaderImpl) FetchPaginated(out interface{}, filter interface{}, exclude interface{}, limit int, offset int, order string) (err error) {
	if db := r.slave.Limit(limit).Offset(offset).Not(exclude).Order(order).Find(out, filter); db != nil {
		if db.Error != nil {
			err = db.Error
		}
		return
	}
	err = errors.New(ErrOnDB)
	return
}

// Use this function to count record by applying filters
func (r ReaderImpl) Count(out interface{}, count interface{}, filter interface{}, exclude interface{}) (err error) {
	if db := r.slave.Not(exclude).Where(filter).Find(out).Count(count); db != nil {
		if db.Error != nil {
			err = db.Error
		}
		return
	}
	err = errors.New(ErrOnDB)
	return
}
