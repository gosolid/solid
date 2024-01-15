//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextContentManagerDelegate
*/

type NSTextContentManagerDelegate struct {

}

func (r *NSTextContentManagerDelegate) TextContentManager() (*NSTextElement, error) {
  return nil, fmt.Errorf("unimplemented")
}

