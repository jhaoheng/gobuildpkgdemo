package autoload

/*
	You can just read the .env file on import just by doing

		import _ "github.com/jhaoheng/printhello/autoload"

	And bob's your mother's brother
*/

import "github.com/jhaoheng/printhello"

func init() {
	printhello.Printhello()
}