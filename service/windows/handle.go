package windows

import (
	"golang.org/x/sys/windows/svc"
)

type Handle struct {
	Start func()
	Stop func()
}

func (h *Handle) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (svcSpecificEC bool, exitCode uint32)  {
	const cmdAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue

	s <- svc.Status{State: svc.StartPending}
	s <- svc.Status{State: svc.Running, Accepts: cmdAccepted}

	go h.Start()

loop:
	for{
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				s <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				break loop
			case svc.Pause:
				s <- svc.Status{State: svc.Paused, Accepts: cmdAccepted}
			case svc.Continue:
				s <- svc.Status{State: svc.Running, Accepts: cmdAccepted}
			}
		}
	}

	s <- svc.Status{State: svc.StopPending}
	h.Stop()
	s <- svc.Status{State: svc.Stopped}

	return
}