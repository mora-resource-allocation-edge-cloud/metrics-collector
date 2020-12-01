package conf

import (
	"os"
	"strconv"
)

const (
	AppName     = "metrics-collector-manager"
	DefaultPort = 7070
)

var (
	MongoHost             = getEnv("MONGO_HOST", "mongo")
	MongoUser             = getEnv("MONGO_INITDB_ROOT_USERNAME", "root")
	MongoPassword         = getEnv("MONGO_INITDB_ROOT_PASSWORD", "root")
	MongoDBName           = getEnv("MONGO_DBNAME", "dashboards")
	MongoPort             = getIntFromEnv("POSTGRES_PORT", 27017)
	MongoAuthSource       = getEnv("MONGO_AUTH_SOURCE", "admin")
	DBMSConnectionTimeout = getIntFromEnv("DBMS_CONNECTION_TIMEOUT", 10)
	SSLMode               = getEnv("SSL_ENABLE", "disabled")
)

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Simple helper function to read an environment as int64 or return a default value
func getIntFromEnv(key string, defaultVal int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		if v, err := strconv.ParseInt(value, 10, 32); err == nil {
			return v
		}
	}
	return defaultVal
}
