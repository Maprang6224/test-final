package backendd

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"testing"
)

type Confirm struct {
	gorm.Model
	Name  string `valid:"required~name not be blank"`
	Age   int    `valid:"required~age must be in range, range(1|149)~age must be in range"`
	Email string `valid:"email~does not validate as email"`
	Url   string `valid:"url~dose not validate as url"`
}

func TestConfirmCorrectAll(t *testing.T) {
	g := NewGomegaWithT(t)

	confirm := Confirm{
		Name:  "Maprang",
		Age:   22,
		Email: "Maprang@gmail.com",
		Url:   "http://www.google.com",
	}
	ok, err := govalidator.ValidateStruct(confirm)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

	fmt.Println(err)
}

func TestConfirmNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	confirm := Confirm{
		Name:  "", //wrong
		Age:   22,
		Email: "Maprang@gmail.com",
		Url:   "http://www.google.com",
	}
	ok, err := govalidator.ValidateStruct(confirm)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("name not be blank"))
}

func TestConfirmAgeNotZero(t *testing.T) {
	g := NewGomegaWithT(t)

	ages := []int{
		-1, 0, 150,
	}

	for _, age := range ages {
		confirm := Confirm{
			Name:  "Maprang",
			Age:   age, //wrong
			Email: "Maprang@gmail.com",
			Url:   "http://www.google.com",
		}
		ok, err := govalidator.ValidateStruct(confirm)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())

		g.Expect(err.Error()).To(Equal("age must be in range"))
	}
}

func TestConfirmEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	confirm := Confirm{
		Name:  "Maprang",
		Age:   22,
		Email: "Maprang@gm", //wrong
		Url:   "http://www.google.com",
	}
	ok, err := govalidator.ValidateStruct(confirm)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("does not validate as email"))
}

func TestConfirmUrl(t *testing.T) {
	g := NewGomegaWithT(t)

	confirm := Confirm{
		Name:  "Maprang",
		Age:   22,
		Email: "Maprang@gmail.com",
		Url:   "//www.google.c",
	}
	ok, err := govalidator.ValidateStruct(confirm)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("dose not validate as url"))
}
