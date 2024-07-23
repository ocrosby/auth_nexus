# AuthNexus

A monorepo to house a generic solution for OAuth 2 authentication and authorization.

## Usage

When updating a shared dependency, you can do so in it's own project directory and then run `go mod tidy` in each 
dependent project to ensure they are using the updated version.

