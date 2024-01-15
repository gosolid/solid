//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.IconFamilyElement
*/

type IconFamilyElement struct {
  ElementType uint `v8:"elementType"`
  ElementSize int `v8:"elementSize"`
  ElementData [1]byte `v8:"elementData"`
}
