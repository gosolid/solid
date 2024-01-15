//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLAccelerationStructureMotionCurveGeometryDescriptor : Metal.MTLAccelerationStructureGeometryDescriptor
*/

type MTLAccelerationStructureMotionCurveGeometryDescriptor struct {
  *MTLAccelerationStructureGeometryDescriptor

}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetControlPointBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetControlPointStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetRadiusBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) RadiusStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) IndexType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) CurveBasis() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) ControlPointStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) ControlPointBuffers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) ControlPointCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) IndexBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) CurveType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetControlPointFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetRadiusFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetCurveType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SegmentCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SegmentControlPointCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetSegmentControlPointCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) CurveEndCaps() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) ControlPointFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetRadiusStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetCurveBasis() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) Descriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) RadiusBuffers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) IndexBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetIndexBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetIndexType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetSegmentCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetControlPointCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) RadiusFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetIndexBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionCurveGeometryDescriptor) SetCurveEndCaps() error {
  return fmt.Errorf("unimplemented")
}

