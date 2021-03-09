package factory_method

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {

	var (
		plusfactory  OperatorFactory
		minusfactory OperatorFactory
	)

	plusTests := []struct {
		a        int
		b        int
		expected int
	}{
		{-1000, 0, -1000},
		{1, 1000, 1001},
		{-1000, 999, -1},
		{0, 1000, 1000},
	}
	plusfactory = PlusOperatorFactory{}
	for _, test := range plusTests {
		require.Equal(t, test.expected, compute(plusfactory, test.a, test.b))
	}

	minusTests := []struct {
		a        int
		b        int
		expected int
	}{
		{-1000, 0, -1000},
		{1, 1000, -999},
		{1000, 999, 1},
		{0, 1000, -1000},
	}
	minusfactory = MinusOperatorFactory{}
	for _, test := range minusTests {
		require.Equal(t, test.expected, compute(minusfactory, test.a, test.b))
	}

}
