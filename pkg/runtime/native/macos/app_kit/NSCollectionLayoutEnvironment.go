//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSCollectionLayoutEnvironment
*/

type NSCollectionLayoutEnvironment struct {

}

func (r *NSCollectionLayoutEnvironment) Container() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

