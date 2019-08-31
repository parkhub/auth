package client

import (
	"github.com/ory/fosite"
	"github.com/justpark/auth/hydra/driver/configuration"
	"github.com/justpark/auth/hydra/x"
)

type InternalRegistry interface {
	x.RegistryWriter
	Registry
}

type Registry interface {
	ClientValidator() *Validator
	ClientManager() Manager
	ClientHasher() fosite.Hasher
}

type Configuration interface {
	configuration.Provider
}
