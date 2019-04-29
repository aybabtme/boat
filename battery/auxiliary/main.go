package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

func main() {
	cellSpecFilename := flag.String("cell", "cell_spec.json", "")
	batterySpecFilename := flag.String("battery", "battery_spec.json", "")
	flag.Parse()

	if *cellSpecFilename == "" {
		log.Fatal("--cell is required")
	}
	if *batterySpecFilename == "" {
		log.Fatal("--battery is required")
	}

	batteryFile, err := os.Open(*batterySpecFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer batteryFile.Close()

	cellFile, err := os.Open(*cellSpecFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer cellFile.Close()

	batterySpec := new(BatterySpec)
	if err := json.NewDecoder(batteryFile).Decode(batterySpec); err != nil {
		log.Fatal(err)
	}

	cellSpec := new(CellSpec)
	if err := json.NewDecoder(cellFile).Decode(cellSpec); err != nil {
		log.Fatal(err)
	}

	result := cellSpec.AsBattery(*batterySpec)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	if err := enc.Encode(result); err != nil {
		log.Fatal(err)
	}
}
