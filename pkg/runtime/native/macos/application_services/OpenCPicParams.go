//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.OpenCPicParams
*/

type OpenCPicParams struct {
  SrcRect objc.Rect `v8:"srcRect"`
  HRes int `v8:"hRes"`
  VRes int `v8:"vRes"`
  Version int16 `v8:"version"`
  Reserved1 int16 `v8:"reserved1"`
  Reserved2 int `v8:"reserved2"`
}
