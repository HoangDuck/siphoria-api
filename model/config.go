package model

type Config struct {
	Server struct {
		Host              string `yaml:"HOST"`
		Port              string `yaml:"PORT"`
		JwtExpired        int    `yaml:"JWT_EXPIRED"`
		JwtRefreshExpired int    `yaml:"JWT_REFRESH_EXPIRED"`
		Domain            string `yaml:"DOMAIN"`
		SecretKey         string `yaml:"SECRET_KEY"`
		SecretRefreshKey  string `yaml:"SECRET_REFRESH_KEY"`
	}
	Database struct {
		DbName     string `yaml:"DB_NAME"`
		DbHost     string `yaml:"DB_HOST"`
		DbPort     string `yaml:"DB_PORT"`
		DbUserName string `yaml:"DB_USER_NAME"`
		DbPassword string `yaml:"DB_PASSWORD"`
	}
	Oauth2 struct {
		ClientId    string `yaml:"GOOGLE_CLIENT_ID"`
		ClientSec   string `yaml:"GOOGLE_CLIENT_SEC"`
		RedirectUrl string `yaml:"REDIRECT_URL"`
	}
	Email struct {
		SenderEmail string `yaml:"SENDER_EMAIL"`
		SenderPass  string `yaml:"SENDER_PASS"`
	}
	FirebaseService struct {
		ProjectID   string `yaml:"PROJECT_ID"`
		JsonKeyPath string `yaml:"JSON_KEY_PATH"`
	}
	Cloudinary struct {
		CloudinaryUrl string `yaml:"CLOUDINARY_URL"`
	}
	FacebookOauth2 struct {
		ClientId    string `yaml:"FACEBOOK_CLIENT_ID"`
		ClientSec   string `yaml:"FACEBOOK_CLIENT_SEC"`
		RedirectUrl string `yaml:"REDIRECT_URL"`
	}
	VNPay struct {
		VNPTmnCode    string `yaml:"vnp_TmnCode"`
		VNPHashSecret string `yaml:"vnp_HashSecret"`
		VNPUrl        string `yaml:"vnp_Url"`
	}
	Stripe struct {
		PublishableKey string `yaml:"publishable_key"`
		SecretKey      string `yaml:"secret_key"`
	}
}

type ConfigurationUrlDefine struct {
	ID          int    `gorm:"primary_key;autoIncrement"`
	Key         string `gorm:"key"`
	Value       string `gorm:"value"`
	Description string `gorm:"description"`
}
