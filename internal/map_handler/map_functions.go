package maphandler

import (
	"encoding/json"
	"errors"
	"os"
)

// Load a map file into a MapInfo Struct and return it
func loadMap(filename string) (*MapData, error) {
	tmpMap := new(MapData)

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("loadMap: Failed to read filepath")
	}

	if err := json.Unmarshal(file, &tmpMap); err != nil {
		return nil, errors.New("loadMap: Failed to unmarshal json into struct")
	}

	tmpMap.CurrentRoom = tmpMap.Rooms[tmpMap.MetaData.StartRoomIndex]

	return tmpMap, nil
}
