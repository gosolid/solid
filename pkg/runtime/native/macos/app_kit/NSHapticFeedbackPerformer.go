//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSHapticFeedbackPerformer
*/

type NSHapticFeedbackPerformer struct {

}

func (r *NSHapticFeedbackPerformer) PerformFeedbackPattern() error {
  return fmt.Errorf("unimplemented")
}
