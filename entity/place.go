package Entity

type Place struct {
	Name           string `json:"name" query:"name"`
	Address        string `json:"address" query:"address"`
	PlaceID        string `json:"place_id" query:"place_id"`
	PhotoReference string `json:"photo_reference" query:"photo_reference"`
}
