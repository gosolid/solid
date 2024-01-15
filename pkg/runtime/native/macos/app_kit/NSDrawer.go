//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSDrawer : AppKit.NSResponder
*/

type NSDrawer struct {
  *NSResponder
  *NSAccessibilityElement
  *NSAccessibility
}

func (r *NSDrawer) ContentView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) SetContentView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) SetContentSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) TrailingOffset() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) SetParentWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) SetLeadingOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) Open() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) Toggle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) State() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) InitWithContentSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) SetPreferredEdge() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) LeadingOffset() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) SetMinContentSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) MaxContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) SetTrailingOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) OpenOnEdge() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) PreferredEdge() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) ContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) MinContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) Close() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawer) ParentWindow() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) Edge() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDrawer) SetMaxContentSize() error {
  return fmt.Errorf("unimplemented")
}

