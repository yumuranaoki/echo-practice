package strategy

type TextSearch struct{}

func (t TextSearch) Formatquery(queries []string) string {
	formattedQuery := queries[0]
	for _, query := range queries[1:] {
		formattedQuery = formattedQuery + "+" + query
	}

	return formattedQuery
}

func (t TextSearch) BaseURL() string {
	return "https://maps.googleapis.com/maps/api/place/textsearch/json?query="
}

func (t TextSearch) Options() string {
	return "&language=ja"
}
