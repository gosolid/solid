//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIStoryboardPopoverSegue : UIKit.UIStoryboardSegue
*/

type UIStoryboardPopoverSegue struct {
  *UIStoryboardSegue

}

func (r *UIStoryboardPopoverSegue) PopoverController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

