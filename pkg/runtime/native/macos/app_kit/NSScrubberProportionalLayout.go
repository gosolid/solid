//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSScrubberProportionalLayout : AppKit.NSScrubberLayout
*/

type NSScrubberProportionalLayout struct {
  *NSScrubberLayout

}

func (r *NSScrubberProportionalLayout) NumberOfVisibleItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubberProportionalLayout) SetNumberOfVisibleItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberProportionalLayout) InitWithNumberOfVisibleItems() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberProportionalLayout) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

