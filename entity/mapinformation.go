package Entity

type MapInformation struct {
	Name           string `json:"name" query:"name"`
	Address        string `json:"address" query:"address"`
	PlaceID        string `json:"place_id" query:"place_id"`
	PhotoReference string `json:"photo_reference" query:"photo_reference"`
	LongName       string `json:"long_name" query:"long_name"`
	Duration       string `json:"duration" query:"duration"`
}
