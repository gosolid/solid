//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFindSession : objc.NSObject
*/

type UIFindSession struct {
  *objc.NSObject

}

func (r *UIFindSession) PerformSearchWithQuery() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) ResultCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) SearchResultDisplayStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) SupportsReplacement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) HighlightedResultIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) SetSearchResultDisplayStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) AllowsReplacementForCurrentlyHighlightedResult() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) AllowsReplacement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) PerformSingleReplacementWithSearchQuery() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) ReplaceAllInstancesOfSearchQuery() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) HighlightNextResultInDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindSession) InvalidateFoundResults() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

