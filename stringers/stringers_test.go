package stringers

import (
	"testing"
)

func TestStringers(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		if name == "loopback" && ip.String() != "127.0.0.1" {
			t.Errorf("ip.String() == %v, want match for '127.0.0.1'", ip.String())
		} else if name == "googleDNS" && ip.String() != "8.8.8.8" {
			t.Errorf("ip.String() == %v, want match for '8.8.8.8'", ip.String())
		}
	}
}
