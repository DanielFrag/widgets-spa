package utils

import (
	"net/http"
)

type HandlerFuncInjector struct {
	dependecies []func(http.HandlerFunc) http.HandlerFunc
	handler     http.HandlerFunc
}

func (hfi *HandlerFuncInjector) InjectDependencies() {
	for _, d := range hfi.dependecies {
		hfi.handler = d(hfi.handler)
	}
}
