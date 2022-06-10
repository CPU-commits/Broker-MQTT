package events

import (
	"fmt"

	"github.com/CPU-commits/broker_mqtt/logs"
	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
)

var logsStruct = &logs.Log{}

type Events struct {
	Server *mqtt.Server
}

func (e *Events) onConnect() {
	e.Server.Events.OnConnect = func(c events.Client, p events.Packet) {
		logsStruct.WriteLog(
			"Connection",
			"Info",
			fmt.Sprintf("Client connected %s. Username: %s", c.ID, string(p.Username)),
		)
		e.Server.Publish("agent/onconnect", []byte(string(p.Username)), false)
	}
}

func (e *Events) onDisconnect() {
	e.Server.Events.OnDisconnect = func(c events.Client, err error) {
		logsStruct.WriteLog(
			"Disconnect",
			"Info",
			fmt.Sprintf("Client disconnected %s", c.ID),
		)
		if err != nil {
			logsStruct.WriteLog(
				"Disconnect",
				"Error",
				fmt.Sprintf("Client disconnected %s. Error: %s", c.ID, err.Error()),
			)
		}
	}
}

func (e *Events) onSubscribe() {
	e.Server.Events.OnSubscribe = func(filter string, cl events.Client, qos byte) {
		logsStruct.WriteLog(
			"Subscribe",
			"Info",
			fmt.Sprintf(
				"Client %s. Subscribed %s. QoS: %s",
				cl.ID,
				filter,
				string(qos),
			),
		)
	}
}

func (e *Events) onUnsubscribe() {
	e.Server.Events.OnUnsubscribe = func(filter string, cl events.Client) {
		logsStruct.WriteLog(
			"Unsubscribe",
			"Info",
			fmt.Sprintf(
				"Client %s. Unsubscribed: %s",
				cl.ID,
				filter,
			),
		)
	}
}

func (e *Events) Init() {
	e.onConnect()
	e.onDisconnect()
	e.onSubscribe()
	e.onUnsubscribe()
}
