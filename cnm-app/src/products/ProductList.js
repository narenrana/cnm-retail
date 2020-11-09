import React, { useEffect } from "react";
import Grid from "@material-ui/core/Grid";
import Product from "./Product";
import Pagination from "../core/Pagination";
import Typography from "@material-ui/core/Typography";
import useStyles from "./style";
import _ from "lodash";

import { useDispatch, useSelector } from "react-redux";
import { getProducts, addCartItem, getCart, updateCartItem } from "./redux";

export default function ProductList() {
  const dispatch = useDispatch();
  const { products = [], isLoading, cart = {} } = useSelector(
    (state) => state.productsStore
  );

  //const { cart } = useSelector((state) => state.productsStore);
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
    dispatch(updateCartItem(product));
  };

  const onReduceQuantity = (product) => {
    dispatch(updateCartItem(product));
  };

  const classes = useStyles();
  const handleSortBy = () => {};

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
        {products.map((product, index) => (
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
