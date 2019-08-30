# readme

This is a sample about how to build a go pkg

# how to build a pkg

1. Create a git, your pkg name "printhello"
    - The Git name does not force has to the same the PKG, but for clarity,  we make it the same.
2. Write your PKG content
    - Remember to need the capital letter.
    - Don't forget to test it.
3. `git push` & release it and use version as v1.0.0

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