package alexa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeManager(t *testing.T) {
	attributes := map[string]interface{}{"attr1": 1, "attr2": "me"}

	bam := BasicAttributeManger{
		Attributes: attributes,
	}
	bam.SetSessionAttributes(map[string]interface{}{"attr3": 3})
	after := bam.GetSessionAttributes()

	assert.Contains(t, after, "attr1")
	assert.Contains(t, after, "attr3")
}
