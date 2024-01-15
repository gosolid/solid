//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionReusableView : UIKit.UIView
*/

type UICollectionReusableView struct {
  *UIView

}

func (r *UICollectionReusableView) PrepareForReuse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionReusableView) ApplyLayoutAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionReusableView) WillTransitionFromLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionReusableView) DidTransitionFromLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionReusableView) PreferredLayoutAttributesFittingAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionReusableView) ReuseIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

