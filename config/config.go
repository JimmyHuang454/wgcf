package config

const (
	DeviceId    = "device_id"
	AccessToken = "access_token"
	PrivateKey  = "private_key"
	PublicKey   = "public_key"
	LicenseKey  = "license_key"
)

type Context struct {
	DeviceId    string
	AccessToken string
	PrivateKey  string
	PublicKey   string
	LicenseKey  string
}
