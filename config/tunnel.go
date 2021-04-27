package config

type TunnelCfg struct {
	Enable bool
	User, Host, Password, Port, Forward, Type string
}
