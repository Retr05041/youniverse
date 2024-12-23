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

// Test load map for if a file is not found
func TestLoadMap_FileNotFound(t *testing.T) {
	filename := "non_existent_file.json"

	_, err := loadMap(filename)
	if err == nil || err.Error() != "loadMap: Failed to read filepath" {
		t.Fatalf("Expected error for missing file, got %v", err)
	}
}

// Test load map on an invalid json file
func TestLoadMap_InvalidJSON(t *testing.T) {
	filename := "invalid_json.json"
	invalidJSON := `{"Rooms": ["Name": "Room1"]}` // Invalid JSON format
	_ = os.WriteFile(filename, []byte(invalidJSON), 0644)
	defer os.Remove(filename)

	_, err := loadMap(filename)
	if err == nil || err.Error() != "loadMap: Failed to unmarshal json into struct" {
		t.Fatalf("Expected error for invalid JSON, got %v", err)
	}
}

// Missing Start Room Index - might need to check more than this in terms of missing values? TODO: Validate this
func TestLoadMap_MissingStartRoomIndex(t *testing.T) {
	filename := "missing_start_room.json"
	missingStartRoomJSON := `{
	"METADATA": {"start": null, end": 0},
		"GAME": [{
			"index": 0,
			"name": "Your Room",
			"needed item": null,
			"north": null,
			"east": null,
			"south": null,
			"west": null,
			"item": null,
			"look": "A basic room description."
		}]
	}`
	_ = os.WriteFile(filename, []byte(missingStartRoomJSON), 0644)
	defer os.Remove(filename)

	_, err := loadMap(filename)
	if err == nil {
		t.Fatalf("Expected error for missing StartRoomIndex, got nil")
	}
}

// Empy map file
func TestLoadMap_EmptyFile(t *testing.T) {
	filename := "empty_file.json"
	_ = os.WriteFile(filename, []byte(""), 0644)
	defer os.Remove(filename)

	_, err := loadMap(filename)
	if err == nil || err.Error() != "loadMap: Failed to unmarshal json into struct" {
		t.Fatalf("Expected error for empty file, got %v", err)
	}
}
