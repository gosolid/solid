//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSDiffableDataSourceSectionTransaction : objc.NSObject
*/

type NSDiffableDataSourceSectionTransaction struct {
  *objc.NSObject

}

func (r *NSDiffableDataSourceSectionTransaction) InitialSnapshot() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionTransaction) FinalSnapshot() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionTransaction) Difference() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionTransaction) SectionIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

