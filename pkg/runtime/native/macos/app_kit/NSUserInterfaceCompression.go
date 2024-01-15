//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSUserInterfaceCompression
*/

type NSUserInterfaceCompression struct {

}

func (r *NSUserInterfaceCompression) CompressWithPrioritizedCompressionOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompression) MinimumSizeWithPrioritizedCompressionOptions() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceCompression) ActiveCompressionOptions() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

