// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package nat_gateway

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.AllocationID, b.ko.Spec.AllocationID) {
		delta.Add("Spec.AllocationID", a.ko.Spec.AllocationID, b.ko.Spec.AllocationID)
	} else if a.ko.Spec.AllocationID != nil && b.ko.Spec.AllocationID != nil {
		if *a.ko.Spec.AllocationID != *b.ko.Spec.AllocationID {
			delta.Add("Spec.AllocationID", a.ko.Spec.AllocationID, b.ko.Spec.AllocationID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ClientToken, b.ko.Spec.ClientToken) {
		delta.Add("Spec.ClientToken", a.ko.Spec.ClientToken, b.ko.Spec.ClientToken)
	} else if a.ko.Spec.ClientToken != nil && b.ko.Spec.ClientToken != nil {
		if *a.ko.Spec.ClientToken != *b.ko.Spec.ClientToken {
			delta.Add("Spec.ClientToken", a.ko.Spec.ClientToken, b.ko.Spec.ClientToken)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ConnectivityType, b.ko.Spec.ConnectivityType) {
		delta.Add("Spec.ConnectivityType", a.ko.Spec.ConnectivityType, b.ko.Spec.ConnectivityType)
	} else if a.ko.Spec.ConnectivityType != nil && b.ko.Spec.ConnectivityType != nil {
		if *a.ko.Spec.ConnectivityType != *b.ko.Spec.ConnectivityType {
			delta.Add("Spec.ConnectivityType", a.ko.Spec.ConnectivityType, b.ko.Spec.ConnectivityType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SubnetID, b.ko.Spec.SubnetID) {
		delta.Add("Spec.SubnetID", a.ko.Spec.SubnetID, b.ko.Spec.SubnetID)
	} else if a.ko.Spec.SubnetID != nil && b.ko.Spec.SubnetID != nil {
		if *a.ko.Spec.SubnetID != *b.ko.Spec.SubnetID {
			delta.Add("Spec.SubnetID", a.ko.Spec.SubnetID, b.ko.Spec.SubnetID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.TagSpecifications, b.ko.Spec.TagSpecifications) {
		delta.Add("Spec.TagSpecifications", a.ko.Spec.TagSpecifications, b.ko.Spec.TagSpecifications)
	}

	return delta
}
