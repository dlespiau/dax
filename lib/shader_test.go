package dax

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testFragmentShaderSource = `
#version 330
out vec4 outputColor;
void main() {
    outputColor = vec4(.8, .8, .8, 1);
}`

const testUniformName = "foo"

func TestCStrings(t *testing.T) {
	s := NewFragmentShader(testFragmentShaderSource)

	for i := UniformKind(0); i < uniformKindMax; i++ {
		s.AddUniform(i, fmt.Sprintf("%s-%d", testUniformName, i))
	}

	assert.Equal(t, testFragmentShaderSource+"\x00", s.source)
	assert.Equal(t, int(uniformKindMax), len(s.uniforms))
	for i := range s.uniforms {
		name := fmt.Sprintf("%s-%d", testUniformName, i)
		assert.Equal(t, name, s.uniforms[i].Name())
	}
}
