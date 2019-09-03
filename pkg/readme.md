# readme

This is a sample about how to build a go pkg

# how to build a pkg
1. Env, if you not use `go module`, you could ignore it.
    - `go mod init github.com/<group>/<project>`
    - remember to fill the full path of git-project, if wrong, you will get wrong when you use `go mod download`.
2. Write your PKG content & test it `go test`.
3. git push & release it and use version as v1.0.0

# how to use it with go module

1. `go mod init {your_project_name}`
2. Edit the go.mod file. Refer this format
```
module {your_project_name}
require (
    github.com/jhaoheng/printhello v1.0.0
)
go 1.12	
```

3. `go mod download` & Use `printhello.Printhello()` in the main.go.
4. `go run main.go`

# 關於 autoload
當 pkg 有此封包時，在 main.go 中使用 `import _ "github.com/jhaoheng/printhello/autoload"`
可以在 init 時，自動調用 `github.com/jhaoheng/printhello`

## For Example

```
package main

import (
	_ "github.com/jhaoheng/printhello/autoload"
)

func main() {
}
```
