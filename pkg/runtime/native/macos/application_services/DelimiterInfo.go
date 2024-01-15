//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.DelimiterInfo
*/

type DelimiterInfo struct {
  StartDelimiter [2]byte `v8:"startDelimiter"`
  EndDelimiter [2]byte `v8:"endDelimiter"`
}
