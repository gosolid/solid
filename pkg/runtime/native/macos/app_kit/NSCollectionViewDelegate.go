//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSCollectionViewDelegate
*/

type NSCollectionViewDelegate struct {

}

func (r *NSCollectionViewDelegate) CollectionView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

