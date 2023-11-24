package examples

//go:generate bash -c "mockgen -source=$GOFILE -destination=mock/$(basename $GOFILE .go)_mock.go -package=${GOPACKAGE}_mock"

type CounterInterface interface {
	PlusOne(i int) int
	MinusOne(i int) int
}

type counter struct{}

func NewCounterInterface() CounterInterface {
	return counter{}
}

func (counter) PlusOne(i int) int {
	return i + 1
}

func (counter) MinusOne(i int) int {
	return i - 1
}
