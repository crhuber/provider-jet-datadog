/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dashboard

import (
	tjconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the monitor group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("datadog_dashboard", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "dashboard"
	})

	p.AddResourceConfigurator("datadog_dashboard_json", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.Kind = "DashboardJSON"
		r.ShortGroup = "dashboard"
	})
}
