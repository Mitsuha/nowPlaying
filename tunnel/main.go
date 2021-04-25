package tunnel

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	subSsh "nowPlaying/tunnel/ssh"
	"strings"
)

func Ssh(server, password, local, remote string) (*subSsh.SSH, error) {
	host, err := parseServerHost(server)

	if err != nil {
		return nil, err
	}

	cfg, err := makeSshConfig(host[0],password)

	if err != nil {
		return nil, err
	}

	return &subSsh.SSH{
		Config:     *cfg,
		ConnAddr:  host[1],
		LocalAddr:  local,
		RemoteAddr: remote,
	}, nil
}

func parseServerHost(server string) ([]string, error) {
	if ! strings.Contains(server, "@") {
		return nil, errors.New("can not parser server name")
	}
	if ! strings.Contains(server, ":") {
		server += ":22"
	}

	return strings.SplitN(server, "@", 2), nil
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

