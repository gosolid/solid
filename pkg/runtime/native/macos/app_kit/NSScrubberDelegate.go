//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSScrubberDelegate
*/

type NSScrubberDelegate struct {

}

func (r *NSScrubberDelegate) Scrubber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberDelegate) DidBeginInteractingWithScrubber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberDelegate) DidFinishInteractingWithScrubber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberDelegate) DidCancelInteractingWithScrubber() error {
  return fmt.Errorf("unimplemented")
}

