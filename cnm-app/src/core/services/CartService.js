function updateCartPrice(cart, product) {
  let cartItems = cart.cartItems.map((cartItem) => {
    if (cartItem.productId == product.productId) {
      return { ...cartItem, quantity: product.quantity };
    }
    return { ...cartItem };
  });
  cartItems = cartItems.filter((item) => item.quantity > 0);
  const { cartId, cartName, userId, discountCoupon } = cart;
  return { cartId, cartName, cartItems, userId, discountCoupon };
}

export { updateCartPrice };
