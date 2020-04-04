// Copyright 2018 Prometheus Team
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

package test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/prometheus/alertmanager/test/with_api_v2"
)

func TestSilencing(t *testing.T) {
	t.Parallel()

	conf := `
route:
  receiver: "default"
  group_by: []
  group_wait:      1s
  group_interval:  1s
  repeat_interval: 1ms

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
`

	at := NewAcceptanceTest(t, &AcceptanceOpts{
		Tolerance: 150 * time.Millisecond,
	})

	co := at.Collector("webhook")
	wh := NewWebhook(co)

	amc := at.AlertmanagerCluster(fmt.Sprintf(conf, wh.Address()), 1)

	// No repeat interval is configured. Thus, we receive an alert
	// notification every second.
	amc.Push(At(1), Alert("alertname", "test1").Active(1))
	amc.Push(At(1), Alert("alertname", "test2").Active(1))

	co.Want(Between(2, 2.5),
		Alert("alertname", "test1").Active(1),
		Alert("alertname", "test2").Active(1),
	)

	// Add a silence that affects the first alert.
	amc.SetSilence(At(2.3), Silence(2.5, 4.5).Match("alertname", "test1"))

	co.Want(Between(3, 3.5), Alert("alertname", "test2").Active(1))
	co.Want(Between(4, 4.5), Alert("alertname", "test2").Active(1))

	// Silence should be over now and we receive both alerts again.

	co.Want(Between(5, 5.5),
		Alert("alertname", "test1").Active(1),
		Alert("alertname", "test2").Active(1),
	)

	at.Run()

	t.Log(co.Check())
}

func TestSilenceDelete(t *testing.T) {
	t.Parallel()

	conf := `
route:
  receiver: "default"
  group_by: []
  group_wait:      1s
  group_interval:  1s
  repeat_interval: 1ms

receivers:
- name: "default"
  webhook_configs:
  - url: 'http://%s'
`

	at := NewAcceptanceTest(t, &AcceptanceOpts{
		Tolerance: 150 * time.Millisecond,
	})

	co := at.Collector("webhook")
	wh := NewWebhook(co)

	amc := at.AlertmanagerCluster(fmt.Sprintf(conf, wh.Address()), 1)

	// No repeat interval is configured. Thus, we receive an alert
	// notification every second.
	amc.Push(At(1), Alert("alertname", "test1").Active(1))
	amc.Push(At(1), Alert("alertname", "test2").Active(1))

	// Silence everything for a long time and delete the silence after
	// two iterations.
	sil := Silence(1.5, 100).MatchRE("alertname", ".*")

	amc.SetSilence(At(1.3), sil)
	amc.DelSilence(At(3.5), sil)

	co.Want(Between(3.5, 4.5),
		Alert("alertname", "test1").Active(1),
		Alert("alertname", "test2").Active(1),
	)

	at.Run()

	t.Log(co.Check())
}
