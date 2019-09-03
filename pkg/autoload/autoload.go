package autoload

/*
	You can just import this

		import _ "github.com/jhaoheng/printhello/pkg/autoload"

	And terminal will show "hello"
*/

import "github.com/jhaoheng/gobuildpkgdemo/pkg"

func init() {
	pkg.Printhello()
}
