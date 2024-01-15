//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSCollectionLayoutContainer
*/

type NSCollectionLayoutContainer struct {

}

func (r *NSCollectionLayoutContainer) EffectiveContentInsets() (NSDirectionalEdgeInsets, error) {
  return NSDirectionalEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutContainer) ContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutContainer) EffectiveContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutContainer) ContentInsets() (NSDirectionalEdgeInsets, error) {
  return NSDirectionalEdgeInsets{}, fmt.Errorf("unimplemented")
}

