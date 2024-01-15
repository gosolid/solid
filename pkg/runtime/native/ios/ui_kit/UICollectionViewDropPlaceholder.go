//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionViewDropPlaceholder : UIKit.UICollectionViewPlaceholder
*/

type UICollectionViewDropPlaceholder struct {
  *UICollectionViewPlaceholder

}

func (r *UICollectionViewDropPlaceholder) PreviewParametersProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDropPlaceholder) SetPreviewParametersProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

