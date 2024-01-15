//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITextSearchingFindSession : UIKit.UIFindSession
*/

type UITextSearchingFindSession struct {
  *UIFindSession

}

func (r *UITextSearchingFindSession) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSearchingFindSession) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSearchingFindSession) SearchableObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSearchingFindSession) InitWithSearchableObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

