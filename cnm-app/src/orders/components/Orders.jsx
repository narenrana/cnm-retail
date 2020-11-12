import React, { useEffect } from "react";
import Grid from "@material-ui/core/Grid";

import Typography from "@material-ui/core/Typography";
import useStyles from "./style";
import _ from "lodash";

import { useDispatch, useSelector } from "react-redux";
import { getOrders } from "../redux";

export default function OrdersList() {
  const classes = useStyles();
  const dispatch = useDispatch();

  const { orders = [], isLoading = {} } = useSelector(
    (state) => state.productsStore
  );

  useEffect(() => {
    if (!isLoading) {
      dispatch(getOrders());
    }
  }, [dispatch, isLoading]);

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom className={classes.listHeader}>
        Fruits
      </Typography>
      <Grid container spacing={4}>
        {orders.map((order) => (
          <d>{order.orderId}</d>
        ))}
      </Grid>
    </React.Fragment>
  );
}
