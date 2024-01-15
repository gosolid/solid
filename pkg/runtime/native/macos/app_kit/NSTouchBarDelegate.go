//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTouchBarDelegate
*/

type NSTouchBarDelegate struct {

}

func (r *NSTouchBarDelegate) TouchBar() (*NSTouchBarItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

