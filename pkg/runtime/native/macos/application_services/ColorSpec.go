//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ColorSpec
*/

type ColorSpec struct {
  Value int16 `v8:"value"`
  Rgb RGBColor `v8:"rgb"`
}
