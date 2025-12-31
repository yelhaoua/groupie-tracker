package handlers

/*
	This page contains all structs used in the project
*/


// For  Handler Error
type info struct {
	Code        int
	Description string
}

// For  Save Artist data
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// For  Save Artist Locations
type LocationStruct struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

// For  Save Artist Dates
type DatesStruct struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

// For  Save Artist Relations
type RelationStruct struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// For  Final output for info page
type AllData struct {
	Artist   Artist
	Location LocationStruct
	Dates    DatesStruct
	Relation RelationStruct
}
