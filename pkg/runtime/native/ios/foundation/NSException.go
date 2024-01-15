//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSException : objc.NSObject
*/

type NSException struct {
  *objc.NSObject

}

func (r *NSException) Raise() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) Reason() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) CallStackReturnAddresses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) CallStackSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) ExceptionWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) InitWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

