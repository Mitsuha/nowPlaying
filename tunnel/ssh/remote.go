package ssh

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
)
// 将远程端口映射到本地
type Remote struct {
	Client *ssh.Client
	Addr string
	MappingAddr string
	listener net.Listener
}

func (r *Remote) Mapping() error {
	listener, err := r.Client.Listen("tcp", r.Addr)

	if err != nil {
		return err
	}
	r.listener = listener
	
	go func() {
		for {
			conn, err := r.listener.Accept()

			if err != nil {
				log.Println(err)
				continue
			}

			go func(remote net.Conn) {
				local, err := net.Dial("tcp", r.MappingAddr)

				if err != nil {
					log.Println(err)
					remote.Close()
				}

				r.forward(remote, local)
				remote.Close()
				local.Close()
			}(conn)
		}
	}()
	return nil
}

func (r *Remote) Stop() error{
	return r.listener.Close()
}

func (r Remote) forward(remote, local net.Conn) {
	done := make(chan bool)
	go func() {
		_, err := io.Copy(local, remote)

		if err != nil {
			done <- true
		}

	}()
	go func() {
		_, err := io.Copy(remote, local)
		if err != nil {
			done <- true
		}
	}()

	<- done
}