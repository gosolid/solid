//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionViewFocusUpdateContext : UIKit.UIFocusUpdateContext
*/

type UICollectionViewFocusUpdateContext struct {
  *UIFocusUpdateContext

}

func (r *UICollectionViewFocusUpdateContext) PreviouslyFocusedIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewFocusUpdateContext) NextFocusedIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

