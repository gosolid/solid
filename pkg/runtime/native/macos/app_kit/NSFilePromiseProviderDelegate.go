//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSFilePromiseProviderDelegate
*/

type NSFilePromiseProviderDelegate struct {

}

func (r *NSFilePromiseProviderDelegate) FilePromiseProvider() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseProviderDelegate) OperationQueueForFilePromiseProvider() (*foundation.NSOperationQueue, error) {
  return nil, fmt.Errorf("unimplemented")
}

