package core

type Writer interface {
	// Create is used to create record in DB
	//
	// Based on object provided it will create record in DB
	Create(object interface{}) error
	// Update is used to update records in DB
	//
	// Based on filters provided it will update corresponding records
	Update(object interface{}, filters ...interface{}) error
	// Ping function is used to ping the database
	//
	// It is used to check the health status of DB
	Ping() error
	// Delete is used to delete records in DB
	//
	// Based on filters provided it will delete corresponding records
	Delete(object interface{}) error
}

// Command to generate the mock implementation of the Writer interface
//go:generate mockgen -source=write.go -destination=mock-impl/write-impl.go -package=mock_impl
