package maphandler

// Global struct for holding all map data
type MapData struct {
	MetaData    metaData   `json:"METADATA"`
	Rooms       []roomData `json:"GAME"`
	CurrentRoom roomData
}

// Meta data struct for database info
type metaData struct {
	StartRoomIndex int `json:"start"`
	EndRoomIndex   int `json:"end"`
}

// Each room will be unmarshaled to this struct for use within the program
type roomData struct {
	Index      int     `json:"index"`
	Name       string  `json:"name"`
	NeededItem *string `json:"needed item"`
	North      *int    `json:"north"`
	East       *int    `json:"east"`
	South      *int    `json:"south"`
	West       *int    `json:"west"`
	Item       *string `json:"item"`
	Look       *string `json:"look"`
}
