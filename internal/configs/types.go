package configs

type (
	Database struct {
		DataSourceName string
	}

	Service struct {
		Port      string
		SecretKey string
	}

	Config struct {
		Service  Service
		Database Database
	}
)
