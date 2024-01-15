//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.LocaleAndVariant
*/

type LocaleAndVariant struct {
  Locale *OpaqueLocaleRef `v8:"locale"`
  OpVariant uint `v8:"opVariant"`
}
