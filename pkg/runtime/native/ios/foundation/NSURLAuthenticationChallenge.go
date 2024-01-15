//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLAuthenticationChallenge : objc.NSObject
*/

type NSURLAuthenticationChallenge struct {
  *objc.NSObject

}

func (r *NSURLAuthenticationChallenge) InitWithProtectionSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) InitWithAuthenticationChallenge() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) ProtectionSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) ProposedCredential() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) PreviousFailureCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) FailureResponse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) Error() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallenge) Sender() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

