//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPositionalSpecifier : objc.NSObject
*/

type NSPositionalSpecifier struct {
  *objc.NSObject

}

func (r *NSPositionalSpecifier) SetInsertionClassDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPositionalSpecifier) Evaluate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPositionalSpecifier) ObjectSpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPositionalSpecifier) InsertionKey() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPositionalSpecifier) InsertionIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPositionalSpecifier) InitWithPosition() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPositionalSpecifier) Position() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPositionalSpecifier) InsertionContainer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPositionalSpecifier) InsertionReplaces() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

