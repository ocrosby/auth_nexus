# AuthNexus

A monorepo to house a generic solution for OAuth 2 authentication and authorization.

## Overview

### OAuth 2.0

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

### Registration

```mermaid
sequenceDiagram
participant U as User
participant C as Client
participant A as Authorization Server

    U->>C: Request to register
    C->>A: Initiate registration process
    A->>U: Request user details
    U->>A: Submit user details
    A->>A: Validate and create user account
    A->>U: Confirm registration success
```

### Authentication

```mermaid
sequenceDiagram
participant U as User
participant C as Client
participant A as Authorization Server

    U->>C: Request to authenticate
    C->>A: Redirect user to Authorization Server
    A->>U: Request credentials
    U->>A: Submit credentials
    A->>A: Validate credentials
    A-->>U: Authentication success or failure
    alt Authentication Success
        A->>C: Redirect with access token
    else Authentication Failure
        A->>U: Show error message
    end
```

### Authroization with JWT

```mermaid
sequenceDiagram
participant U as User
participant C as Client
participant A as Authorization Server
participant R as Resource Server

    U->>C: Request access to protected resource
    alt JWT exists and is valid
        C->>R: Request resource with JWT
        R->>C: Provide resource
    else JWT not present or invalid
        C->>U: Redirect to Authorization Server
        U->>A: Authenticate and request JWT
        A-->>U: Provide JWT
        U->>C: Provide JWT to Client
        C->>R: Request resource with JWT
        R->>C: Provide resource
    end
```

## Usage

When updating a shared dependency, you can do so in it's own project directory and then run `go mod tidy` in each 
dependent project to ensure they are using the updated version.

