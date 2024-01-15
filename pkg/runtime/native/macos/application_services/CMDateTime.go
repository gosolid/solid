//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMDateTime
*/

type CMDateTime struct {
  Year uint16 `v8:"year"`
  Month uint16 `v8:"month"`
  DayOfTheMonth uint16 `v8:"dayOfTheMonth"`
  Hours uint16 `v8:"hours"`
  Minutes uint16 `v8:"minutes"`
  Seconds uint16 `v8:"seconds"`
}
