package core

type Reader interface {
	First(out interface{}, filter interface{}) error
	Last(out interface{}, filter interface{}) error
	Where(out interface{}, query string, params ...interface{}) (interface{}, error)
	Find(out interface{}, filter interface{}, exclude interface{}) error
	FetchPaginated(out interface{}, filter interface{}, exclude interface{}, limit int, offset int, order string) (err error)
	Count(out interface{}, count interface{}, filter interface{}, exclude interface{}) (err error)
}

// Command to generate the mock implementation of the Reader interface
//go:generate mockgen -source=read.go -destination=mock-impl/read-impl.go -package=mock_impl
