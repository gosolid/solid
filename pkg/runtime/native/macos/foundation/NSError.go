//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSError : objc.NSObject
*/

type NSError struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSError) Code() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSError) UnderlyingErrors() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) UserInfo() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) LocalizedRecoveryOptions() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) SetUserInfoValueProviderForDomain() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSError) UserInfoValueProviderForDomain() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) Domain() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) LocalizedDescription() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) RecoveryAttempter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) HelpAnchor() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) InitWithDomain() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) ErrorWithDomain() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) LocalizedFailureReason() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSError) LocalizedRecoverySuggestion() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

