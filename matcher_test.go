package withgomega

import (
	"testing"

	"github.com/onsi/gomega"
)

func Test_Basic(t *testing.T) {
	tt := gomega.NewGomegaWithT(t)

	matcher := Matcher{}

	tt.Expect(1).To(matcher.Equal(1))
}
