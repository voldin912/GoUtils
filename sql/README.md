# JSON utility package for SQL

[![Go Reference](https://pkg.go.dev/badge/github.com/kamiaka/go-jsonutil/sql.svg)](https://pkg.go.dev/github.com/kamiaka/go-jsonutil/sql)

JSON Utilities for Go-lang `encoding/json` package in SQL-based API response.

```go
var v *struct {
  Date      sql.Date
  Str       string
  IsDeleted sql.Bool
  CreatedAt sql.DateTime
  DeletedAt *sql.DateTime
}
json.Unmarshal([]byte(`{"Date":"2021-07-08","Str":"val","IsDeleted":0,"CreatedAt":"2021-07-07 12:34:56","DeletedAt":null}`), &v)
```