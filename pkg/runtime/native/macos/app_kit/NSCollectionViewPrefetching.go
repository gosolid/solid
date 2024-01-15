//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSCollectionViewPrefetching
*/

type NSCollectionViewPrefetching struct {

}

func (r *NSCollectionViewPrefetching) CollectionView() error {
  return fmt.Errorf("unimplemented")
}

