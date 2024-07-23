module github.com/ocrosby/auth_nexus/cmd

go 1.22.5

require (
	github.com/ocrosby/auth-nexus/internal v0.0.0
)

replace github.com/ocrosby/auth-nexus/internal => ../internal
