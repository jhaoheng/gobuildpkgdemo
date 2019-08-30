package autoload

/*
	You can just import this

		import _ "github.com/jhaoheng/printhello/autoload"

	And terminal will show "hello"
*/

import "github.com/jhaoheng/printhello"

func init() {
	printhello.Printhello()
}
