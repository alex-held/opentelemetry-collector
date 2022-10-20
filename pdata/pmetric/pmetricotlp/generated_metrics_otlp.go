// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package pmetricotlp

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpcollectormetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/metrics/v1"
)

// ExportMetricsPartialSuccess represents the details of a partially successful export request.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewExportMetricsPartialSuccess function to create new instances.
// Important: zero-initialized instance is not valid for use.

type ExportMetricsPartialSuccess internal.ExportMetricsPartialSuccess

func newExportMetricsPartialSuccess(orig *otlpcollectormetrics.ExportMetricsPartialSuccess) ExportMetricsPartialSuccess {
	return ExportMetricsPartialSuccess(internal.NewExportMetricsPartialSuccess(orig))
}

func (ms ExportMetricsPartialSuccess) getOrig() *otlpcollectormetrics.ExportMetricsPartialSuccess {
	return internal.GetOrigExportMetricsPartialSuccess(internal.ExportMetricsPartialSuccess(ms))
}

// NewExportMetricsPartialSuccess creates a new empty ExportMetricsPartialSuccess.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewExportMetricsPartialSuccess() ExportMetricsPartialSuccess {
	return newExportMetricsPartialSuccess(&otlpcollectormetrics.ExportMetricsPartialSuccess{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ExportMetricsPartialSuccess) MoveTo(dest ExportMetricsPartialSuccess) {
	*dest.getOrig() = *ms.getOrig()
	*ms.getOrig() = otlpcollectormetrics.ExportMetricsPartialSuccess{}
}

// RejectedDataPoints returns the rejecteddatapoints associated with this ExportMetricsPartialSuccess.
func (ms ExportMetricsPartialSuccess) RejectedDataPoints() int64 {
	return ms.getOrig().RejectedDataPoints
}

// SetRejectedDataPoints replaces the rejecteddatapoints associated with this ExportMetricsPartialSuccess.
func (ms ExportMetricsPartialSuccess) SetRejectedDataPoints(v int64) {
	ms.getOrig().RejectedDataPoints = v
}

// ErrorMessage returns the errormessage associated with this ExportMetricsPartialSuccess.
func (ms ExportMetricsPartialSuccess) ErrorMessage() string {
	return ms.getOrig().ErrorMessage
}

// SetErrorMessage replaces the errormessage associated with this ExportMetricsPartialSuccess.
func (ms ExportMetricsPartialSuccess) SetErrorMessage(v string) {
	ms.getOrig().ErrorMessage = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ExportMetricsPartialSuccess) CopyTo(dest ExportMetricsPartialSuccess) {
	dest.SetRejectedDataPoints(ms.RejectedDataPoints())
	dest.SetErrorMessage(ms.ErrorMessage())
}
