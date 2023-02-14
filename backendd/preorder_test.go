package backendd

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"testing"
)

type Preorder struct {
	gorm.Model
	Name  string `valid:"required~Name not blank"`
	Email string `valid:"email~does not validate as email"`
	Url   string `valid:"url~does not validate as url"`
	Age   int    `valid:"required~age must be in range, range(1|29)~age must be in range"`
}

func TestCorrectAll(t *testing.T) {
	g := NewGomegaWithT(t)

	pre := Preorder{
		Name:  "Maprang",
		Email: "Maprang@gmail.com",
		Url:   "http://www.google.com",
		Age:   22,
	}
	ok, err := govalidator.ValidateStruct(pre)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

	fmt.Println(err)
}

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	pre := Preorder{
		Name:  "", //wrong
		Email: "Maprang@gmail.com",
		Url:   "http://www.google.com",
		Age:   22,
	}
	ok, err := govalidator.ValidateStruct(pre)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Name not blank"))
}

func TestEmailIsMatch(t *testing.T) {
	g := NewGomegaWithT(t)

	pre := Preorder{
		Name:  "Maprang",
		Email: "Maprang@", //wrong
		Url:   "http://www.google.com",
		Age:   22,
	}
	ok, err := govalidator.ValidateStruct(pre)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("does not validate as email"))
}

func TestUrlValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	pre := Preorder{
		Name:  "Maprang",
		Email: "Maprang@gmail.com",
		Url:   "//www.google.", //wrong
		Age:   22,
	}
	ok, err := govalidator.ValidateStruct(pre)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("does not validate as url"))
}

func TestAgeNoTZero(t *testing.T) {
	g := NewGomegaWithT(t)

	ages := []int{
		0,
		-1,
		30,
	}

	for _, age := range ages {
		pre := Preorder{
			Name:  "Maprang",
			Email: "Maprang@gmail.com",
			Url:   "http://www.google.com",
			Age:   age,
		}
		ok, err := govalidator.ValidateStruct(pre)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())

		g.Expect(err.Error()).To(Equal("age must be in range"))
	}
}
