//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.Picture
*/

type Picture struct {
  PicSize int16 `v8:"picSize"`
  PicFrame objc.Rect `v8:"picFrame"`
}
