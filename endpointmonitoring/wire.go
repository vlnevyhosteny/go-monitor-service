//go:build wireinject
// +build wireinject

package endpointmonitoring

import "github.com/google/wire"

func Wire(databaseFile string) *Service {
	panic(wire.Build(Set))
}
