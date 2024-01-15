//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSException : objc.NSObject
*/

type NSException struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSException) CallStackReturnAddresses() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) CallStackSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) ExceptionWithName() (*NSException, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) Raise() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSException) Name() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) Reason() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSException) UserInfo() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

