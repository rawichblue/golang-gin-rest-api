package config

// Config is a struct that contains all the configuration of the application.
type Config struct {
	AppName string
	AppKey  string
	AppEnv  string
	Debug   bool

	Port           int
	HttpPrefork    bool
	HttpJsonNaming string

	SslCaPath      string
	SslPrivatePath string
	SslCertPath    string

	OtelEnable             bool
	OtelDebugMetricDisable bool
	OtelCollectorEndpoint  string
	OtelTraceRatio         float64

	JwtSecret   string
	JwtDuration int64

	RedirectUrl  string
	ClientId     string
	ClientSecret string
	GrantType    string

	S3BucketUrl    string
	S3BucketName   string
	S3BucketKey    string
	S3BucketSecret string

	FirebaseUrl string
	FirebaseKey string

	GcsKeyPath    string
	GcsBucketName string
	IsStaging     bool
}

var App = Config{

	AppName: "go_app",
	Port:    8080,
	AppKey:  "secret",
	AppEnv:  "development",
	Debug:   false,

	HttpPrefork:    false,
	HttpJsonNaming: "snake_case",

	SslCaPath:      "storage/cert/ca.pem",
	SslPrivatePath: "storage/cert/server.pem",
	SslCertPath:    "storage/cert/server-key.pem",

	OtelCollectorEndpoint:  "localhost:4317",
	OtelEnable:             false,
	OtelDebugMetricDisable: true,
	OtelTraceRatio:         1,

	JwtSecret:   "secret",
	JwtDuration: 720,

	RedirectUrl:  "",
	ClientId:     "",
	ClientSecret: "",
	GrantType:    "",

	S3BucketUrl:    "",
	S3BucketName:   "",
	S3BucketKey:    "",
	S3BucketSecret: "",
	FirebaseUrl:    "",
	FirebaseKey:    "",

	GcsKeyPath:    "",
	GcsBucketName: "",
}
