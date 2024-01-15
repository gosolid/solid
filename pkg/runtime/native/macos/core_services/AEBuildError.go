//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AEBuildError
*/

type AEBuildError struct {
  FError uint `v8:"fError"`
  FErrorPos uint `v8:"fErrorPos"`
}
