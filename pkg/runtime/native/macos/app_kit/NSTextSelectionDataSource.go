//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextSelectionDataSource
*/

type NSTextSelectionDataSource struct {

}

func (r *NSTextSelectionDataSource) EnumerateContainerBoundariesFromLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionDataSource) EnumerateSubstringsFromLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionDataSource) BaseWritingDirectionAtLocation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionDataSource) EnumerateCaretOffsetsInLineFragmentAtLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionDataSource) LineFragmentRangeForPoint() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionDataSource) DocumentRange() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionDataSource) TextRangeForSelectionGranularity() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionDataSource) LocationFromLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionDataSource) OffsetFromLocation() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionDataSource) TextLayoutOrientationAtLocation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

