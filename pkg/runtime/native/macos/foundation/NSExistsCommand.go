//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
interface Foundation.NSExistsCommand : Foundation.NSScriptCommand
*/

type NSExistsCommand struct {
  *NSScriptCommand

}

