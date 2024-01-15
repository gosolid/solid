//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSCandidateListTouchBarItemDelegate
*/

type NSCandidateListTouchBarItemDelegate struct {

}

func (r *NSCandidateListTouchBarItemDelegate) CandidateListTouchBarItem() error {
  return fmt.Errorf("unimplemented")
}

