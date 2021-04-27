package tunnel

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"nowPlaying/config"
	subSsh "nowPlaying/tunnel/ssh"
	"strings"
)

//func Ssh(user, host, password, local, remote string) (*subSsh.SSH, error) {
func Ssh(tunCfg *config.TunnelCfg) (*subSsh.SSH, error) {
	cfg, err := makeSshConfig(tunCfg.User,tunCfg.Password)

	if err != nil {
		return nil, err
	}

	var host = tunCfg.Host
	if ! strings.Contains(host, ":") {
		host = host + ":22"
	}
	return &subSsh.SSH{
		Config:     *cfg,
		ConnAddr:  host,
		LocalAddr:  config.App.Listen,
		RemoteAddr: "0.0.0.0:" + tunCfg.Port,
	}, nil
}

func parsePrivateKey(keyPath string) (ssh.Signer, error) {
	buff, err := ioutil.ReadFile(keyPath)

	if err != nil {
		return nil, err
	}

	return ssh.ParsePrivateKey(buff)
}

func makeSshConfig(user, password string) (*ssh.ClientConfig, error) {
	key, err := parsePrivateKey(privateKetPath())

	if err != nil {
		return nil, err
	}

	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}, nil
}

