import React, { useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import { Product } from "../../../products";
import _ from "lodash";
import { updateCartPrice } from "../../../core";
import CartDetails from "../cart-details/CartDetails";
import { useDispatch, useSelector } from "react-redux";
import {
  addCartItem,
  updateCart,
  getCart,
  getProducts,
} from "../../../products/redux";

const useStyles = makeStyles((theme) => ({
  root: {
    "& > *": {
      marginTop: theme.spacing(2),
    },
  },
}));

export default function Checkout(props) {
  const dispatch = useDispatch();
  const { products = [], isLoading, cart = {} } = useSelector(
    (state) => state.productsStore
  );
  const classes = useStyles();

  const { cartItems = [] } = cart;
  const cartItemsMap = {};
  cartItems.map((cartItem) => {
    cartItemsMap[cartItem.productId] = cartItem;
  });

  const onAddToCart = (product) => {
    const { productId, productName, productPrice, quantity = 1 } = product;
    dispatch(addCartItem({ productId, productName, productPrice, quantity }));
  };

  const isExpanded = (product) => {
    return !_.isNil(cartItemsMap[product.productId]);
  };

  const cartQuantity = (product) => {
    return _.get(cartItemsMap, "[" + product.productId + "].quantity", 0);
  };

  const onIncreaseQuantity = (product) => {
    const updatedCart = updateCartPrice(cart, product);
    dispatch(updateCart(updatedCart));
  };

  const onReduceQuantity = (product) => {
    const updatedCart = updateCartPrice(cart, product);
    dispatch(updateCart(updatedCart));
  };

  const getCartProduct = (cartItem) => {
    return products.find((product) => product.productId === cartItem.productId);
  };

  const addCouponToCart = (discountCoupon) => {
    const { userId, cartId, cartItems } = cart;
    dispatch(
      updateCart({
        userId,
        cartId,
        cartItems,
        discountCoupon,
      })
    );
  };

  useEffect(() => {
    if (!isLoading) {
      dispatch(getProducts());
      dispatch(getCart());
    }
  }, [dispatch, isLoading]);

  return (
    <div className={classes.root}>
      <Grid container>
        <Grid item xs={6} md={6} xl={6}>
          <Grid container spacing={4}>
            {cartItems.map((cartItem, index) => (
              <Grid item xs={6} md={6} xl={3} key={"suggestion-key-" + index}>
                <Product
                  product={getCartProduct(cartItem)}
                  cartQuantity={cartQuantity(cartItem)}
                  isExpanded={isExpanded(cartItem)}
                  onAddToCart={onAddToCart}
                  onIncreaseQuantity={onIncreaseQuantity}
                  onReduceQuantity={onReduceQuantity}
                />
              </Grid>
            ))}
          </Grid>
        </Grid>
        <Grid item xs={6} md={6} xl={6}>
          <CartDetails cart={cart} addCouponToCart={addCouponToCart} />
        </Grid>
      </Grid>
    </div>
  );
}
