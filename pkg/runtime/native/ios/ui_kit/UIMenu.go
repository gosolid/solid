//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIMenu : UIKit.UIMenuElement
*/

type UIMenu struct {
  *UIMenuElement

}

func (r *UIMenu) Children() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) SelectedElements() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) MenuByReplacingChildren() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) PreferredElementSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) SetPreferredElementSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) MenuWithChildren() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) MenuWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenu) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

