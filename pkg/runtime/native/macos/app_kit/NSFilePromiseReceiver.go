//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSFilePromiseReceiver : objc.NSObject
*/

type NSFilePromiseReceiver struct {
  *objc.NSObject
  *NSPasteboardReading
}

func (r *NSFilePromiseReceiver) FileTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseReceiver) FileNames() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseReceiver) ReceivePromisedFilesAtDestination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFilePromiseReceiver) ReadableDraggedTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

