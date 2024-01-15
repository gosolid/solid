//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ScriptCodeRun
*/

type ScriptCodeRun struct {
  Offset uint64 `v8:"offset"`
  Script int16 `v8:"script"`
}
