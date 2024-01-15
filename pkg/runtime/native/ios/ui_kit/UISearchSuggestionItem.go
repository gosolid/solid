//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISearchSuggestionItem : objc.NSObject
*/

type UISearchSuggestionItem struct {
  *objc.NSObject

}

func (r *UISearchSuggestionItem) LocalizedAttributedSuggestion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchSuggestionItem) LocalizedSuggestion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchSuggestionItem) IconImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchSuggestionItem) RepresentedObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchSuggestionItem) SuggestionWithLocalizedSuggestion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchSuggestionItem) InitWithLocalizedSuggestion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchSuggestionItem) LocalizedDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchSuggestionItem) SetRepresentedObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchSuggestionItem) SuggestionWithLocalizedAttributedSuggestion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchSuggestionItem) InitWithLocalizedAttributedSuggestion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

