//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSCollectionViewFlowLayoutInvalidationContext : AppKit.NSCollectionViewLayoutInvalidationContext
*/

type NSCollectionViewFlowLayoutInvalidationContext struct {
  *NSCollectionViewLayoutInvalidationContext

}

func (r *NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes() error {
  return fmt.Errorf("unimplemented")
}

