package main

import (
	"reflect"
	"testing"
)

func TestCalculateSubnets(t *testing.T) {
	egs := []struct {
		network Network
		subnets []Subnet
		err     string
	}{
		{
			network: Network{IP: "11.22.33.44"},
			err:     "invalid CIDR address: 11.22.33.44",
		},
		{
			network: Network{
				IP: "11.22.33.44/24",
				Subnets: []Subnet{
					{
						Mode: Minimum,
						Size: 255,
					},
				},
			},
			err: "Network is too small. Subnet requested 255, max available is 254",
		},
		{
			network: Network{
				IP: "11.22.33.44/24",
				Subnets: []Subnet{
					{
						Mode: Minimum,
						Size: 10,
					},
				},
			},
			subnets: []Subnet{
				{
					Mode:      Minimum,
					Size:      16,
					IP:        "11.22.33.0/28",
					Mask:      "255.255.255.240",
					RangeMin:  "11.22.33.1",
					RangeMax:  "11.22.33.14",
					Broadcast: "11.22.33.15",
				},
			},
		},
		{
			network: Network{
				IP: "192.168.1.2/30",
				Subnets: []Subnet{
					{
						Mode: Minimum,
						Size: 2,
					},
				},
			},
			subnets: []Subnet{
				{
					Mode:      Minimum,
					Size:      4,
					IP:        "192.168.1.0/30",
					Mask:      "255.255.255.252",
					RangeMin:  "192.168.1.1",
					RangeMax:  "192.168.1.2",
					Broadcast: "192.168.1.3",
				},
			},
		},
		{
			network: Network{
				IP: "192.137.28.3/26",
				Subnets: []Subnet{
					{
						Mode: Minimum,
						Size: 20,
					},
				},
			},
			subnets: []Subnet{
				{
					Mode:      Minimum,
					Size:      32,
					IP:        "192.137.28.0/27",
					Mask:      "255.255.255.224",
					RangeMin:  "192.137.28.1",
					RangeMax:  "192.137.28.30",
					Broadcast: "192.137.28.31",
				},
			},
		},
	}

	for _, eg := range egs {
		got, err := calculateSubnets(eg.network)
		if err != nil && eg.err != err.Error() {
			t.Errorf("CalculateSubnets(%v)\n exp: error %s\n got: %s\n\n", eg.network, eg.err, err)
			continue
		}

		if !reflect.DeepEqual(eg.subnets, got) {
			t.Errorf("CalculateSubnets(%v)\n exp: %v\n got: %v\n\n", eg.network, eg.subnets, got)
		}

	}
}
