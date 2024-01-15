//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UILexiconEntry : objc.NSObject
*/

type UILexiconEntry struct {
  *objc.NSObject

}

func (r *UILexiconEntry) DocumentText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILexiconEntry) UserInput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

