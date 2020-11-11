import React from "react";
import Typography from "@material-ui/core/Typography";
import InputBase from "@material-ui/core/InputBase";
import Fab from "@material-ui/core/Fab";
import useStyles from "./style";
import { Grid } from "@material-ui/core";
import useMediaQuery from "@material-ui/core/useMediaQuery";

function SubscribeGrid() {
  const classes = useStyles();
  return (
    <div>
      <Grid container spacing={0}>
        <Grid item xs={12} md={12} xl={12}>
          <Typography className={classes.subscribeTypehead}>
            Subscribe to our news latter and receive exclusive offer every week
          </Typography>
        </Grid>
        <Grid item xs={12} md={8} xl={8}>
          <InputBase
            className={classes.subscribeInput}
            placeholder="Enter your email"
          />
        </Grid>
        <Grid item xs={12} md={4} xl={4}>
          <Fab
            variant="extended"
            size="small"
            color="primary"
            aria-label="subscribe"
            className={classes.subscribeButton}
          >
            <div className={classes.subscribeButtonText}>Subscribe</div>
          </Fab>
        </Grid>
      </Grid>
    </div>
  );
}

export default function Subscribe() {
  const classes = useStyles();
  const matches = useMediaQuery("(min-width:600px)");

  return matches ? (
    <div className={classes.subscribeRoot}>
      <div className={classes.subscribeMenu}>
        <SubscribeGrid />
      </div>
    </div>
  ) : (
    <div>
      <SubscribeGrid />
    </div>
  );
}
