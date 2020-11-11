function updateCartPrice(cart, product) {
  let cartItems = cart.cartItems.map((cartItem) => {
    if (cartItem.productId === product.productId) {
      return { ...cartItem, quantity: product.quantity };
    }
    return { ...cartItem };
  });
  cartItems = cartItems.filter((item) => item.quantity > 0);

  const { cartId, cartName, userId, discountCoupon } = cart;
  return { cartId, cartName, cartItems, userId, discountCoupon };
}

const filterCartItem = (cart, product) => {
  const { cartId, userId } = cart;
  const cartItemIds = [];

  cart.cartItems.forEach((item) => {
    if (item.productId === product.productId) {
      cartItemIds.push(item.cartItemsId);
    }
  });

  const modifiedCart = {
    cartId,
    userId,
    cartItemIds,
  };

  return modifiedCart;
};

const deleteCartItemPayload = (cart, product) => {
  const { cartId, userId } = cart;
  const cartItemIds = [];
  cart.cartItems.forEach((item) => {
    cartItemIds.push(item.cartItemsId);
  });

  const modifiedCart = {
    cartId,
    userId,
    cartItemIds,
  };

  return modifiedCart;
};

export { updateCartPrice, filterCartItem, deleteCartItemPayload };
