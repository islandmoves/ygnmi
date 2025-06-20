// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package uexampleocpath is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: (devel): (ygot: v0.32.0)
using the following YANG input files:
  - ../../pathgen/testdata/yang/openconfig-simple.yang
  - ../../pathgen/testdata/yang/openconfig-withlistval.yang
  - ../../pathgen/testdata/yang/openconfig-nested.yang

Imported modules were sourced from:
*/
package uexampleocpath

import (
	oc "github.com/openconfig/ygnmi/internal/uexampleoc"
	"github.com/openconfig/ygnmi/internal/uexampleoc/nested"
	"github.com/openconfig/ygnmi/internal/uexampleoc/simple"
	"github.com/openconfig/ygnmi/internal/uexampleoc/withlistval"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// DevicePath represents the /device YANG schema element.
type DevicePath struct {
	*ygnmi.DeviceRootBase
}

// Root returns a root path object from which YANG paths can be constructed.
func Root() *DevicePath {
	return &DevicePath{ygnmi.NewDeviceRootBase()}
}

// A (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "a"
//	Path from root:       "/a"
func (n *DevicePath) A() *nested.OpenconfigNested_APath {
	ps := &nested.OpenconfigNested_APath{
		NodePath: ygnmi.NewNodePath(
			[]string{"a"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_A](
		"OpenconfigNested_A",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// Container (container):
//
//	Defining module:      "openconfig-nested"
//	Instantiating module: "openconfig-nested"
//	Path from parent:     "container"
//	Path from root:       "/container"
func (n *DevicePath) Container() *nested.OpenconfigNested_ContainerPath {
	ps := &nested.OpenconfigNested_ContainerPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"container"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigNested_Container](
		"OpenconfigNested_Container",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// Model (container):
//
//	Defining module:      "openconfig-withlistval"
//	Instantiating module: "openconfig-withlistval"
//	Path from parent:     "model"
//	Path from root:       "/model"
func (n *DevicePath) Model() *withlistval.OpenconfigWithlistval_ModelPath {
	ps := &withlistval.OpenconfigWithlistval_ModelPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"model"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigWithlistval_Model](
		"OpenconfigWithlistval_Model",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// Parent (container): I am a parent container
// that has 4 children.
//
//	Defining module:      "openconfig-simple"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "parent"
//	Path from root:       "/parent"
func (n *DevicePath) Parent() *simple.OpenconfigSimple_ParentPath {
	ps := &simple.OpenconfigSimple_ParentPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"parent"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigSimple_Parent](
		"OpenconfigSimple_Parent",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// RemoteContainer (container):
//
//	Defining module:      "openconfig-remote"
//	Instantiating module: "openconfig-simple"
//	Path from parent:     "remote-container"
//	Path from root:       "/remote-container"
func (n *DevicePath) RemoteContainer() *simple.OpenconfigSimple_RemoteContainerPath {
	ps := &simple.OpenconfigSimple_RemoteContainerPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"remote-container"},
			map[string]interface{}{},
			n,
		),
	}
	ps.ConfigQuery = ygnmi.NewConfigQuery[*oc.OpenconfigSimple_RemoteContainer](
		"OpenconfigSimple_RemoteContainer",
		true,
		false,
		false,
		false,
		false,
		false,
		ps,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)

	return ps
}

// Batch contains a collection of paths.
// Use batch to call Lookup, Watch, etc. on multiple paths at once.
type Batch struct {
	paths []ygnmi.PathStruct
}

// AddPaths adds the paths to the batch.
func (b *Batch) AddPaths(paths ...ygnmi.PathStruct) *Batch {
	b.paths = append(b.paths, paths...)
	return b
}

// Query returns a Query that can be used in gNMI operations.
// The returned query is immutable, adding paths does not modify existing queries.
func (b *Batch) Query() ygnmi.SingletonQuery[*oc.Device] {
	queryPaths := make([]ygnmi.PathStruct, len(b.paths))
	copy(queryPaths, b.paths)
	return ygnmi.NewSingletonQuery[*oc.Device](
		"Device",
		true,
		false,
		false,
		false,
		false,
		false,
		ygnmi.NewDeviceRootBase(),
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Device{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		queryPaths,
		nil,
	)
}

func binarySliceToFloatSlice(in []oc.Binary) []float32 {
	converted := make([]float32, 0, len(in))
	for _, binary := range in {
		converted = append(converted, ygot.BinaryToFloat32(binary))
	}
	return converted
}
