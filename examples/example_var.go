package examples

//go:generate mockgen -source=$GOFILE -destination=mock/${GOFILE}_mock.go -package=${GOPACKAGE}_mock

type PrefixInterface interface {
	AddPrefix(prefix, str string) string
	AddPostfix(postfix, str string) string
}

type prefix struct{}

func NewPrefixInterface() PrefixInterface {
	return prefix{}
}

func (prefix) AddPrefix(prefix, str string) string {
	return prefix + str
}

func (prefix) AddPostfix(postfix, str string) string {
	return str + postfix
}
