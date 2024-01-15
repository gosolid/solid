//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.NumFormatString
*/

type NumFormatString struct {
  FLength byte `v8:"fLength"`
  FVersion byte `v8:"fVersion"`
  Data [254]byte `v8:"data"`
}
