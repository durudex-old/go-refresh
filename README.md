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

Generation of refresh token and receiving payload and full token.

```go
import (
	"fmt"

	"github.com/durudex/go-refresh"
)

func main() {
	r, err := refresh.New()
	if err != nil { ... }

	sessionId := "123"

	fmt.Println("Payload:", r.String())
	fmt.Println("Token:", r.Token(sessionId))
}
```

**Result:**
```
Payload: jeUiq2jueOdYGD1bKWaQHMRaGJQv4BlCC
Token: 123.jeUiq2jueOdYGD1bKWaQHMRaGJQv4BlCC
```
