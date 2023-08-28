//js:package assert

package assert

//go:generate go run github.com/grexie/isolates/codegen

type Import interface{}

//js:method
//js:export notStrictEqual
func NotStrictEqual(actual any, expected any, message ...any) error {
	return nil
}

//js:method assert
//js:export assert
//js:export default
func Assert() error {
	return nil
}
