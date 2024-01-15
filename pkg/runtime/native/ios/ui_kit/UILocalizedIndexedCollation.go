//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UILocalizedIndexedCollation : objc.NSObject
*/

type UILocalizedIndexedCollation struct {
  *objc.NSObject

}

func (r *UILocalizedIndexedCollation) SectionForSectionIndexTitleAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalizedIndexedCollation) SectionForObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalizedIndexedCollation) SortedArrayFromArray() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalizedIndexedCollation) SectionTitles() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalizedIndexedCollation) SectionIndexTitles() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILocalizedIndexedCollation) CurrentCollation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

