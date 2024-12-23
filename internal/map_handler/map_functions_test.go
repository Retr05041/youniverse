package maphandler

import (
	"os"
	"testing"
)

// Test if a valid map can be loaded
func TestLoadMap_ValidMapFile(t *testing.T) {
	// Dumby file
	filename := "valid_map.json"
	expectedStartRoomIndex := 0
	// Mock data
	validJSON := `
	{
		"METADATA": {
			"start": 0,
			"end": 1
		},
		"GAME": [
			{
				"index": 0,
				"name": "Room1",
				"needed item": null,
				"north": null,
				"east": null,
				"south": 1,
				"west": null,
				"item": null,
				"look": null
			},
			{
				"index": 1,
				"name": "Room2",
				"needed item": null,
				"north": 0,
				"east": null,
				"south": null,
				"west": null,
				"item": null,
				"look": null
			}
		]
	}
	`
	// Write it to a file for testing, defer the deletion of the file
	_ = os.WriteFile(filename, []byte(validJSON), 0644)
	defer os.Remove(filename)

	// Test
	mapData, err := loadMap(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if mapData.CurrentRoom.Name != "Room1" {
		t.Fatalf("Expected CurrentRoom to be Room1, got %v", mapData.CurrentRoom.Name)
	}
	if mapData.MetaData.StartRoomIndex != expectedStartRoomIndex {
		t.Fatalf("Expected StartRoomIndex to be %d, got %d", expectedStartRoomIndex, mapData.MetaData.StartRoomIndex)
	}
}
