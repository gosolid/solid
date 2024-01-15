//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSTextContent
*/

type NSTextContent struct {

}

func (r *NSTextContent) ContentType() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContent) SetContentType() error {
  return fmt.Errorf("unimplemented")
}

