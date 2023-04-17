package config

type Configuration struct {
	Listen         string `envconfig:"listen" default:":8080"`
	ListenAdmin    string `envconfig:"listen" default:":8082"`
	ListenInternal string `envconfig:"listen_internal" default:":8081"`

	LogLevel string `envconfig:"log_level" default:"debug"`

	MongoConnString string `envconfig:"MONGO_CONN_STRING" default:"mongodb://192.168.1.60:27017/"`
	MongoDBName     string `envconfig:"MONGO_DB_NAME" default:"hismap"`
}
