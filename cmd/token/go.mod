module github.com/ocrosby/auth_nexus/token

go 1.22.5

require (
	github.com/ocrosby/auth_nexus/internal v0.0.0
	github.com/ocrosby/auth_nexus/pkg v0.0.0
)

require golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect

replace github.com/ocrosby/auth_nexus/pkg => ../../pkg

replace github.com/ocrosby/auth_nexus/internal => ../../internal
