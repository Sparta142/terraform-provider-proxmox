/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatasourceNode(t *testing.T) {
	t.Parallel()

	te := initTestEnvironment(t)

	tests := []struct {
		name  string
		steps []resource.TestStep
	}{
		{"read node attributes", []resource.TestStep{{
			Config: te.renderConfig(`data "proxmox_virtual_environment_node" "test" { node_name = "{{.NodeName}}" }`),
			Check: resource.ComposeTestCheckFunc(
				testResourceAttributesSet("data.proxmox_virtual_environment_node.test", []string{
					"cpu_count",
					"cpu_sockets",
					"cpu_model",
					"memory_available",
					"memory_used",
					"memory_total",
					"uptime",
				}),
			),
		}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource.ParallelTest(t, resource.TestCase{
				ProtoV6ProviderFactories: te.accProviders,
				Steps:                    tt.steps,
			})
		})
	}
}
