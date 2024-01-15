//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
interface Foundation.NSConstantString : Foundation.NSSimpleCString
*/

type NSConstantString struct {
  *NSSimpleCString

}

