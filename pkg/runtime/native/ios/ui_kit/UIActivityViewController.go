//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIActivityViewController : UIKit.UIViewController
*/

type UIActivityViewController struct {
  *UIViewController

}

func (r *UIActivityViewController) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) InitWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) InitWithActivityItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) CompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) CompletionWithItemsHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) SetCompletionWithItemsHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) SetExcludedActivityTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) SetCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) ExcludedActivityTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) AllowsProminentActivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityViewController) SetAllowsProminentActivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

