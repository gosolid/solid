//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.MacPolygon
*/

type MacPolygon struct {
  PolySize int16 `v8:"polySize"`
  PolyBBox objc.Rect `v8:"polyBBox"`
  PolyPoints [1]objc.Point `v8:"polyPoints"`
}
