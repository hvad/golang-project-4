package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/shirou/gopsutil/disk"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Hostname string
	Address  string
	Disks    []string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var config Config

	var warning int = 80
	var critical int = 90

	DefaultConfigfile := "/etc/golang-monitor.yml"
	ConfigFile := flag.String("f", DefaultConfigfile, fmt.Sprintf("Config file path, default = %s", DefaultConfigfile))
	flag.Parse()

	CFG, err := ioutil.ReadFile(*ConfigFile)
	check(err)

	err = yaml.Unmarshal(CFG, &config)
	check(err)

	DisplayAddr := config.Address
	DisplayHostname := config.Hostname
	DisplayDisks := config.Disks

	fmt.Printf("Address : %v\n", DisplayAddr)
	fmt.Printf("Hostname : %v\n", DisplayHostname)

	for _, d := range DisplayDisks {
		u, err := disk.Usage(d)
		check(err)
		var DisplayUsage int = int(u.UsedPercent)
		if DisplayUsage >= critical {
			fmt.Printf("Critical - Disk %v percent usage %v %%\n", u.Path, DisplayUsage)

		} else if DisplayUsage >= warning {
			fmt.Printf("Warning - Disk %v percent usage %v %%\n", u.Path, DisplayUsage)

		} else {

			fmt.Printf("Ok - Disk %v percent usage %v %%\n", u.Path, DisplayUsage)

		}

	}

}
