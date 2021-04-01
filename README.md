# Go google.type Adapters

![ci](https://github.com/googleapis/go-type-adapters/workflows/ci/badge.svg)
![latest release](https://img.shields.io/github/v/release/googleapis/go-type-adapters)
![go version](https://img.shields.io/github/go-mod/go-version/googleapis/go-type-adapters)

This library provides helper functions for converting between the Golang
proto messages in `google.type` (as found in [genproto][]) and Golang native
types.

Full docs are at https://pkg.go.dev/github.com/googleapis/go-type-adapters.

### Example

As a simple example, this library can convert between a `google.type.Decimal`
([proto definition][], [Go docs][]) and a Golang [big.Float][]:

```go
import (
  "github.com/googleapis/go-type-adapters/adapters"
  dpb "google.golang.org/genproto/type/decimal"
)

func main() {
  decimal := &dpb.Decimal{Value: "12345.678"}
  flt, err := adapters.DecimalToFloat(decimal)
  if err != nil {
    panic(err)
  }
  // flt is a Golang *big.Float and can be used as such...
}
```

[genproto]: https://pkg.go.dev/google.golang.org/genproto
[proto definition]: https://github.com/googleapis/googleapis/blob/master/google/type/decimal.proto
[go docs]: https://pkg.go.dev/google.golang.org/genproto/googleapis/type/decimal
[big.float]: https://golang.org/pkg/math/big/#Float

## License

This software is made available under the [Apache 2.0][] license.

[apache 2.0]: https://www.apache.org/licenses/LICENSE-2.0
