package scanner_test

import (
	"reflect"
	"testing"

	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/scanner"
)

func TestToken(t *testing.T) {
	tkn := &scanner.Token{
		Value:     `foo`,
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	c := []freefloating.String{
		{
			Value:      "test comment",
			StringType: freefloating.CommentType,
			Position:   nil,
		},
	}

	tkn.FreeFloating = c

	if !reflect.DeepEqual(tkn.FreeFloating, c) {
		t.Errorf("comments are not equal\n")
	}

	if tkn.String() != `foo` {
		t.Errorf("token value is not equal\n")
	}
}
