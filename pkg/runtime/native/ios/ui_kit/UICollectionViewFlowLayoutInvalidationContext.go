//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionViewFlowLayoutInvalidationContext : UIKit.UICollectionViewLayoutInvalidationContext
*/

type UICollectionViewFlowLayoutInvalidationContext struct {
  *UICollectionViewLayoutInvalidationContext

}

func (r *UICollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

