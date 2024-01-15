//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.DateTimeRec
*/

type DateTimeRec struct {
  Year int16 `v8:"year"`
  Month int16 `v8:"month"`
  Day int16 `v8:"day"`
  Hour int16 `v8:"hour"`
  Minute int16 `v8:"minute"`
  Second int16 `v8:"second"`
  DayOfWeek int16 `v8:"dayOfWeek"`
}
