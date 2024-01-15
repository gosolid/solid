//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.NCMConcatProfileSet
*/

type NCMConcatProfileSet struct {
  Cmm uint `v8:"cmm"`
  Flags uint `v8:"flags"`
  FlagsMask uint `v8:"flagsMask"`
  ProfileCount uint `v8:"profileCount"`
  ProfileSpecs [1]NCMConcatProfileSpec `v8:"profileSpecs"`
}
