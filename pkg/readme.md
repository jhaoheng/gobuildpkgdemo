# how to use this pkg
1. `go mod init <projectName>`
2. 在 main 中 `import "github.com/jhaoheng/gobuildpkgdemo/pkg"`
3. 使用 `pkg.Printhello()`
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
