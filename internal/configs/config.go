package configs

import "github.com/spf13/viper"

var config *Config

type option struct {
	configFolders []string
	configFile    string
	configType    string
}

type Option func(*option)

func getDefaultConfigFolders() []string {
	return []string{"./configs"}
}

func getConfigFile() string {
	return "config"
}

func getConfigType() string {
	return "yaml"
}

func WithConfigFolders(configFolders []string) Option {
	return func(opt *option) {
		opt.configFolders = configFolders
	}
}

func WithConfigFile(configFile string) Option {
	return func(opt *option) {
		opt.configFile = configFile
	}
}

func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

func Init(opts ...Option) error {
	opt := &option{
		configFolders: getDefaultConfigFolders(),
		configFile:    getConfigFile(),
		configType:    getConfigType()}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	for _, configFolder := range opt.configFolders {
		viper.AddConfigPath(configFolder)
	}

	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	config = new(Config)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(&config)
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
