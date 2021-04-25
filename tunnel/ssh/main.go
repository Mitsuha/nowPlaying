package ssh

import (
	"golang.org/x/crypto/ssh"
)

type SSH struct {
	Config ssh.ClientConfig
	ConnAddr string
	LocalAddr string
	RemoteAddr string

	Remote *Remote
}

func (s *SSH) MappingRemote() error {
	conn,err := ssh.Dial("tcp", s.ConnAddr, &s.Config)

	if err != nil {
		return err
	}

	if s.Remote == nil {
		s.Remote = & Remote{
			Client:      conn,
			Addr:        s.RemoteAddr,
			MappingAddr: s.LocalAddr,
		}
	}

	return s.Remote.Mapping()
}

func (s *SSH) StopMappingRemote() error {
	return s.Remote.Stop()
}