//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSLayoutYAxisAnchor : UIKit.NSLayoutAnchor
*/

type NSLayoutYAxisAnchor struct {
  *NSLayoutAnchor

}

func (r *NSLayoutYAxisAnchor) AnchorWithOffsetToAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

