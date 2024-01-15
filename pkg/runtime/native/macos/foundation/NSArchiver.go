//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSArchiver : Foundation.NSCoder
*/

type NSArchiver struct {
  *NSCoder

}

func (r *NSArchiver) ClassNameEncodedForTrueClassName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArchiver) ArchiverData() (*NSMutableData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArchiver) InitForWritingWithMutableData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArchiver) ArchivedDataWithRootObject() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArchiver) ArchiveRootObject() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArchiver) EncodeClassName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArchiver) EncodeRootObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArchiver) EncodeConditionalObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArchiver) ReplaceObject() error {
  return fmt.Errorf("unimplemented")
}

