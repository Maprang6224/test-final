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
	Amount   int    `valid:"required~ราคาหนังสือไม่ถูกต้อง, range(1|5)~ราคาหนังสือไม่ถูกต้อง"`
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

func TestAmountNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	fixture := []int{
		-1, 0, 6}

	for _, price := range fixture {
		pr := Preorder{
			NameBook: "bookname",
			Url:      "http://www.facebook.com",
			Email:    "maprang@gmail.com",
			Year:     "2555",
			Amount:   price,
		}

		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(pr)

		//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ err ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("ราคาหนังสือไม่ถูกต้อง"))

	}

}
