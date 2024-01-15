//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFontPickerViewControllerConfiguration : objc.NSObject
*/

type UIFontPickerViewControllerConfiguration struct {
  *objc.NSObject

}

func (r *UIFontPickerViewControllerConfiguration) FilterPredicateForFilteredLanguages() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewControllerConfiguration) DisplayUsingSystemFont() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewControllerConfiguration) SetFilteredTraits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewControllerConfiguration) FilteredLanguagesPredicate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewControllerConfiguration) SetFilteredLanguagesPredicate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewControllerConfiguration) IncludeFaces() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewControllerConfiguration) SetIncludeFaces() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewControllerConfiguration) SetDisplayUsingSystemFont() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewControllerConfiguration) FilteredTraits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

