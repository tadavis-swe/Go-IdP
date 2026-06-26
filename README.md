# IdP — OAuth2 / OpenID Connect Identity Provider (Go)
![Go Version](https://img.shields.io/github/go-mod/go-version/tadavis-swe/Go-IdP)


This project is a lightweight OAuth2 / OIDC Identity Provider written in Go.  
It implements the Authorization Code flow end‑to‑end, including:

- `/authorize` — issues authorization codes  
- `/token` — exchanges codes for ID & access tokens  
- `/jwks.json` — exposes public signing keys  
- `/.well-known/openid-configuration` — OIDC discovery document  
- `testclient/` — minimal OAuth client for local testing

## Features

- Pure Go implementation (no external frameworks)
- RSA‑signed ID tokens (RS256)
- In‑memory authorization code store
- Minimal, readable code structure
- Test client demonstrating the full OAuth2 flow

## Project Structure

cmd/
server/        # Main IdP server
internal/
auth/          # Token signing, auth code storage
http/          # Handlers for OAuth2 endpoints
testclient/      # Simple OAuth client for local testing


## Running the Server

```bash
go run ./cmd/server
```
- Server runs on port 8080

## Running the Test client

```bash
go run ./testclient
```
- Client listens on 3000 and prints received auth codes anad tokens

## OAuth2 Flow (Local)
1. Start both the IdP server and test client
2. Visit
```bash
http://localhost:8080/authorize?client_id=test&redirect_uri=http://localhost:3000/cb&response_type=code&state=123
```
3. The test client receives the code and exchanges it for tokens.

## License

MIT
