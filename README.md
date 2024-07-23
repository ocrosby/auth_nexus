# AuthNexus

A monorepo to house a generic solution for OAuth 2 authentication and authorization.

## Overview

```mermaid
sequenceDiagram
    participant U as User
    participant C as Client
    participant A as Authorization Server
    participant R as Resource Server

    U->>C: Request access to resources
    C->>U: Redirect to Authorization Server login
    U->>A: Login and approve access
    A->>U: Authorization code
    U->>C: Redirect with authorization code
    C->>A: Exchange authorization code for access token
    A->>C: Access token
    C->>R: Request resources with access token
    R->>C: Resources
```

## Usage

When updating a shared dependency, you can do so in it's own project directory and then run `go mod tidy` in each 
dependent project to ensure they are using the updated version.

