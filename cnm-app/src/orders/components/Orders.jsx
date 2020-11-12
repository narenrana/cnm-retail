import React, { useEffect } from "react";
import Grid from "@material-ui/core/Grid";

import Typography from "@material-ui/core/Typography";
import useStyles from "./style";
import _ from "lodash";

import { useDispatch, useSelector } from "react-redux";
import { getOrders } from "../redux";
import OrderItem from "./OrderItem";

import CardHeader from "@material-ui/core/CardHeader";
import Card from "@material-ui/core/Card";

export default function OrdersList() {
  const classes = useStyles();
  const dispatch = useDispatch();

  const { orders = [], isLoading = {} } = useSelector(
    (state) => state.ordersStore
  );

  useEffect(() => {
    if (!isLoading) {
      dispatch(getOrders());
    }
  }, [dispatch, isLoading]);

  const getOrderData = (order) => {
    return JSON.parse(order.orderData) || { cartItems: [] };
  };

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom className={classes.listHeader}>
        Orders
      </Typography>
      {orders.map((order) => (
        <Card className={classes.root} spacing={4}>
          <CardHeader title={`$${getOrderData(order).totalAmount || 0}`} />
          <Grid container spacing={4}>
            {(getOrderData(order) || {}).cartItems.map((cartItem, index) => (
              <Grid
                item
                xs={6}
                md={3}
                xl={3}
                key={"suggestion-key-" + index}
                spacing={4}
              >
                <OrderItem
                  product={cartItem.products}
                  cartQuantity={cartItem.quantity}
                  isExpanded={false}
                  cart={cartItem}
                  onAddToCart={() => {}}
                  onIncreaseQuantity={() => {}}
                  onReduceQuantity={() => {}}
                />
              </Grid>
            ))}
          </Grid>
        </Card>
      ))}
    </React.Fragment>
  );
}
