// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
//
// SPDX-License-Identifier: Apache-2.0
package web2wrapper

import (
	"context"
	"github.com/nats-io/nats.go"
	"log"
)

type natsBroker struct {
	connection    *nats.Conn
	encConnection *nats.EncodedConn
	clientName    string
	topic         string
	queue         string
}

func NewNatsBrokerFor(config *Config) (Broker, error) {
	brokerCfg := config.Broker
	return NewNatsBroker(brokerCfg.Url, brokerCfg.Topic, brokerCfg.Queue, brokerCfg.Clientname)
}

func NewNatsBroker(natsHost string, topic string, queue string, clientName string, natsOptions ...nats.Option) (Broker, error) {
	errorHandler := nats.ErrorHandler(func(nc *nats.Conn, s *nats.Subscription, err error) {
		if s != nil {
			log.Printf("Async error in %q/%q: %v", s.Subject, s.Queue, err)
		} else {
			log.Printf("Async error outside subscription: %v", err)
		}
	})
	natsOptions = append(natsOptions, errorHandler)
	natsOptions = append(natsOptions, nats.Name(clientName))
	nc, err := nats.Connect(natsHost, natsOptions...)
	if err != nil {
		return nil, err
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	return &natsBroker{
		connection:    nc,
		encConnection: ec,
		topic:         topic,
		queue:         queue,
		clientName:    clientName,
	}, nil
}

func (b *natsBroker) Close() error {
	b.connection.Close()
	b.encConnection.Close()
	return nil
}

func (b *natsBroker) Submit(msg *Msg) error {
	return b.encConnection.Publish(b.topic, msg)
}

func (b *natsBroker) Receive(ctx context.Context) <-chan *Msg {
	result := make(chan *Msg)
	subscription, err := b.encConnection.QueueSubscribe(b.topic, b.queue, func(msg *Msg) {
		result <- msg
	})
	if err != nil {
		log.Print(err)
		close(result)
		return result
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("Closing subscrition %s - %s on %s", b.topic, b.queue, b.clientName)
				if err := subscription.Unsubscribe(); err != nil {
					log.Print(err)
				}
				close(result)
				return
			}
		}
	}()
	return result
}
