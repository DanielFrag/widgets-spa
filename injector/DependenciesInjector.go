package injector

import (
	"net/http"
)

type DependenciesInjector struct {
	dependecies []func(http.HandlerFunc) http.HandlerFunc
	handler     http.HandlerFunc
}

func (di *DependenciesInjector) InjectDependencies(w http.ResponseWriter, r *http.Request) {
	for _, d := range di.dependecies {
		di.handler = d(di.handler)
	}
}
