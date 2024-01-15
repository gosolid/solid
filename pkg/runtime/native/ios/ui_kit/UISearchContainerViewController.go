//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UISearchContainerViewController : UIKit.UIViewController
*/

type UISearchContainerViewController struct {
  *UIViewController

}

func (r *UISearchContainerViewController) InitWithSearchController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISearchContainerViewController) SearchController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

