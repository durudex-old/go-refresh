<div align="center">
    <a href="https://discord.gg/4qcXbeVehZ">
        <img alt="Discord" src="https://img.shields.io/discord/882288646517035028?label=%F0%9F%92%AC%20discord">
    </a>
    <a href="https://github.com/durudex/go-refresh/blob/main/LICENSE">
        <img alt="License" src="https://img.shields.io/github/license/durudex/go-refresh?label=%F0%9F%93%95%20license">
    </a>
    <a href="https://github.com/durudex/go-refresh/stargazers">
        <img alt="GitHub Stars" src="https://img.shields.io/github/stars/durudex/go-refresh?label=%E2%AD%90%20stars&logo=sdf">
    </a>
    <a href="https://github.com/durudex/go-refresh/network">
        <img alt="GitHub Forks" src="https://img.shields.io/github/forks/durudex/go-refresh?label=%F0%9F%93%81%20forks">
    </a>
</div>

<h1 align="center">Go Refresh</h1>

<p align="center">Durudex Refresh Token implementation.</p>

## Setup

```
go get github.com/durudex/go-refresh
```

## Usage

Generation of refresh token and receiving payload and full token:

```go
import (
	"fmt"

	"github.com/durudex/go-refresh"
)

const (
	sessionId = "2Fj7R4ERGtDdM5SBGl6mDesP2Qm"
	objectId = "2Fj7R4WcK9KMn1FY2kCnKGbFtA0"
)

func main() {
	r, err := refresh.New()
	if err != nil { ... }

	fmt.Println("Payload:", r.String())
	fmt.Println("Token:", r.Token(sessionId, objectId))
}
```

**Result:**

```
Payload: ZpgdubxwoCROmgHJ0g8EzFhFsy5KopL4
Token: 2Fj7R4ERGtDdM5SBGl6mDesP2Qm.2Fj7R4WcK9KMn1FY2kCnKGbFtA0.ZpgdubxwoCROmgHJ0g8EzFhFsy5KopL4
```

Hashing of the refresh token using a secret key:

```go
import (
	"fmt"

	"github.com/durudex/go-refresh"
)

const secretKey = "durudex"

func main() {
	r, err := refresh.New()
	if err != nil { ... }

	h := r.Hash([]byte(secretKey))

	fmt.Println("Hash:", fmt.Sprintf("%x", h))
}
```

**Result:**

```
Hash: 91b9b4ddda35be0338407fbaa76bb6adfe2dba8ad6719fe0ebae006c297b529f
```

Parsing refresh token from string:

```go
import (
	"fmt"

	"github.com/durudex/go-refresh"
)

const fullToken = "2Fj7R4ERGtDdM5SBGl6mDesP2Qm.2Fj7R4WcK9KMn1FY2kCnKGbFtA0.ZpgdubxwoCROmgHJ0g8EzFhFsy5KopL4"

func main() {
	token, err := refresh.Parse(fullToken)
	if err != nil { ... }

	fmt.Println("Payload:", token.Payload.String())
	fmt.Println("Session Id:", token.Session)
	fmt.Println("Object Id:", token.Object)
}
```

**Result:**

```
Payload: ZpgdubxwoCROmgHJ0g8EzFhFsy5KopL4
Session Id: 2Fj7R4ERGtDdM5SBGl6mDesP2Qm
Object Id: 2Fj7R4WcK9KMn1FY2kCnKGbFtA0
```
