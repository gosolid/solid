//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UILexicon : objc.NSObject
*/

type UILexicon struct {
  *objc.NSObject

}

func (r *UILexicon) Entries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

