package stmt_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/node/stmt"
	"github.com/i582/php2go/src/php/php5"
	"github.com/i582/php2go/src/php/php7"
	"github.com/i582/php2go/src/php/position"
)

func TestTrait(t *testing.T) {
	src := `<? trait Foo {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    15,
		},
		Stmts: []node.Node{
			&stmt.Trait{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    15,
				},
				PhpDocComment: "",
				TraitName: &node.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  9,
						EndPos:    12,
					},
					Value: "Foo",
				},
				Stmts: []node.Node{},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
