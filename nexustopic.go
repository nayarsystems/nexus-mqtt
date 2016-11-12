// Copyright (c) 2014 The SurgeMQ Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/nayarsystems/surgemq/topics"
	"github.com/surgemq/message"
)

var (
	// MaxQosAllowed is the maximum QOS supported by this server
	MaxQosAllowed = message.QosExactlyOnce
)

type nexusTopics struct {
}

func init() {
	topics.Register("nexus", NewNexusTopicsProvider())
}

func NewNexusTopicsProvider() *nexusTopics {
	return &nexusTopics{}
}

func (this *nexusTopics) Subscribe(topic []byte, qos byte, sub interface{}) (byte, error) {
	if !message.ValidQos(qos) {
		return message.QosFailure, fmt.Errorf("Invalid QoS %d", qos)
	}

	if sub == nil {
		return message.QosFailure, fmt.Errorf("Subscriber cannot be nil")
	}

	if qos > MaxQosAllowed {
		qos = MaxQosAllowed
	}
	return qos, nil
}

func (this *nexusTopics) Unsubscribe(topic []byte, sub interface{}) error {
	return nil
}

func (this *nexusTopics) Subscribers(topic []byte, qos byte, subs *[]interface{}, qoss *[]byte) error {
	if !message.ValidQos(qos) {
		return fmt.Errorf("Invalid QoS %d", qos)
	}
	*subs = []interface{}{}
	*qoss = []byte{}
	return nil
}

func (this *nexusTopics) Retain(msg *message.PublishMessage) error {
	return nil
}

func (this *nexusTopics) Retained(topic []byte, msgs *[]*message.PublishMessage) error {
	return nil
}

func (this *nexusTopics) Close() error {
	return nil
}
