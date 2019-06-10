package keys

type env struct {
	DatabaseHost         string
	DatabaseUsername     string
	DatabasePassword     string
	DatabasePort         string
	DatabaseName         string
	DatabaseSSL          string
	TestDatabaseHost     string
	TestDatabaseUsername string
	TestDatabasePassword string
	TestDatabasePort     string
	TestDatabaseName     string
	TestDatabaseSSL      string
	AppPort              string
	AppEnv               string
	ListenerReadTimeout  string
	ListenerWriteTimeout string
	JWTSecretKey         string
}

// Env holds Key for env variables
var Env = env{
	DatabaseHost:         "DATABASE_HOST",
	DatabaseUsername:     "DATABASE_USERNAME",
	DatabasePassword:     "DATABASE_PASSWORD",
	DatabasePort:         "DATABASE_PORT",
	DatabaseName:         "DATABASE_NAME",
	DatabaseSSL:          "DATABASE_SSL",
	TestDatabaseHost:     "TEST_DATABASE_HOST",
	TestDatabaseUsername: "TEST_DATABASE_USERNAME",
	TestDatabasePassword: "TEST_DATABASE_PASSWORD",
	TestDatabasePort:     "TEST_DATABASE_PORT",
	TestDatabaseName:     "TEST_DATABASE_NAME",
	TestDatabaseSSL:      "TEST_DATABASE_SSL",
	AppPort:              "APP_PORT",
	AppEnv:               "APP_ENV",
	ListenerReadTimeout:  "LISTENER_READ_TIMEOUT",
	ListenerWriteTimeout: "LISTENER_WRITE_TIMEOUT",
	JWTSecretKey:         "JWT_SECRET_KEY",
}
