package backend

import (
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"testing"
)

type Preorder struct {
	gorm.Model
	NameBook string `valid:"required~Name not blank"`
	Url      string `gorm:"uniqueIndex" valid:"url~does not validate as url"`
	Email    string `gorm:"uniqueIndex" valid:"email~does not validate as email"`
	Year     string `valid:"required~year not blank"`
	Amount   int
}

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	p := Preorder{
		NameBook: "",
		Url:      "https://www.facebook.com/",
		Email:    "bookname@gmail.com",
		Year:     "2565",
		Amount:   2,
	}
	ok, err := govalidator.ValidateStruct(p)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name not blank"))
}

func TestUrl(t *testing.T) {
	g := NewGomegaWithT(t)

	p := Preorder{
		NameBook: "bookname",
		Url:      "//www.facebook", //ผิด
		Email:    "bookname@gmail.com",
		Year:     "2565",
		Amount:   2,
	}
	ok, err := govalidator.ValidateStruct(p)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("does not validate as url"))
}

func TestEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	p := Preorder{
		NameBook: "bookname",
		Url:      "https://www.facebook.com/",
		Email:    "bookname@gmai", //ผิด
		Year:     "2565",
		Amount:   2,
	}
	ok, err := govalidator.ValidateStruct(p)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("does not validate as email"))
}

func TestYearNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	p := Preorder{
		NameBook: "bookname",
		Url:      "https://www.facebook.com/",
		Email:    "bookname@gmail.com",
		Year:     "", //ผิด
		Amount:   2,
	}
	ok, err := govalidator.ValidateStruct(p)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("year not blank"))

}
