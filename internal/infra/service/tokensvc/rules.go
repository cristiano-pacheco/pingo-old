package tokensvc

import (
	_ "embed"
)

// Package name of our rego code.
const (
	opaPackage string = "pingo.rego"
)

// Core OPA policies.
var (
	//go:embed rego/authentication.rego
	regoAuthenticationPolicy string
)
