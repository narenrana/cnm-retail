import React, { useEffect } from "react";
import Grid from "@material-ui/core/Grid";
import { Product } from "../../";
import Pagination from "../../../core/Pagination";
import Typography from "@material-ui/core/Typography";
import useStyles from "../../style";
import _ from "lodash";
import { updateCartPrice } from "../../../core";

import { useDispatch, useSelector } from "react-redux";
import { getProducts, addCartItem, getCart, updateCart } from "../../redux";

export default function ProductList() {
  const dispatch = useDispatch();
  const { products = [], isLoading, cart = {} } = useSelector(
    (state) => state.productsStore
  );

  const { cartItems = [] } = cart;
  const cartItemsMap = {};

  cartItems.map((cartItem) => {
    cartItemsMap[cartItem.productId] = cartItem;
  });

  const onAddToCart = (product) => {
    const { productId, quantity = 0 } = product;
    dispatch(addCartItem({ productId, quantity, cartId: cart.cartId }));

    //const updatedCart = updateCartPrice(cart, { productId, quantity });
    //dispatch(updateCart(updatedCart));
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

  const classes = useStyles();

  useEffect(() => {
    if (!isLoading) {
      dispatch(getProducts());
      dispatch(getCart());
    }
  }, [isLoading]);

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom className={classes.listHeader}>
        Fruits
      </Typography>
      <Grid container spacing={4}>
        {_.sortBy(products, ["productId"]).map((product, index) => (
          <Grid item xs={12} md={3} xl={3} key={"suggestion-key-" + index}>
            <Product
              product={product}
              cartQuantity={cartQuantity(product)}
              isExpanded={isExpanded(product)}
              onAddToCart={onAddToCart}
              onIncreaseQuantity={onIncreaseQuantity}
              onReduceQuantity={onReduceQuantity}
            />
          </Grid>
        ))}
      </Grid>
      <Pagination count={10} />
    </React.Fragment>
  );
}
