# Snowflake

Just an experimentation in making my own id generator for use as an identifier to be mapped with users, or for auth, that is short and readable.

The name was chosen as each individual snowflake is different from the other. So it fits with a unique identifier generator.

## Usage

### CLI

Get a new token

```bash
> snowflake
n4Xo9g.r6cguxp6
```

Get the timestamp of when the token was generated.

```bash
> snowflake -time n4Xo9g.r6cguxp6
1570762319
```

### Library

Generating a new token.

```go
import "github.com/Daegalus/snowflakes"

func main() {
  token, _err := snowflakes.Snowflake()
  fmt.Println(token)
}
```

Getting the time from the token.

```go
import "github.com/Daegalus/snowflakes"

func main() {
  token, _err := snowflakes.Snowflake()
  time, _err := snowflakes.GetTimeFromHash()
  fmt.Println(time.Unix())
}
```
