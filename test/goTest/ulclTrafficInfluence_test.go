package test

import (
	"encoding/json"
	"os/exec"
	"strings"
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
	time.Sleep(3 * time.Millisecond)

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

	// check charging record
	t.Run("Check Charging Record", func(t *testing.T) {
		checkChargingRecord(t)
	})

	// deactivate PacketRusher
	pr.Deactivate()
}

func checkChargingRecord(t *testing.T) {
	cmd := exec.Command("bash", "../api-webconsole-charging-record.sh", "get", "../json/webconsole-login-data.json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Get charging record failed: %v, output: %s", err, output)
		return
	}

	outputStr := string(output)

	lines := strings.Split(outputStr, "\n")
	var jsonLine string
	for i := len(lines) - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) != "" {
			jsonLine = lines[i]
			break
		}
	}

	var chargingRecords []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonLine), &chargingRecords); err != nil {
		t.Errorf("Failed to parse charging record JSON: %v\nJSON content: %s", err, jsonLine)
		return
	}

	if len(chargingRecords) == 0 {
		t.Error("No charging records found")
		return
	}

	if chargingRecords[0]["TotalVol"].(float64) == 0 {
		t.Error("Charging record is empty")
	}
}
