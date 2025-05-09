// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ControllerCertificateApplyConfiguration represents a declarative configuration of the ControllerCertificate type for use
// with apply.
type ControllerCertificateApplyConfiguration struct {
	Subject    *string      `json:"subject,omitempty"`
	Signer     *string      `json:"signer,omitempty"`
	NotBefore  *metav1.Time `json:"notBefore,omitempty"`
	NotAfter   *metav1.Time `json:"notAfter,omitempty"`
	BundleFile *string      `json:"bundleFile,omitempty"`
}

// ControllerCertificateApplyConfiguration constructs a declarative configuration of the ControllerCertificate type for use with
// apply.
func ControllerCertificate() *ControllerCertificateApplyConfiguration {
	return &ControllerCertificateApplyConfiguration{}
}

// WithSubject sets the Subject field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subject field is set to the value of the last call.
func (b *ControllerCertificateApplyConfiguration) WithSubject(value string) *ControllerCertificateApplyConfiguration {
	b.Subject = &value
	return b
}

// WithSigner sets the Signer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Signer field is set to the value of the last call.
func (b *ControllerCertificateApplyConfiguration) WithSigner(value string) *ControllerCertificateApplyConfiguration {
	b.Signer = &value
	return b
}

// WithNotBefore sets the NotBefore field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NotBefore field is set to the value of the last call.
func (b *ControllerCertificateApplyConfiguration) WithNotBefore(value metav1.Time) *ControllerCertificateApplyConfiguration {
	b.NotBefore = &value
	return b
}

// WithNotAfter sets the NotAfter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NotAfter field is set to the value of the last call.
func (b *ControllerCertificateApplyConfiguration) WithNotAfter(value metav1.Time) *ControllerCertificateApplyConfiguration {
	b.NotAfter = &value
	return b
}

// WithBundleFile sets the BundleFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BundleFile field is set to the value of the last call.
func (b *ControllerCertificateApplyConfiguration) WithBundleFile(value string) *ControllerCertificateApplyConfiguration {
	b.BundleFile = &value
	return b
}
