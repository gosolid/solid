//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSLayoutXAxisAnchor : UIKit.NSLayoutAnchor
*/

type NSLayoutXAxisAnchor struct {
  *NSLayoutAnchor

}

func (r *NSLayoutXAxisAnchor) AnchorWithOffsetToAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

