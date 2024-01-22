//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSCollectionViewDelegateFlowLayout
*/

type NSCollectionViewDelegateFlowLayout struct {

}

func (r *NSCollectionViewDelegateFlowLayout) CollectionView() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}
