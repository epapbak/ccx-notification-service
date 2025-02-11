/*
Copyright © 2022 Red Hat, Inc.

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

// Package disabled contains an empty implementation of Producer interface.
package disabled

// Generated documentation is available at:
// https://pkg.go.dev/github.com/RedHatInsights/ccx-notification-service/producer
//
// Documentation in literate-programming-style is available at:
// https://redhatinsights.github.io/ccx-notification-service/packages/disabled/disabled_producer.html

import "github.com/RedHatInsights/ccx-notification-service/types"

// Producer is an implementation of Producer interface where no message is sent
type Producer struct {
}

// ProduceMessage doesn't publish any message.
func (producer *Producer) ProduceMessage(msg types.ProducerMessage) (partitionID int32, offset int64, err error) {
	return 0, -1, nil
}

// Close return nil
func (producer *Producer) Close() error {
	return nil
}
