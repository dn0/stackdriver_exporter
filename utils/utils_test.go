// Copyright 2020 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/prometheus-community/stackdriver_exporter/utils"
)

var _ = Describe("NormalizeMetricName", func() {
	It("returns a normalized metric name", func() {
		Expect(NormalizeMetricName("This_is__a-MetricName.Example/with:0totals")).To(Equal("this_is_a_metric_name_example_with_0_totals"))
	})
})

var _ = Describe("ProjectResource", func() {
	It("returns a project resource", func() {
		Expect(ProjectResource("fake-project-1")).To(Equal("projects/fake-project-1"))
	})
})

var _ = Describe("SplitExtraFilter", func() {
	It("returns an empty string from incomplete filter", func() {
		metricPrefix, filterQuery := SplitExtraFilter("This_is__a-MetricName.Example/with/no/filter", ":")
		Expect(metricPrefix).To(Equal(""))
		Expect(filterQuery).To(Equal(""))
	})

	It("returns a metric prefix and filter query from basic filter", func() {
		metricPrefix, filterQuery := SplitExtraFilter("This_is__a-MetricName.Example/with:filter.name=filter_value", ":")
		Expect(metricPrefix).To(Equal("This_is__a-MetricName.Example/with"))
		Expect(filterQuery).To(Equal("filter.name=filter_value"))
	})

	It("returns a metric prefix and filter query filter with colon", func() {
		metricPrefix, filterQuery := SplitExtraFilter(`This_is__a-MetricName.Example/with:filter.name="filter:value"`, ":")
		Expect(metricPrefix).To(Equal("This_is__a-MetricName.Example/with"))
		Expect(filterQuery).To(Equal("filter.name=\"filter:value\""))
	})
})
