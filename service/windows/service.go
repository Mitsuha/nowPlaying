package windows

import (
	"fmt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"time"
)

func Install(path, name, desc string,param ...string ) error {
	manager, err := mgr.Connect()

	if err != nil {
		return err
	}
	// see https://docs.microsoft.com/zh-cn/dotnet/api/system.serviceprocess.servicestartmode
	service, err := manager.CreateService(name, path, mgr.Config{
		ServiceType:      windows.SERVICE_WIN32_OWN_PROCESS,
		StartType:        windows.SERVICE_AUTO_START,
		//ServiceStartName: name,
		DisplayName:      name,
		Description:      desc,
		DelayedAutoStart: true,
	}, param...)
	if err != nil {
		return err
	}

	defer service.Close()
	return nil
}

func Uninstall(name string) error {
	manager, err := mgr.Connect()

	if err != nil {
		return err
	}

	defer manager.Disconnect()
	svs, err := manager.OpenService(name)
	if err != nil {
		return err
	}

	return svs.Delete()
}

func Start(name string) error {
	manager, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer manager.Disconnect()

	s, err := manager.OpenService(name)
	if err !=nil {
		return err
	}
	defer s.Close()

	return s.Start()
}

func Stop(name string) error {
	return controlService(name, svc.Stop, svc.Stopped, 2 * time.Second)
}

func State(name string) (*svc.Status, error) {
	m, err := mgr.Connect()

	if err != nil {
		return nil, err
	}
	
	s, err := m.OpenService(name)
	if err != nil {
		return nil, err
	}
	status, err := s.Query()
	return &status, err
}

func InServiceModel() bool {
	isWinSvs, err := svc.IsWindowsService()
	if err != nil {
		fmt.Println(err)
	}
	return isWinSvs
}

func controlService(name string, c svc.Cmd, to svc.State, timeout time.Duration) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(name)
	if err != nil {
		return err
	}
	defer s.Close()

	// 发送修改状态的请求
	status, err := s.Control(c)

	if err != nil {
		return err
	}

	limit := time.Now().Add(timeout)

	for status.State != to {
		if limit.Before(time.Now()) {
			return fmt.Errorf("controlService: timeout waiting for service to go to state=%d", to)
		}
		time.Sleep(300* time.Microsecond)
		status, err = s.Query()

		if err != nil {
			return err
		}
	}
	return nil
}

func RunAsService(name string, start, stop func()) error {
	return svc.Run(name, &Handle{Start: start, Stop: stop})
}