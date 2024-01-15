//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMFloatBitmap
*/

type CMFloatBitmap struct {
  Version uint64 `v8:"version"`
  Buffers [16]*float32 `v8:"buffers"`
  Height uint64 `v8:"height"`
  Width uint64 `v8:"width"`
  RowStride int64 `v8:"rowStride"`
  ColStride int64 `v8:"colStride"`
  Space uint `v8:"space"`
  Flags int `v8:"flags"`
}
