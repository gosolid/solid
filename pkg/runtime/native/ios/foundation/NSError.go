//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSError : objc.NSObject
*/

type NSError struct {
  *objc.NSObject

}

func (r *NSError) Code() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) LocalizedDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) RecoveryAttempter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) UnderlyingErrors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) ErrorWithDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) Domain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) LocalizedRecoveryOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) HelpAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) InitWithDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) SetUserInfoValueProviderForDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) UserInfoValueProviderForDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) LocalizedRecoverySuggestion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) LocalizedFailureReason() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

