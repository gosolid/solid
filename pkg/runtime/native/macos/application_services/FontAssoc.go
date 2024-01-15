//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.FontAssoc
*/

type FontAssoc struct {
  NumAssoc int16 `v8:"numAssoc"`
}
