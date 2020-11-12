package orders

import (
	"encoding/json"
	"errors"
	"shopping-cart/cnm-carts/models"
	ce "shopping-cart/cnm-carts/models"
	cartService "shopping-cart/cnm-carts/services"
	"shopping-cart/cnm-core/constants"
	"shopping-cart/cnm-core/utils"
	"shopping-cart/cnm-orders/entities"
	repository "shopping-cart/cnm-orders/repository"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides booking methods.
type  Service interface {
	add(request addOrdersRequest) (addOrdersResponse, error)
	get(request  GetOrderRequest) (getOrderListResponse, error)
	CartToOrder(ce.CartResponse) entities.Orders
}

type service struct {
}

func (s *service) add(request addOrdersRequest) (addOrdersResponse, error) {
	instance:= repository.OrderRepositoryInstance()
	cartServiceInstance:=cartService.NewService()
	var cartRequest models.GetCartRequest
	cartRequest.UserId=request.UserId
	var cartResponse, _ =cartServiceInstance.Get(cartRequest)
	//TODO fetch order from cart
	var userOrder=s.CartToOrder(cartResponse);
	orderAdded, err:=instance.Add(userOrder)

    token,_:=utils.PaymentToken(userOrder.OrderId,userOrder.UserId, userOrder.Amount)
	//clear order
	return  addOrdersResponse{
		Orders: orderAdded,
		PaymentHost:"http://localhost:3501",
		PaymentUrl: "/payment",
		Token:token,
	}, err
}

func (s *service) get(request  GetOrderRequest) (getOrderListResponse, error) {
	repository:= repository.OrderRepositoryInstance()
	cart, err:=repository.List(request.UserId);
	return getOrderListResponse{Orders: cart}, err
}


func (s *service) CartToOrder(cart ce.CartResponse) (entities.Orders) {
	//Convert Cart items into order
	var oderItems [] entities.OrdersItems;
	b, _:=json.Marshal(cart)
	for _, cartItem := range cart.CartItems {
		oderItems = append(oderItems,entities.OrdersItems{
			ProductId : cartItem.Product.ProductId,
			ProductName : cartItem.Product.ProductName,
			ProductPrice :cartItem.Product.ProductPrice,
			Quantity: cartItem.Quantity,

		})
	}
	 order:= entities.Orders{
		  UserId :  cart.UserId,
		  CartId :  cart.CartId,
		  Amount :  cart.TotalAmount,
		  Status :  constants.PENDING_DELIVERY,
		  OrdersItems: oderItems,
		  OrderData: string(b),
	 }

	return order
}


// NewService creates a  service with necessary dependencies.
func NewService() Service {
	return &service{}
}

