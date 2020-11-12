import React, { useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Input from "@material-ui/core/Input";
import InputLabel from "@material-ui/core/InputLabel";
import FormControl from "@material-ui/core/FormControl";
import { useDispatch, useSelector } from "react-redux";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import clsx from "clsx";
import Button from "@material-ui/core/Button";
import { withRouter } from "react-router-dom";
import { generateCoupons, getCoupons } from "../redux";

const useStyles = makeStyles((theme) => ({
  root: {
    marginTop: "10%",
    display: "flex",
    width: "100%",
    justifyContent: "center",
  },
  margin: {
    margin: theme.spacing(1),
  },
  container: {
    display: "flex",
    maxWidth: 500,
    margin: "100px",
    background: "rgb(76 175 80 / 5%)",
    padding: "100px",
  },
}));

export default withRouter(function GenerateCoupon(props) {
  const classes = useStyles();

  const dispatch = useDispatch();

  const { discountCoupons = [], isLoading } = useSelector(
    (state) => state.couponsStore
  );
  const [values, setValues] = React.useState({
    expireTimeInSeconds: 10,
    quantity: 3,
  });

  const handleChange = (prop) => (event) => {
    setValues({ ...values, [prop]: event.target.value });
  };

  const onCouponGenerate = () => {
    dispatch(generateCoupons({ values }));
  };

  useEffect(() => {
    if (!isLoading) {
      dispatch(getCoupons());
    }
  }, [dispatch, isLoading]);

  return (
    <div className={classes.root}>
      <div className={classes.container}>
        <Grid container>
          <Grid item xs={12} md={12} xl={6}>
            <FormControl
              className={clsx(classes.margin, classes.textField)}
              fullWidth
            >
              <InputLabel htmlFor="standard-adornment-expireTimeInSeconds">
                Expire in Seconds
              </InputLabel>
              <Input
                id="standard-adornment-expireTimeInSeconds"
                type={"text"}
                value={values.expireTimeInSeconds}
                onChange={handleChange("Expire Time in Seconds")}
              />
            </FormControl>

            <FormControl
              className={clsx(classes.margin, classes.textField)}
              fullWidth
            >
              <InputLabel htmlFor="standard-adornment-password">
                Quantity
              </InputLabel>
              <Input
                id="standard-adornment-quantity"
                type={"text"}
                value={values.quantity || 5}
                onChange={handleChange("")}
              />
            </FormControl>

            <FormControl
              className={clsx(classes.margin, classes.textField)}
              fullWidth
            >
              <Button
                variant="outlined"
                color="primary"
                className={classes.checkout}
                onClick={onCouponGenerate}
              >
                Generate Coupons
              </Button>
            </FormControl>
            <FormControl
              className={clsx(classes.margin, classes.textField)}
              fullWidth
            >
              {discountCoupons.map((coupon) => (
                <Typography className={classes.countTypography}>
                  {`${coupon.description}  - ${coupon.discountCoupon} - Expiry : ${coupon.expiryDate}`}
                </Typography>
              ))}
            </FormControl>
          </Grid>
        </Grid>
      </div>
    </div>
  );
});
