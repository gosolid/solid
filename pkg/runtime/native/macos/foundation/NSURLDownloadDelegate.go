//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLDownloadDelegate
*/

type NSURLDownloadDelegate struct {

}

func (r *NSURLDownloadDelegate) DownloadDidFinish() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLDownloadDelegate) DownloadDidBegin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLDownloadDelegate) Download() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLDownloadDelegate) DownloadShouldUseCredentialStorage() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

