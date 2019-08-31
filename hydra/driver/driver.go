package driver

import (
	"github.com/justpark/auth/hydra/driver/configuration"
)

type Driver interface {
	Configuration() configuration.Provider
	Registry() Registry
	CallRegistry() Driver
}
