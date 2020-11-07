package payments

import (
	"github.com/jackc/pgtype"
	core "shopping-cart/cnm-core"
)

type Service interface {
	List() ([] Payments,error)
	FindBy(u Payments) (  Payments,error)
	Add(u Payments) ( Payments,error)
	Delete(u Payments) ( Payments,error)
}

type Payments struct {

	PaymentId         int          `gorm:"primaryKey" json:"description,omitempty"`
	UserId            string       `json:"userId,omitempty"`
	PaymentMode       string       `json:"paymentMode,omitempty"`
	OrderId           string       `json:"orderId,omitempty"`
	OrderDate         string       `json:"orderDate,omitempty"`
	OrderAmount       string       `json:"orderAmount,omitempty"`
	PaymentStatus     string       `json:"PaymentStatus,omitempty"`
	NameOnCard        pgtype.Date  `json:"nameOnCard,omitempty"`
	CardType          pgtype.Date  `json:"cardType,omitempty"`
	CardLastDigit     string       `json:"cardLastDigit,omitempty"`
	active            bool         `json:"active,omitempty"`
	DateCreated   pgtype.Date    `json:"dateCreated,omitempty"`
	DateUpdated   pgtype.Date    `json:"dateUpdated,omitempty"`
}

func (*Payments) List() ([] Payments, error){
	db,err :=core.GetDB()
	var found [] Payments;
	if err != nil {
		return nil, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Payments) FindBy(item Payments) (Payments, error){
	db,err :=core.GetDB()
	 var found Payments;
	if err != nil {
		return item, err;
	}
	db.Find(&found);
	return  found, err;
}

func (*Payments) Add(item Payments) (Payments, error){
	db,err :=core.GetDB()
	if err != nil {
		return item, err;
	}
	db.Create(&item);
	return  item, err;
}

func (*Payments) Delete(item Payments) (Payments, error){
	db,err :=core.GetDB()
	if err != nil {
		return item, err;
	}
	db.Delete(&item);
	return  item, err;
}

func PaymentsRepositoryInstance() Service {
	return &Payments{};
}

