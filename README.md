# Antelope Leap HTTP Api

[Leap HTTP API](https://developers.eos.io/welcome/latest/reference/index) and [hyperion API](https://hyperion.docs.eosrio.io) written in go.

### Install package

``` bash
go get -u github.com/eosswedenorg-go/leapapi@latest
```

### Types

API Request parameters struct

```go
type ReqParams struct {
	Url string
	Host string
}
```

chain info struct

```go
type Info struct {
	ServerVersion string
	HeadBlockNum int64
	HeadBlockTime time.Time
	HTTPStatusCode int
}
```

Hyperion health struct (not all fields).

```go
type Health struct {
	VersionHash string
	Health []Service
	HTTPStatusCode int
}
```

Service struct from Hyperion health
```go
type Service struct {
	Name string
	Status string
	Data map[string]interface{}
	Time int64 // unix timestamp.
}
```

### Functions

```go
func GetInfo(params ReqParams) (Info, error)
```

Call `v1/chain/get_info` and return the results.

```go
func GetHealth(params ReqParams) (Health, error)
```

Call `v2/health` (hyperion) and return the results.

### Author

Henrik Hautakoski - [Sw/eden](https://eossweden.org/) - [henrik@eossweden.org](mailto:henrik@eossweden.org)
