// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	hive "github.com/operator-framework/operator-metering/pkg/hive"
	presto "github.com/operator-framework/operator-metering/pkg/presto"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSBillingDataSource) DeepCopyInto(out *AWSBillingDataSource) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		if *in == nil {
			*out = nil
		} else {
			*out = new(S3Bucket)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSBillingDataSource.
func (in *AWSBillingDataSource) DeepCopy() *AWSBillingDataSource {
	if in == nil {
		return nil
	}
	out := new(AWSBillingDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenQueryView) DeepCopyInto(out *GenQueryView) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenQueryView.
func (in *GenQueryView) DeepCopy() *GenQueryView {
	if in == nil {
		return nil
	}
	out := new(GenQueryView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HiveStorage) DeepCopyInto(out *HiveStorage) {
	*out = *in
	in.TableProperties.DeepCopyInto(&out.TableProperties)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HiveStorage.
func (in *HiveStorage) DeepCopy() *HiveStorage {
	if in == nil {
		return nil
	}
	out := new(HiveStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoTable) DeepCopyInto(out *PrestoTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoTable.
func (in *PrestoTable) DeepCopy() *PrestoTable {
	if in == nil {
		return nil
	}
	out := new(PrestoTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrestoTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoTableList) DeepCopyInto(out *PrestoTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*PrestoTable, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(PrestoTable)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoTableList.
func (in *PrestoTableList) DeepCopy() *PrestoTableList {
	if in == nil {
		return nil
	}
	out := new(PrestoTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrestoTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrestoTableStatus) DeepCopyInto(out *PrestoTableStatus) {
	*out = *in
	in.Parameters.DeepCopyInto(&out.Parameters)
	in.Properties.DeepCopyInto(&out.Properties)
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = make([]TablePartition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrestoTableStatus.
func (in *PrestoTableStatus) DeepCopy() *PrestoTableStatus {
	if in == nil {
		return nil
	}
	out := new(PrestoTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMetricsDataSource) DeepCopyInto(out *PrometheusMetricsDataSource) {
	*out = *in
	if in.QueryConfig != nil {
		in, out := &in.QueryConfig, &out.QueryConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(PrometheusQueryConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		if *in == nil {
			*out = nil
		} else {
			*out = new(StorageLocationRef)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMetricsDataSource.
func (in *PrometheusMetricsDataSource) DeepCopy() *PrometheusMetricsDataSource {
	if in == nil {
		return nil
	}
	out := new(PrometheusMetricsDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusQueryConfig) DeepCopyInto(out *PrometheusQueryConfig) {
	*out = *in
	if in.QueryInterval != nil {
		in, out := &in.QueryInterval, &out.QueryInterval
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	if in.StepSize != nil {
		in, out := &in.StepSize, &out.StepSize
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	if in.ChunkSize != nil {
		in, out := &in.ChunkSize, &out.ChunkSize
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusQueryConfig.
func (in *PrometheusQueryConfig) DeepCopy() *PrometheusQueryConfig {
	if in == nil {
		return nil
	}
	out := new(PrometheusQueryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Report) DeepCopyInto(out *Report) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Report.
func (in *Report) DeepCopy() *Report {
	if in == nil {
		return nil
	}
	out := new(Report)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Report) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportDataSource) DeepCopyInto(out *ReportDataSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportDataSource.
func (in *ReportDataSource) DeepCopy() *ReportDataSource {
	if in == nil {
		return nil
	}
	out := new(ReportDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportDataSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportDataSourceList) DeepCopyInto(out *ReportDataSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*ReportDataSource, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(ReportDataSource)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportDataSourceList.
func (in *ReportDataSourceList) DeepCopy() *ReportDataSourceList {
	if in == nil {
		return nil
	}
	out := new(ReportDataSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportDataSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportDataSourceSpec) DeepCopyInto(out *ReportDataSourceSpec) {
	*out = *in
	if in.Promsum != nil {
		in, out := &in.Promsum, &out.Promsum
		if *in == nil {
			*out = nil
		} else {
			*out = new(PrometheusMetricsDataSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AWSBilling != nil {
		in, out := &in.AWSBilling, &out.AWSBilling
		if *in == nil {
			*out = nil
		} else {
			*out = new(AWSBillingDataSource)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportDataSourceSpec.
func (in *ReportDataSourceSpec) DeepCopy() *ReportDataSourceSpec {
	if in == nil {
		return nil
	}
	out := new(ReportDataSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportDataSourceStatus) DeepCopyInto(out *ReportDataSourceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportDataSourceStatus.
func (in *ReportDataSourceStatus) DeepCopy() *ReportDataSourceStatus {
	if in == nil {
		return nil
	}
	out := new(ReportDataSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportGenerationQuery) DeepCopyInto(out *ReportGenerationQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportGenerationQuery.
func (in *ReportGenerationQuery) DeepCopy() *ReportGenerationQuery {
	if in == nil {
		return nil
	}
	out := new(ReportGenerationQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportGenerationQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportGenerationQueryColumn) DeepCopyInto(out *ReportGenerationQueryColumn) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportGenerationQueryColumn.
func (in *ReportGenerationQueryColumn) DeepCopy() *ReportGenerationQueryColumn {
	if in == nil {
		return nil
	}
	out := new(ReportGenerationQueryColumn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportGenerationQueryInputDefinition) DeepCopyInto(out *ReportGenerationQueryInputDefinition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportGenerationQueryInputDefinition.
func (in *ReportGenerationQueryInputDefinition) DeepCopy() *ReportGenerationQueryInputDefinition {
	if in == nil {
		return nil
	}
	out := new(ReportGenerationQueryInputDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportGenerationQueryInputValue) DeepCopyInto(out *ReportGenerationQueryInputValue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportGenerationQueryInputValue.
func (in *ReportGenerationQueryInputValue) DeepCopy() *ReportGenerationQueryInputValue {
	if in == nil {
		return nil
	}
	out := new(ReportGenerationQueryInputValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ReportGenerationQueryInputValues) DeepCopyInto(out *ReportGenerationQueryInputValues) {
	{
		in := &in
		*out = make(ReportGenerationQueryInputValues, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportGenerationQueryInputValues.
func (in ReportGenerationQueryInputValues) DeepCopy() ReportGenerationQueryInputValues {
	if in == nil {
		return nil
	}
	out := new(ReportGenerationQueryInputValues)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportGenerationQueryList) DeepCopyInto(out *ReportGenerationQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*ReportGenerationQuery, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(ReportGenerationQuery)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportGenerationQueryList.
func (in *ReportGenerationQueryList) DeepCopy() *ReportGenerationQueryList {
	if in == nil {
		return nil
	}
	out := new(ReportGenerationQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportGenerationQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportGenerationQuerySpec) DeepCopyInto(out *ReportGenerationQuerySpec) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]ReportGenerationQueryColumn, len(*in))
		copy(*out, *in)
	}
	out.View = in.View
	if in.ReportQueries != nil {
		in, out := &in.ReportQueries, &out.ReportQueries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DynamicReportQueries != nil {
		in, out := &in.DynamicReportQueries, &out.DynamicReportQueries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DataSources != nil {
		in, out := &in.DataSources, &out.DataSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Reports != nil {
		in, out := &in.Reports, &out.Reports
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ScheduledReports != nil {
		in, out := &in.ScheduledReports, &out.ScheduledReports
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]ReportGenerationQueryInputDefinition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportGenerationQuerySpec.
func (in *ReportGenerationQuerySpec) DeepCopy() *ReportGenerationQuerySpec {
	if in == nil {
		return nil
	}
	out := new(ReportGenerationQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportGenerationQueryStatus) DeepCopyInto(out *ReportGenerationQueryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportGenerationQueryStatus.
func (in *ReportGenerationQueryStatus) DeepCopy() *ReportGenerationQueryStatus {
	if in == nil {
		return nil
	}
	out := new(ReportGenerationQueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportList) DeepCopyInto(out *ReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*Report, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Report)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportList.
func (in *ReportList) DeepCopy() *ReportList {
	if in == nil {
		return nil
	}
	out := new(ReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportPrometheusQuery) DeepCopyInto(out *ReportPrometheusQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportPrometheusQuery.
func (in *ReportPrometheusQuery) DeepCopy() *ReportPrometheusQuery {
	if in == nil {
		return nil
	}
	out := new(ReportPrometheusQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportPrometheusQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportPrometheusQueryList) DeepCopyInto(out *ReportPrometheusQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*ReportPrometheusQuery, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(ReportPrometheusQuery)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportPrometheusQueryList.
func (in *ReportPrometheusQueryList) DeepCopy() *ReportPrometheusQueryList {
	if in == nil {
		return nil
	}
	out := new(ReportPrometheusQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportPrometheusQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportPrometheusQuerySpec) DeepCopyInto(out *ReportPrometheusQuerySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportPrometheusQuerySpec.
func (in *ReportPrometheusQuerySpec) DeepCopy() *ReportPrometheusQuerySpec {
	if in == nil {
		return nil
	}
	out := new(ReportPrometheusQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportSpec) DeepCopyInto(out *ReportSpec) {
	*out = *in
	if in.ReportingStart != nil {
		in, out := &in.ReportingStart, &out.ReportingStart
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.ReportingEnd != nil {
		in, out := &in.ReportingEnd, &out.ReportingEnd
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make(ReportGenerationQueryInputValues, len(*in))
		copy(*out, *in)
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		if *in == nil {
			*out = nil
		} else {
			*out = new(StorageLocationRef)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportSpec.
func (in *ReportSpec) DeepCopy() *ReportSpec {
	if in == nil {
		return nil
	}
	out := new(ReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportStatus) DeepCopyInto(out *ReportStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportStatus.
func (in *ReportStatus) DeepCopy() *ReportStatus {
	if in == nil {
		return nil
	}
	out := new(ReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bucket) DeepCopyInto(out *S3Bucket) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bucket.
func (in *S3Bucket) DeepCopy() *S3Bucket {
	if in == nil {
		return nil
	}
	out := new(S3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReport) DeepCopyInto(out *ScheduledReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReport.
func (in *ScheduledReport) DeepCopy() *ScheduledReport {
	if in == nil {
		return nil
	}
	out := new(ScheduledReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportCondition) DeepCopyInto(out *ScheduledReportCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportCondition.
func (in *ScheduledReportCondition) DeepCopy() *ScheduledReportCondition {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportList) DeepCopyInto(out *ScheduledReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*ScheduledReport, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(ScheduledReport)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportList.
func (in *ScheduledReportList) DeepCopy() *ScheduledReportList {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportSchedule) DeepCopyInto(out *ScheduledReportSchedule) {
	*out = *in
	if in.Cron != nil {
		in, out := &in.Cron, &out.Cron
		if *in == nil {
			*out = nil
		} else {
			*out = new(ScheduledReportScheduleCron)
			**out = **in
		}
	}
	if in.Hourly != nil {
		in, out := &in.Hourly, &out.Hourly
		if *in == nil {
			*out = nil
		} else {
			*out = new(ScheduledReportScheduleHourly)
			**out = **in
		}
	}
	if in.Daily != nil {
		in, out := &in.Daily, &out.Daily
		if *in == nil {
			*out = nil
		} else {
			*out = new(ScheduledReportScheduleDaily)
			**out = **in
		}
	}
	if in.Weekly != nil {
		in, out := &in.Weekly, &out.Weekly
		if *in == nil {
			*out = nil
		} else {
			*out = new(ScheduledReportScheduleWeekly)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Monthly != nil {
		in, out := &in.Monthly, &out.Monthly
		if *in == nil {
			*out = nil
		} else {
			*out = new(ScheduledReportScheduleMonthly)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportSchedule.
func (in *ScheduledReportSchedule) DeepCopy() *ScheduledReportSchedule {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportScheduleCron) DeepCopyInto(out *ScheduledReportScheduleCron) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportScheduleCron.
func (in *ScheduledReportScheduleCron) DeepCopy() *ScheduledReportScheduleCron {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportScheduleCron)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportScheduleDaily) DeepCopyInto(out *ScheduledReportScheduleDaily) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportScheduleDaily.
func (in *ScheduledReportScheduleDaily) DeepCopy() *ScheduledReportScheduleDaily {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportScheduleDaily)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportScheduleHourly) DeepCopyInto(out *ScheduledReportScheduleHourly) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportScheduleHourly.
func (in *ScheduledReportScheduleHourly) DeepCopy() *ScheduledReportScheduleHourly {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportScheduleHourly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportScheduleMonthly) DeepCopyInto(out *ScheduledReportScheduleMonthly) {
	*out = *in
	if in.DayOfMonth != nil {
		in, out := &in.DayOfMonth, &out.DayOfMonth
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportScheduleMonthly.
func (in *ScheduledReportScheduleMonthly) DeepCopy() *ScheduledReportScheduleMonthly {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportScheduleMonthly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportScheduleWeekly) DeepCopyInto(out *ScheduledReportScheduleWeekly) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportScheduleWeekly.
func (in *ScheduledReportScheduleWeekly) DeepCopy() *ScheduledReportScheduleWeekly {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportScheduleWeekly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportSpec) DeepCopyInto(out *ScheduledReportSpec) {
	*out = *in
	in.Schedule.DeepCopyInto(&out.Schedule)
	if in.ReportingStart != nil {
		in, out := &in.ReportingStart, &out.ReportingStart
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.ReportingEnd != nil {
		in, out := &in.ReportingEnd, &out.ReportingEnd
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make(ReportGenerationQueryInputValues, len(*in))
		copy(*out, *in)
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		if *in == nil {
			*out = nil
		} else {
			*out = new(StorageLocationRef)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportSpec.
func (in *ScheduledReportSpec) DeepCopy() *ScheduledReportSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledReportStatus) DeepCopyInto(out *ScheduledReportStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ScheduledReportCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastReportTime != nil {
		in, out := &in.LastReportTime, &out.LastReportTime
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledReportStatus.
func (in *ScheduledReportStatus) DeepCopy() *ScheduledReportStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduledReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageLocation) DeepCopyInto(out *StorageLocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageLocation.
func (in *StorageLocation) DeepCopy() *StorageLocation {
	if in == nil {
		return nil
	}
	out := new(StorageLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageLocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageLocationList) DeepCopyInto(out *StorageLocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*StorageLocation, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(StorageLocation)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageLocationList.
func (in *StorageLocationList) DeepCopy() *StorageLocationList {
	if in == nil {
		return nil
	}
	out := new(StorageLocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageLocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageLocationRef) DeepCopyInto(out *StorageLocationRef) {
	*out = *in
	if in.StorageSpec != nil {
		in, out := &in.StorageSpec, &out.StorageSpec
		if *in == nil {
			*out = nil
		} else {
			*out = new(StorageLocationSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageLocationRef.
func (in *StorageLocationRef) DeepCopy() *StorageLocationRef {
	if in == nil {
		return nil
	}
	out := new(StorageLocationRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageLocationSpec) DeepCopyInto(out *StorageLocationSpec) {
	*out = *in
	if in.Hive != nil {
		in, out := &in.Hive, &out.Hive
		if *in == nil {
			*out = nil
		} else {
			*out = new(HiveStorage)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageLocationSpec.
func (in *StorageLocationSpec) DeepCopy() *StorageLocationSpec {
	if in == nil {
		return nil
	}
	out := new(StorageLocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableParameters) DeepCopyInto(out *TableParameters) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]hive.Column, len(*in))
		copy(*out, *in)
	}
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = make([]hive.Column, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableParameters.
func (in *TableParameters) DeepCopy() *TableParameters {
	if in == nil {
		return nil
	}
	out := new(TableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TablePartition) DeepCopyInto(out *TablePartition) {
	*out = *in
	if in.PartitionSpec != nil {
		in, out := &in.PartitionSpec, &out.PartitionSpec
		*out = make(presto.PartitionSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TablePartition.
func (in *TablePartition) DeepCopy() *TablePartition {
	if in == nil {
		return nil
	}
	out := new(TablePartition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableProperties) DeepCopyInto(out *TableProperties) {
	*out = *in
	if in.SerdeRowProperties != nil {
		in, out := &in.SerdeRowProperties, &out.SerdeRowProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableProperties.
func (in *TableProperties) DeepCopy() *TableProperties {
	if in == nil {
		return nil
	}
	out := new(TableProperties)
	in.DeepCopyInto(out)
	return out
}
