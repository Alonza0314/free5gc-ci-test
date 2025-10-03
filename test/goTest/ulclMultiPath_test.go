package test

import (
	packetRusher "test/packetRusher"
	pinger "test/pinger"
	"testing"
	"time"
)

var testMpCases = []struct {
	name        string
	destination string
}{
	{
		name:        "DN",
		destination: ONE_IP,
	},
	{
		name:        "MEC",
		destination: MEC_IP,
	},
}

func TestULCLMultiPathCi1(t *testing.T) {
	pr := packetRusher.NewPacketRusher()
	pr.Activate()
	defer pr.Deactivate()

	time.Sleep(5 * time.Second)

	for _, testCase := range testMpCases {
		t.Run(testCase.name, func(t *testing.T) {
			if err := pinger.Pinger(testCase.destination, NIC_1); err != nil {
				t.Errorf("Ping %s failed: expected ping success, but got %v", testCase.destination, err)
			}
		})
	}
}

func TestULCLMultiPathCi2(t *testing.T) {
	pr := packetRusher.NewPacketRusher()
	pr.Activate()
	defer pr.Deactivate()

	time.Sleep(5 * time.Second)

	for _, testCase := range testMpCases {
		t.Run(testCase.name, func(t *testing.T) {
			if err := pinger.Pinger(testCase.destination, NIC_2); err != nil {
				t.Errorf("Ping %s failed: expected ping success, but got %v", testCase.destination, err)
			}
		})
	}
}
