//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pack) DeepCopyInto(out *Pack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pack.
func (in *Pack) DeepCopy() *Pack {
	if in == nil {
		return nil
	}
	out := new(Pack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackInitParameters) DeepCopyInto(out *PackInitParameters) {
	*out = *in
	if in.CertificateAuthority != nil {
		in, out := &in.CertificateAuthority, &out.CertificateAuthority
		*out = new(string)
		**out = **in
	}
	if in.CloudflareBranding != nil {
		in, out := &in.CloudflareBranding, &out.CloudflareBranding
		*out = new(bool)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidationErrors != nil {
		in, out := &in.ValidationErrors, &out.ValidationErrors
		*out = make([]ValidationErrorsInitParameters, len(*in))
		copy(*out, *in)
	}
	if in.ValidationMethod != nil {
		in, out := &in.ValidationMethod, &out.ValidationMethod
		*out = new(string)
		**out = **in
	}
	if in.ValidationRecords != nil {
		in, out := &in.ValidationRecords, &out.ValidationRecords
		*out = make([]ValidationRecordsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ValidityDays != nil {
		in, out := &in.ValidityDays, &out.ValidityDays
		*out = new(float64)
		**out = **in
	}
	if in.WaitForActiveStatus != nil {
		in, out := &in.WaitForActiveStatus, &out.WaitForActiveStatus
		*out = new(bool)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDRef != nil {
		in, out := &in.ZoneIDRef, &out.ZoneIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneIDSelector != nil {
		in, out := &in.ZoneIDSelector, &out.ZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackInitParameters.
func (in *PackInitParameters) DeepCopy() *PackInitParameters {
	if in == nil {
		return nil
	}
	out := new(PackInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackList) DeepCopyInto(out *PackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackList.
func (in *PackList) DeepCopy() *PackList {
	if in == nil {
		return nil
	}
	out := new(PackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackObservation) DeepCopyInto(out *PackObservation) {
	*out = *in
	if in.CertificateAuthority != nil {
		in, out := &in.CertificateAuthority, &out.CertificateAuthority
		*out = new(string)
		**out = **in
	}
	if in.CloudflareBranding != nil {
		in, out := &in.CloudflareBranding, &out.CloudflareBranding
		*out = new(bool)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidationErrors != nil {
		in, out := &in.ValidationErrors, &out.ValidationErrors
		*out = make([]ValidationErrorsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ValidationMethod != nil {
		in, out := &in.ValidationMethod, &out.ValidationMethod
		*out = new(string)
		**out = **in
	}
	if in.ValidationRecords != nil {
		in, out := &in.ValidationRecords, &out.ValidationRecords
		*out = make([]ValidationRecordsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ValidityDays != nil {
		in, out := &in.ValidityDays, &out.ValidityDays
		*out = new(float64)
		**out = **in
	}
	if in.WaitForActiveStatus != nil {
		in, out := &in.WaitForActiveStatus, &out.WaitForActiveStatus
		*out = new(bool)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackObservation.
func (in *PackObservation) DeepCopy() *PackObservation {
	if in == nil {
		return nil
	}
	out := new(PackObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackParameters) DeepCopyInto(out *PackParameters) {
	*out = *in
	if in.CertificateAuthority != nil {
		in, out := &in.CertificateAuthority, &out.CertificateAuthority
		*out = new(string)
		**out = **in
	}
	if in.CloudflareBranding != nil {
		in, out := &in.CloudflareBranding, &out.CloudflareBranding
		*out = new(bool)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidationErrors != nil {
		in, out := &in.ValidationErrors, &out.ValidationErrors
		*out = make([]ValidationErrorsParameters, len(*in))
		copy(*out, *in)
	}
	if in.ValidationMethod != nil {
		in, out := &in.ValidationMethod, &out.ValidationMethod
		*out = new(string)
		**out = **in
	}
	if in.ValidationRecords != nil {
		in, out := &in.ValidationRecords, &out.ValidationRecords
		*out = make([]ValidationRecordsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ValidityDays != nil {
		in, out := &in.ValidityDays, &out.ValidityDays
		*out = new(float64)
		**out = **in
	}
	if in.WaitForActiveStatus != nil {
		in, out := &in.WaitForActiveStatus, &out.WaitForActiveStatus
		*out = new(bool)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDRef != nil {
		in, out := &in.ZoneIDRef, &out.ZoneIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneIDSelector != nil {
		in, out := &in.ZoneIDSelector, &out.ZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackParameters.
func (in *PackParameters) DeepCopy() *PackParameters {
	if in == nil {
		return nil
	}
	out := new(PackParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackSpec) DeepCopyInto(out *PackSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackSpec.
func (in *PackSpec) DeepCopy() *PackSpec {
	if in == nil {
		return nil
	}
	out := new(PackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackStatus) DeepCopyInto(out *PackStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackStatus.
func (in *PackStatus) DeepCopy() *PackStatus {
	if in == nil {
		return nil
	}
	out := new(PackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationErrorsInitParameters) DeepCopyInto(out *ValidationErrorsInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationErrorsInitParameters.
func (in *ValidationErrorsInitParameters) DeepCopy() *ValidationErrorsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ValidationErrorsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationErrorsObservation) DeepCopyInto(out *ValidationErrorsObservation) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationErrorsObservation.
func (in *ValidationErrorsObservation) DeepCopy() *ValidationErrorsObservation {
	if in == nil {
		return nil
	}
	out := new(ValidationErrorsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationErrorsParameters) DeepCopyInto(out *ValidationErrorsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationErrorsParameters.
func (in *ValidationErrorsParameters) DeepCopy() *ValidationErrorsParameters {
	if in == nil {
		return nil
	}
	out := new(ValidationErrorsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationRecordsInitParameters) DeepCopyInto(out *ValidationRecordsInitParameters) {
	*out = *in
	if in.CnameName != nil {
		in, out := &in.CnameName, &out.CnameName
		*out = new(string)
		**out = **in
	}
	if in.CnameTarget != nil {
		in, out := &in.CnameTarget, &out.CnameTarget
		*out = new(string)
		**out = **in
	}
	if in.Emails != nil {
		in, out := &in.Emails, &out.Emails
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.HTTPBody != nil {
		in, out := &in.HTTPBody, &out.HTTPBody
		*out = new(string)
		**out = **in
	}
	if in.HTTPURL != nil {
		in, out := &in.HTTPURL, &out.HTTPURL
		*out = new(string)
		**out = **in
	}
	if in.TxtName != nil {
		in, out := &in.TxtName, &out.TxtName
		*out = new(string)
		**out = **in
	}
	if in.TxtValue != nil {
		in, out := &in.TxtValue, &out.TxtValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationRecordsInitParameters.
func (in *ValidationRecordsInitParameters) DeepCopy() *ValidationRecordsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ValidationRecordsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationRecordsObservation) DeepCopyInto(out *ValidationRecordsObservation) {
	*out = *in
	if in.CnameName != nil {
		in, out := &in.CnameName, &out.CnameName
		*out = new(string)
		**out = **in
	}
	if in.CnameTarget != nil {
		in, out := &in.CnameTarget, &out.CnameTarget
		*out = new(string)
		**out = **in
	}
	if in.Emails != nil {
		in, out := &in.Emails, &out.Emails
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.HTTPBody != nil {
		in, out := &in.HTTPBody, &out.HTTPBody
		*out = new(string)
		**out = **in
	}
	if in.HTTPURL != nil {
		in, out := &in.HTTPURL, &out.HTTPURL
		*out = new(string)
		**out = **in
	}
	if in.TxtName != nil {
		in, out := &in.TxtName, &out.TxtName
		*out = new(string)
		**out = **in
	}
	if in.TxtValue != nil {
		in, out := &in.TxtValue, &out.TxtValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationRecordsObservation.
func (in *ValidationRecordsObservation) DeepCopy() *ValidationRecordsObservation {
	if in == nil {
		return nil
	}
	out := new(ValidationRecordsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationRecordsParameters) DeepCopyInto(out *ValidationRecordsParameters) {
	*out = *in
	if in.CnameName != nil {
		in, out := &in.CnameName, &out.CnameName
		*out = new(string)
		**out = **in
	}
	if in.CnameTarget != nil {
		in, out := &in.CnameTarget, &out.CnameTarget
		*out = new(string)
		**out = **in
	}
	if in.Emails != nil {
		in, out := &in.Emails, &out.Emails
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.HTTPBody != nil {
		in, out := &in.HTTPBody, &out.HTTPBody
		*out = new(string)
		**out = **in
	}
	if in.HTTPURL != nil {
		in, out := &in.HTTPURL, &out.HTTPURL
		*out = new(string)
		**out = **in
	}
	if in.TxtName != nil {
		in, out := &in.TxtName, &out.TxtName
		*out = new(string)
		**out = **in
	}
	if in.TxtValue != nil {
		in, out := &in.TxtValue, &out.TxtValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationRecordsParameters.
func (in *ValidationRecordsParameters) DeepCopy() *ValidationRecordsParameters {
	if in == nil {
		return nil
	}
	out := new(ValidationRecordsParameters)
	in.DeepCopyInto(out)
	return out
}
