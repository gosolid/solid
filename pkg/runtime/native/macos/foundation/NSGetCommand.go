//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
interface Foundation.NSGetCommand : Foundation.NSScriptCommand
*/

type NSGetCommand struct {
  *NSScriptCommand

}

