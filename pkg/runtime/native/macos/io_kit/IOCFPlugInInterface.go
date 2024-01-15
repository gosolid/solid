//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOCFPlugInInterface
*/

type IOCFPlugInInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  Probe func(...any) (any, error) `v8:"probe"`
  Start func(...any) (any, error) `v8:"start"`
  Stop func(...any) (any, error) `v8:"stop"`
}
