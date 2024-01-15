//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextSearchOptions : objc.NSObject
*/

type UITextSearchOptions struct {
  *objc.NSObject

}

func (r *UITextSearchOptions) WordMatchMethod() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSearchOptions) StringCompareOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

