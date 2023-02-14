package backend

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Confirmation struct {
	gorm.Model
	Name  string `valid:"required~Name not blank"`
	Email string
}

func TestNameCNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	c := Confirmation{
		Name:  "",
		Email: "maprang@gmail.com",
	}
	ok, err := govalidator.ValidateStruct(c)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name not blank"))

}
