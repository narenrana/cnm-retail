import React, { useEffect } from "react";
import Grid from "@material-ui/core/Grid";

import Pagination from "../../../core/Pagination";
import Typography from "@material-ui/core/Typography";
import useStyles from "../../style";
import _ from "lodash";

import { useDispatch, useSelector } from "react-redux";
import { processOrder } from "../../redux";

export default function OrdersList() {
  const classes = useStyles();
  const dispatch = useDispatch();

  useEffect(() => {}, [dispatch, isLoading]);

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom className={classes.listHeader}>
        Fruits
      </Typography>
      <Grid container spacing={4}></Grid>
      <Pagination count={10} />
    </React.Fragment>
  );
}
