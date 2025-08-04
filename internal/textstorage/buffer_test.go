package textstorage_test

import (
	"reflect"
	"testing"

	"dido/internal/textstorage"
)

func TestBufferInsert(t *testing.T) {
	tests := map[string]struct {
		buffer   textstorage.Buffer
		input    rune
		idx      int
		expected []rune
	}{
		`empty buffer`: {
			buffer:   textstorage.NewBuffer([]rune{0}),
			input:    'x',
			idx:      0,
			expected: []rune{'x'},
		},
		`insert at the beginning`: {
			buffer:   textstorage.NewBuffer([]rune{'a', 'b', 'c'}),
			input:    'x',
			idx:      0,
			expected: []rune{'x', 'a', 'b', 'c'},
		},
		`insert in the middle`: {
			buffer:   textstorage.NewBuffer([]rune{'a', 'b', 'c'}),
			input:    'x',
			idx:      2,
			expected: []rune{'a', 'b', 'x', 'c'},
		},
		`insert at the end`: {
			buffer:   textstorage.NewBuffer([]rune{'a', 'b', 'c'}),
			input:    'x',
			idx:      3,
			expected: []rune{'a', 'b', 'c', 'x'},
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := tt.buffer.Insert(tt.input, tt.idx); err != nil {
				t.Fail()
			}
			got := tt.buffer.Get()
			if !reflect.DeepEqual(*got, tt.expected) {
				t.Errorf("got: `%s`, expected: `%s`", string(*got), string(tt.expected))
			}
		})
	}
}
