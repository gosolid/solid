//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.PixPat
*/

type PixPat struct {
  PatType int16 `v8:"patType"`
  PatMap **PixMap `v8:"patMap"`
  PatData **byte `v8:"patData"`
  PatXData **byte `v8:"patXData"`
  PatXValid int16 `v8:"patXValid"`
  PatXMap **byte `v8:"patXMap"`
  Pat1Data Pattern `v8:"pat1Data"`
}
