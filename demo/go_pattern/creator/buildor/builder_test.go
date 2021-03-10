package Buildor

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuilder1(t *testing.T) {

	tests := []struct {
		InitValue string
		Excepted  string
	}{
		{"xxxx", "xxxx123"},
		{"_x81", "_x81123"},
		{"", "123"},
	}

	for _, item := range tests {
		b1 := Builder1{item.InitValue}
		d1 := NewDirector(&b1)
		d1.Construct()
		require.Equal(t, item.Excepted, b1.GetResult())
	}
}

func TestBuilder2(t *testing.T) {

	tests := []struct {
		InitValue int
		Excepted  int
	}{
		{-6, -6},
		{1234, 1240},
		{-1000, -994},
		{0, 6},
	}

	for _, item := range tests {
		b2 := Builder2{item.InitValue}
		d2 := NewDirector(&b2)
		d2.Construct()
		require.Equal(t, item.Excepted, b2.GetResult())
	}
}
