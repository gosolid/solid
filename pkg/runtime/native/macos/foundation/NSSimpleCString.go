//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
interface Foundation.NSSimpleCString : Foundation.NSString
*/

type NSSimpleCString struct {
  *NSString

}

