package utils

import (
	"os"
	"syscall"
)

// ShutDownSignals returns all the singals that are being watched for to shut down services.
func ShutDownSignals() []os.Signal {
	return []os.Signal{
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL,
	}
}
