//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextContentStorageDelegate
*/

type NSTextContentStorageDelegate struct {

}

func (r *NSTextContentStorageDelegate) TextContentStorage() (*NSTextParagraph, error) {
  return nil, fmt.Errorf("unimplemented")
}

