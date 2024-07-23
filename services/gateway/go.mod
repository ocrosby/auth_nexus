module github.com/ocrosby/auth-nexus/services/gateway

go 1.22.5

require (
	github.com/ocrosby/auth-nexus/internal v0.0.0
	github.com/ocrosby/auth-nexus/pkg v0.0.0
)

require golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect

replace github.com/ocrosby/auth-nexus/pkg => ../../pkg

replace github.com/ocrosby/auth-nexus/internal => ../../internal
