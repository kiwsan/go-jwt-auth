package utils

//https://github.com/jinzhu/configor
var Config = struct {
	AppName string `default:"app name"`

	Database struct {
		Host     string `default:"localhost"`
		User     string `default:"root"`
		Password string `required:"true" env:"password"`
		Port     string `default:"27017"`
	}

	RabbitMq struct {
		Host     string `default:"localhost"`
		User     string `default:"user"`
		Password string `required:"true" env:"secret"`
		Port     string `default:"5672"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}
