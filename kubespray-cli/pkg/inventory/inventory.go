package inventory

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/toszo/kubespray-cli/pkg/constants"
	"gopkg.in/yaml.v2"
)

func Init() {
	inventory := InventoryWrapper{
		Inventory: Inventory{
			Vars:  InventoryVars{},
			Hosts: map[string]Node{},
		},
	}
	saveState(&inventory)
}

func Apply() {
	log.Println("Nothing to see here yet.")
}

func saveState(model *InventoryWrapper) {
	stateLocation, state := readStateFile()

	d, err := yaml.Marshal(&model)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	state["kubespray"] = string(d)

	stateUpdated, err := yaml.Marshal(&state)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	file, err := os.Create(stateLocation)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	l, err := file.WriteString(string(stateUpdated))
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	log.Println(l, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getStateFileLocation() string {
	statefile := os.Getenv(constants.StateFileLocationVarName)
	sharedLoc := os.Getenv(constants.SharedDirVarName)
	return fmt.Sprintf("%s/%s", sharedLoc, statefile)
}

func readStateFile() (string, map[string]string) {

	stateFileLoc := getStateFileLocation()
	log.Println("Reading file: ", stateFileLoc)
	yamlFile, err := ioutil.ReadFile(stateFileLoc)
	if err != nil {
		log.Printf("#%v ", err)
	}
	var result map[string]string
	errUnmarhal := yaml.Unmarshal(yamlFile, &result)
	if errUnmarhal != nil {
		log.Printf("#%v ", err)
	}
	return stateFileLoc, result
}
