//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSDiffableDataSourceTransaction : objc.NSObject
*/

type NSDiffableDataSourceTransaction struct {
  *objc.NSObject

}

func (r *NSDiffableDataSourceTransaction) InitialSnapshot() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceTransaction) FinalSnapshot() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceTransaction) Difference() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceTransaction) SectionTransactions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

