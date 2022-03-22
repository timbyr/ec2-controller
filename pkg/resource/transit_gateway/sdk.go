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

package transit_gateway

import (
	"context"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/ec2-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.EC2{}
	_ = &svcapitypes.TransitGateway{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeTransitGatewaysOutput
	resp, err = rm.sdkapi.DescribeTransitGatewaysWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeTransitGateways", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.TransitGateways {
		if elem.CreationTime != nil {
			ko.Status.CreationTime = &metav1.Time{*elem.CreationTime}
		} else {
			ko.Status.CreationTime = nil
		}
		if elem.Description != nil {
			ko.Spec.Description = elem.Description
		} else {
			ko.Spec.Description = nil
		}
		if elem.Options != nil {
			f2 := &svcapitypes.TransitGatewayRequestOptions{}
			if elem.Options.AmazonSideAsn != nil {
				f2.AmazonSideASN = elem.Options.AmazonSideAsn
			}
			if elem.Options.AutoAcceptSharedAttachments != nil {
				f2.AutoAcceptSharedAttachments = elem.Options.AutoAcceptSharedAttachments
			}
			if elem.Options.DefaultRouteTableAssociation != nil {
				f2.DefaultRouteTableAssociation = elem.Options.DefaultRouteTableAssociation
			}
			if elem.Options.DefaultRouteTablePropagation != nil {
				f2.DefaultRouteTablePropagation = elem.Options.DefaultRouteTablePropagation
			}
			if elem.Options.DnsSupport != nil {
				f2.DNSSupport = elem.Options.DnsSupport
			}
			if elem.Options.MulticastSupport != nil {
				f2.MulticastSupport = elem.Options.MulticastSupport
			}
			if elem.Options.TransitGatewayCidrBlocks != nil {
				f2f8 := []*string{}
				for _, f2f8iter := range elem.Options.TransitGatewayCidrBlocks {
					var f2f8elem string
					f2f8elem = *f2f8iter
					f2f8 = append(f2f8, &f2f8elem)
				}
				f2.TransitGatewayCIDRBlocks = f2f8
			}
			if elem.Options.VpnEcmpSupport != nil {
				f2.VPNECMPSupport = elem.Options.VpnEcmpSupport
			}
			ko.Spec.Options = f2
		} else {
			ko.Spec.Options = nil
		}
		if elem.OwnerId != nil {
			ko.Status.OwnerID = elem.OwnerId
		} else {
			ko.Status.OwnerID = nil
		}
		if elem.State != nil {
			ko.Status.State = elem.State
		} else {
			ko.Status.State = nil
		}
		if elem.Tags != nil {
			f5 := []*svcapitypes.Tag{}
			for _, f5iter := range elem.Tags {
				f5elem := &svcapitypes.Tag{}
				if f5iter.Key != nil {
					f5elem.Key = f5iter.Key
				}
				if f5iter.Value != nil {
					f5elem.Value = f5iter.Value
				}
				f5 = append(f5, f5elem)
			}
			ko.Status.Tags = f5
		} else {
			ko.Status.Tags = nil
		}
		if elem.TransitGatewayArn != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.TransitGatewayArn)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.TransitGatewayId != nil {
			ko.Status.TransitGatewayID = elem.TransitGatewayId
		} else {
			ko.Status.TransitGatewayID = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return r.ko.Status.TransitGatewayID == nil

}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeTransitGatewaysInput, error) {
	res := &svcsdk.DescribeTransitGatewaysInput{}

	if r.ko.Status.TransitGatewayID != nil {
		f4 := []*string{}
		f4 = append(f4, r.ko.Status.TransitGatewayID)
		res.SetTransitGatewayIds(f4)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateTransitGatewayOutput
	_ = resp
	resp, err = rm.sdkapi.CreateTransitGatewayWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateTransitGateway", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.TransitGateway.CreationTime != nil {
		ko.Status.CreationTime = &metav1.Time{*resp.TransitGateway.CreationTime}
	} else {
		ko.Status.CreationTime = nil
	}
	if resp.TransitGateway.Description != nil {
		ko.Spec.Description = resp.TransitGateway.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.TransitGateway.Options != nil {
		f2 := &svcapitypes.TransitGatewayRequestOptions{}
		if resp.TransitGateway.Options.AmazonSideAsn != nil {
			f2.AmazonSideASN = resp.TransitGateway.Options.AmazonSideAsn
		}
		if resp.TransitGateway.Options.AutoAcceptSharedAttachments != nil {
			f2.AutoAcceptSharedAttachments = resp.TransitGateway.Options.AutoAcceptSharedAttachments
		}
		if resp.TransitGateway.Options.DefaultRouteTableAssociation != nil {
			f2.DefaultRouteTableAssociation = resp.TransitGateway.Options.DefaultRouteTableAssociation
		}
		if resp.TransitGateway.Options.DefaultRouteTablePropagation != nil {
			f2.DefaultRouteTablePropagation = resp.TransitGateway.Options.DefaultRouteTablePropagation
		}
		if resp.TransitGateway.Options.DnsSupport != nil {
			f2.DNSSupport = resp.TransitGateway.Options.DnsSupport
		}
		if resp.TransitGateway.Options.MulticastSupport != nil {
			f2.MulticastSupport = resp.TransitGateway.Options.MulticastSupport
		}
		if resp.TransitGateway.Options.TransitGatewayCidrBlocks != nil {
			f2f8 := []*string{}
			for _, f2f8iter := range resp.TransitGateway.Options.TransitGatewayCidrBlocks {
				var f2f8elem string
				f2f8elem = *f2f8iter
				f2f8 = append(f2f8, &f2f8elem)
			}
			f2.TransitGatewayCIDRBlocks = f2f8
		}
		if resp.TransitGateway.Options.VpnEcmpSupport != nil {
			f2.VPNECMPSupport = resp.TransitGateway.Options.VpnEcmpSupport
		}
		ko.Spec.Options = f2
	} else {
		ko.Spec.Options = nil
	}
	if resp.TransitGateway.OwnerId != nil {
		ko.Status.OwnerID = resp.TransitGateway.OwnerId
	} else {
		ko.Status.OwnerID = nil
	}
	if resp.TransitGateway.State != nil {
		ko.Status.State = resp.TransitGateway.State
	} else {
		ko.Status.State = nil
	}
	if resp.TransitGateway.Tags != nil {
		f5 := []*svcapitypes.Tag{}
		for _, f5iter := range resp.TransitGateway.Tags {
			f5elem := &svcapitypes.Tag{}
			if f5iter.Key != nil {
				f5elem.Key = f5iter.Key
			}
			if f5iter.Value != nil {
				f5elem.Value = f5iter.Value
			}
			f5 = append(f5, f5elem)
		}
		ko.Status.Tags = f5
	} else {
		ko.Status.Tags = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.TransitGateway.TransitGatewayArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.TransitGateway.TransitGatewayArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.TransitGateway.TransitGatewayId != nil {
		ko.Status.TransitGatewayID = resp.TransitGateway.TransitGatewayId
	} else {
		ko.Status.TransitGatewayID = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateTransitGatewayInput, error) {
	res := &svcsdk.CreateTransitGatewayInput{}

	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.Options != nil {
		f1 := &svcsdk.TransitGatewayRequestOptions{}
		if r.ko.Spec.Options.AmazonSideASN != nil {
			f1.SetAmazonSideAsn(*r.ko.Spec.Options.AmazonSideASN)
		}
		if r.ko.Spec.Options.AutoAcceptSharedAttachments != nil {
			f1.SetAutoAcceptSharedAttachments(*r.ko.Spec.Options.AutoAcceptSharedAttachments)
		}
		if r.ko.Spec.Options.DefaultRouteTableAssociation != nil {
			f1.SetDefaultRouteTableAssociation(*r.ko.Spec.Options.DefaultRouteTableAssociation)
		}
		if r.ko.Spec.Options.DefaultRouteTablePropagation != nil {
			f1.SetDefaultRouteTablePropagation(*r.ko.Spec.Options.DefaultRouteTablePropagation)
		}
		if r.ko.Spec.Options.DNSSupport != nil {
			f1.SetDnsSupport(*r.ko.Spec.Options.DNSSupport)
		}
		if r.ko.Spec.Options.MulticastSupport != nil {
			f1.SetMulticastSupport(*r.ko.Spec.Options.MulticastSupport)
		}
		if r.ko.Spec.Options.TransitGatewayCIDRBlocks != nil {
			f1f6 := []*string{}
			for _, f1f6iter := range r.ko.Spec.Options.TransitGatewayCIDRBlocks {
				var f1f6elem string
				f1f6elem = *f1f6iter
				f1f6 = append(f1f6, &f1f6elem)
			}
			f1.SetTransitGatewayCidrBlocks(f1f6)
		}
		if r.ko.Spec.Options.VPNECMPSupport != nil {
			f1.SetVpnEcmpSupport(*r.ko.Spec.Options.VPNECMPSupport)
		}
		res.SetOptions(f1)
	}
	if r.ko.Spec.TagSpecifications != nil {
		f2 := []*svcsdk.TagSpecification{}
		for _, f2iter := range r.ko.Spec.TagSpecifications {
			f2elem := &svcsdk.TagSpecification{}
			if f2iter.ResourceType != nil {
				f2elem.SetResourceType(*f2iter.ResourceType)
			}
			if f2iter.Tags != nil {
				f2elemf1 := []*svcsdk.Tag{}
				for _, f2elemf1iter := range f2iter.Tags {
					f2elemf1elem := &svcsdk.Tag{}
					if f2elemf1iter.Key != nil {
						f2elemf1elem.SetKey(*f2elemf1iter.Key)
					}
					if f2elemf1iter.Value != nil {
						f2elemf1elem.SetValue(*f2elemf1iter.Value)
					}
					f2elemf1 = append(f2elemf1, f2elemf1elem)
				}
				f2elem.SetTags(f2elemf1)
			}
			f2 = append(f2, f2elem)
		}
		res.SetTagSpecifications(f2)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteTransitGatewayOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteTransitGatewayWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteTransitGateway", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteTransitGatewayInput, error) {
	res := &svcsdk.DeleteTransitGatewayInput{}

	if r.ko.Status.TransitGatewayID != nil {
		res.SetTransitGatewayId(*r.ko.Status.TransitGatewayID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.TransitGateway,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
