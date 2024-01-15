//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSColorSampler : objc.NSObject
*/

type NSColorSampler struct {
  *objc.NSObject

}

func (r *NSColorSampler) ShowSamplerWithSelectionHandler() error {
  return fmt.Errorf("unimplemented")
}

