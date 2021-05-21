package dto

type ServerConfig struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	IdleConnections int    `json:"idleConnections,string"`
	OpenConnections int    `json:"openConnections,string"`
	NeedsMigration  bool   `json:"needsMigration,string"`
}
type AWSConfig struct {
	ID            string `json:"id"`
	Token         string `json:"token"`
	Region        string `json:"region"`
	Bucket        string `json:"bucket"`
	WorkerBucket  string `json:"workerBucket"`
	ImageBucket   string `json:"imageBucket"`
	MappingBucket string `json:"mappingBucket"`
}

type Database struct {
	Master DatabaseConfig `json:"master"`
	Slave  DatabaseConfig `json:"slave"`
}

//Config struct contains the application config read from the json file.
type Config struct {
	Server    ServerConfig `json:"server"`
	Database  Database     `json:"database"`
	Aws       AWSConfig    `json:"aws"`
	BaseUrl   string       `json:"baseUrl"`
	StatusUrl string       `json:"statusUrl"`
}
