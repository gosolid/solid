//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLAuthenticationChallengeSender
*/

type NSURLAuthenticationChallengeSender struct {

}

func (r *NSURLAuthenticationChallengeSender) RejectProtectionSpaceAndContinueWithChallenge() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallengeSender) UseCredential() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallengeSender) ContinueWithoutCredentialForAuthenticationChallenge() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallengeSender) CancelAuthenticationChallenge() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLAuthenticationChallengeSender) PerformDefaultHandlingForAuthenticationChallenge() error {
  return fmt.Errorf("unimplemented")
}

