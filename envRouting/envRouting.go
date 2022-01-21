// Package envRouting that provides environment variable access with static and secure bindings/configuration
package envRouting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// Port ...
	Port string

	// SecretKey ...
	SecretKey string

	// StaticWebLocation ...
	StaticWebLocation string

	// DatabaseName ...
	DatabaseName string

	// SQLiteFilename ...
	SQLiteFilename string

	// MySQLUsername ...
	MySQLUsername string

	// MySQLPassword ...
	MySQLPassword string

	// MySQLHost ...
	MySQLHost string

	// PostgresUsername ...
	PostgresUsername string

	// PostgresPassword ...
	PostgresPassword string

	// PostgresHost ...
	PostgresHost string

	// PostgresPort ...
	PostgresPort string

	// PostgresSSLMode ...
	PostgresSSLMode string

	// PostgresTimezone ...
	PostgresTimezone string

	// TwilioAccountSID ..
	TwilioAccountSID string

	// TwilioAuthenticationToken ..
	TwilioAuthenticationToken string

	// TwilioPhoneNumber ..
	TwilioPhoneNumber string

	// PostgresURL
	PostgresURL string

	// SSL
	SSL string

	// SSLCertificate ...
	SSLCertificate string

	// SSLKey ...
	SSLKey string

	// K2CServicesHost
	K2CServicesHost string

	// K2CServicesPort
	K2CServicesPort string

	// K2CServicesAPIVersionEndpoint
	K2CServicesAPIVersionEndpoint string

	// K2CServicesTransactionEndpoint
	K2CServicesTransactionEndpoint string

	// K2CServicesTransactionEndpoint
	K2CServicesTransactionSummaryEndpoint string

	// K2CServicesByBranchUserIDEndpoint
	K2CServicesByBranchUserIDEndpoint string

	// K2CServicesTransactionEndpoint
	K2CServicesInstitutionEndpoint string

	// K2CServicesTransactionEndpoint
	K2CServicesBranchEndpoint string

	// K2CServicesTransactionEndpoint
	K2CServicesBranchUserEndpoint string
)

// LoadEnv Staticly load environment variables
func LoadEnv() {
	Port = getEnv("PORT")
	SecretKey = getEnv("SECRET_KEY")
	StaticWebLocation = getEnv("STATIC_WEB_LOCATION")
	DatabaseName = getEnv("DATABASE_NAME")

	SQLiteFilename = getEnv("SQLITE_FILENAME")

	MySQLUsername = getEnv("MYSQL_USERNAME")
	MySQLPassword = getEnv("MYSQL_PASSWORD")
	MySQLHost = getEnv("MYSQL_HOST")

	PostgresUsername = getEnv("POSTGRES_USERNAME")
	PostgresPassword = getEnv("POSTGRES_PASSWORD")
	PostgresHost = getEnv("POSTGRES_HOST")
	PostgresPort = getEnv("POSTGRES_PORT")
	PostgresSSLMode = getEnv("POSTGRES_SSL_MODE")
	PostgresTimezone = getEnv("POSTGRES_TIMEZONE")

	TwilioAccountSID = getEnv("TWILIO_ACCOUNT_SID")
	TwilioAuthenticationToken = getEnv("TWILIO_AUTH_TOKEN")
	TwilioPhoneNumber = getEnv("TWILIO_PHONE_NUMBER")

	PostgresURL = getEnv("POSTGRES_URL")

	SSL = getEnv("SSL")
	SSLCertificate = getEnv("SSL_CERTIFICATE")
	SSLKey = getEnv("SSL_KEY")

	K2CServicesHost = getEnv("K2C_SERVICES_HOST")
	K2CServicesPort = getEnv("K2C_SERVICES_PORT")
	K2CServicesAPIVersionEndpoint = getEnv("K2C_SERVICES_API_VERSION_ENDPOINT")
	K2CServicesTransactionEndpoint = getEnv("K2C_SERVICES_TRANSACTION_ENDPOINT")
	K2CServicesTransactionSummaryEndpoint = getEnv("K2C_SERVICES_TRANSACTION_SUMMARY_ENDPOINT")
	K2CServicesByBranchUserIDEndpoint = getEnv("K2C_SERVICES_BY_BRANCH_USER_ID_ENDPOINT")
	K2CServicesInstitutionEndpoint = getEnv("K2C_SERVICES_INSTITUTION_ENDPOINT")
	K2CServicesBranchEndpoint = getEnv("K2C_SERVICES_BRANCH_ENDPOINT")
	K2CServicesBranchUserEndpoint = getEnv("K2C_SERVICES_BRANCH_USER_ENDPOINT")
}

func getEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
