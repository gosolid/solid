//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMBitmap
*/

type CMBitmap struct {
  Image byte `v8:"image"`
  Width uint64 `v8:"width"`
  Height uint64 `v8:"height"`
  RowBytes uint64 `v8:"rowBytes"`
  PixelSize uint64 `v8:"pixelSize"`
  Space uint `v8:"space"`
  User1 uint `v8:"user1"`
  User2 uint `v8:"user2"`
}
