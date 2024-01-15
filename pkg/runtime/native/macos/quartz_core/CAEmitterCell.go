//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface QuartzCore.CAEmitterCell : objc.NSObject
*/

type CAEmitterCell struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *CAMediaTiming
}

func (r *CAEmitterCell) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) BirthRate() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) Spin() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) RedSpeed() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) Velocity() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) GreenSpeed() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) ContentsRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) ContentsScale() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetSpin() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) DefaultValueForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) EmissionRange() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetZAcceleration() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) ScaleRange() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) BlueSpeed() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetMinificationFilterBias() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetEmitterCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) Lifetime() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) LifetimeRange() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetEmissionLongitude() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetVelocity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) MagnificationFilter() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) EmitterCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) EmissionLongitude() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) BlueRange() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetRedSpeed() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) Contents() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) MinificationFilter() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetMinificationFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetVelocityRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetYAcceleration() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetContentsRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetMagnificationFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetEmissionRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetXAcceleration() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SpinRange() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) GreenRange() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) EmitterCells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetScaleSpeed() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) RedRange() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) ScaleSpeed() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) Color() (*core_graphics.CGColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) XAcceleration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) ZAcceleration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetRedRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) AlphaRange() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetAlphaSpeed() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) MinificationFilterBias() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) EmissionLatitude() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) VelocityRange() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetScaleRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetBlueSpeed() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) Style() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetLifetimeRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetSpinRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetGreenRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetBlueRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) AlphaSpeed() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetContentsScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetLifetime() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetEmissionLatitude() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) Scale() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetAlphaRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetBirthRate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) YAcceleration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterCell) SetGreenSpeed() error {
  return fmt.Errorf("unimplemented")
}

