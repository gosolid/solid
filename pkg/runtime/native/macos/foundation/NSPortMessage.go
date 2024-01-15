//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPortMessage : objc.NSObject
*/

type NSPortMessage struct {
  *objc.NSObject

}

func (r *NSPortMessage) SendBeforeDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPortMessage) Components() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPortMessage) ReceivePort() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPortMessage) SendPort() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPortMessage) Msgid() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPortMessage) SetMsgid() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPortMessage) InitWithSendPort() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

