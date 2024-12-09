package test

import (
	"os/exec"
	"testing"
	"time"

	packetRusher "test/packetRusher"
	pinger "test/pinger"
)

func TestULCLTrafficInfluence(t *testing.T) {
	var (
		err    error
		output []byte
		cmd    *exec.Cmd
		pr     *packetRusher.PacketRusher
	)

	// activate PacketRusher
	pr = packetRusher.NewPacketRusher()
	pr.Activate()

	time.Sleep(3 * time.Second)

	// before TI
	t.Run("Before TI", func(t *testing.T) {
		err = pinger.Pinger(N6GW_IP, NIC)
		if err != nil {
			t.Errorf("Ping n6gw failed: expected ping success, but got %v", err)
		}
		err = pinger.Pinger(MEC_IP, NIC)
		if err == nil {
			t.Errorf("Ping mec failed: expected ping failed, but got %v", err)
		}
	})

	// post TI
	cmd = exec.Command("bash", "../api-udr-ti-data-action.sh", "put")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Insert ti data failed: expected insert success, but got %v, output: %s", err, output)
	}
	time.Sleep(3 * time.Millisecond)

	// after TI
	t.Run("After TI", func(t *testing.T) {
		err = pinger.Pinger(N6GW_IP, NIC)
		if err != nil {
			t.Errorf("Ping n6gw failed: expected ping failed, but got %v", err)
		}
		err = pinger.Pinger(MEC_IP, NIC)
		if err == nil {
			t.Errorf("Ping mec failed: expected ping sucess, but got %v", err)
		}
	})

	// delete TI
	cmd = exec.Command("bash", "../api-udr-ti-data-action.sh", "delete")
	output, err = cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Delete ti data failed: expected delete success, but got %v, output: %s", err, output)
	}
	time.Sleep(500 * time.Millisecond)

	// reset TI
	t.Run("Reset TI", func(t *testing.T) {
		err = pinger.Pinger(N6GW_IP, NIC)
		if err != nil {
			t.Errorf("Ping n6gw failed: expected ping success, but got %v", err)
		}
		err = pinger.Pinger(MEC_IP, NIC)
		if err == nil {
			t.Errorf("Ping mec failed: expected ping failed, but got %v", err)
		}
	})

	// deactivate PacketRusher
	pr.Deactivate()
}
