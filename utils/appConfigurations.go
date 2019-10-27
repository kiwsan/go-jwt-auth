package utils

//https://github.com/jinzhu/configor
var Config = struct {
	AppName string `default:"app name"`

	Database struct {
		Host     string `default:"localhost"`
		User     string `default:"root"`
		Password string `required:"true" env:"password"`
		Port     uint   `default:"27017"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}
