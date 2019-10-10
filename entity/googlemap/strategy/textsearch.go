package strategy

type TextSearch struct{}

func (t TextSearch) formatquery(queries []string) string {
	formattedQuery := queries[0]
	for _, query := range queries[1:] {
		formattedQuery = formattedQuery + "+" + query
	}

	return formattedQuery
}

func (t TextSearch) baseURL() string {
	return "https://maps.googleapis.com/maps/api/place/textsearch/json?query="
}

func (t TextSearch) options() string {
	return "&language=ja"
}
