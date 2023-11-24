//go:generate mockgen -source=$GOFILE -destination=mock/${GOFILE}_mock.go -package=${GOPACKAGE}_mock

package examples

type PrefixCounter interface {
	AddCountPrefix(i int, prefix, str string) string
	AddCountPostfix(i int, postfix, str string) string
}

type prefixCounter struct {
	counter CounterInterface
	prefix  PrefixInterface
}

func NewPrefixCounter(c CounterInterface, p PrefixInterface) PrefixCounter {
	return prefixCounter{
		counter: c,
		prefix:  p,
	}
}

func (p prefixCounter) AddCountPrefix(i int, prefix, str string) string {
	for count := 0; count < i; count = p.counter.PlusOne(count) {
		str = p.prefix.AddPrefix(prefix, str)
	}
	return str
}

func (p prefixCounter) AddCountPostfix(i int, postfix, str string) string {
	for count := 0; count < i; count = p.counter.PlusOne(count) {
		str = p.prefix.AddPostfix(postfix, str)
	}
	return str
}
