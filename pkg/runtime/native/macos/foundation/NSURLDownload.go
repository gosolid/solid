//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLDownload : objc.NSObject
*/

type NSURLDownload struct {
  *objc.NSObject

}

func (r *NSURLDownload) Cancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLDownload) SetDestination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLDownload) Request() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLDownload) DeletesFileUponFailure() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLDownload) InitWithRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLDownload) InitWithResumeData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLDownload) ResumeData() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLDownload) SetDeletesFileUponFailure() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLDownload) CanResumeDownloadDecodedWithEncodingMIMEType() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

