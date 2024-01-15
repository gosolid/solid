//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreLocation.CLLocationCoordinate2D
*/

type CLLocationCoordinate2D struct {
  Latitude float64 `v8:"latitude"`
  Longitude float64 `v8:"longitude"`
}
