//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSTextTableBlock : AppKit.NSTextBlock
*/

type NSTextTableBlock struct {
  *NSTextBlock

}

func (r *NSTextTableBlock) InitWithTable() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextTableBlock) Table() (*NSTextTable, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextTableBlock) StartingRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextTableBlock) RowSpan() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextTableBlock) StartingColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextTableBlock) ColumnSpan() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

