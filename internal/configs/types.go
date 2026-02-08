package configs

type (
	Database struct {
		DataSourceName string
	}

	Service struct {
		Ports     string
		SecretKey string
	}

	Config struct {
		Service  Service
		Database Database
	}
)
