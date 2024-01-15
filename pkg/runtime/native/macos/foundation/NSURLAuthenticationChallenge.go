//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLAuthenticationChallenge : objc.NSObject
*/

type NSURLAuthenticationChallenge struct {
  *objc.NSObject
  *NSSecureCoding
}

func (r *NSURLAuthenticationChallenge) ProposedCredential() (*NSURLCredential, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) PreviousFailureCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) FailureResponse() (*NSURLResponse, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) Error() (*NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) Sender() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) InitWithProtectionSpace() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) InitWithAuthenticationChallenge() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) ProtectionSpace() (*NSURLProtectionSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

