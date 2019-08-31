package jwk

import (
	"github.com/justpark/auth/hydra/driver/configuration"
	"github.com/justpark/auth/hydra/x"
)

type InternalRegistry interface {
	x.RegistryWriter
	x.RegistryLogger
	Registry
}

type Registry interface {
	KeyManager() Manager
	KeyGenerators() map[string]KeyGenerator
	KeyCipher() *AEAD
}

type Configuration interface {
	configuration.Provider
}
