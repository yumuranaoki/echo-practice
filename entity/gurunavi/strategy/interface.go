package strategy

type Strategy interface {
	build() string
	parse() Restaurant
}
