import React from "react";

import Card from "@material-ui/core/Card";
import CardHeader from "@material-ui/core/CardHeader";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import ShoppingBasket from "@material-ui/icons/ShoppingBasket";
import MoreVertIcon from "@material-ui/icons/MoreVert";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import { useDispatch, useSelector } from "react-redux";
import { get } from "lodash";
import useStyles from "../style";
import { loadCoupon } from "../../redux";

export default function CartDetails(props) {
  const { cart, addCouponToCart } = props;
  const { discountCoupons } = useSelector((state) => state.checkoutStore);
  const dispatch = useDispatch();
  const classes = useStyles();
  const [couponValue, setCouponValue] = React.useState("");

  const onCouponApply = (e) => {
    addCouponToCart(couponValue);
  };

  const onCouponRemove = (e) => {
    addCouponToCart("");
  };

  const onCouponEnter = (e) => {
    console.log({ e: e.target.value });
    setCouponValue(e.target.value);
    dispatch(loadCoupon(e.target.value));
  };

  const getHelperMessage = () =>
    couponValue && !discountCoupons.valid && "Invalid Coupon";

  const isInValidCoupon = () => {
    return couponValue && !discountCoupons.valid;
  };

  return (
    <Card className={classes.root}>
      <CardHeader
        action={
          <IconButton aria-label="settings">
            <MoreVertIcon />
          </IconButton>
        }
        title={`$${cart.totalAmount || 0}`}
        subheader={`Total Savings - $${cart.totalDiscount || 0}`}
      />

      <CardContent>
        <Typography variant="h6" component="h2" className={classes.heading}>
          Offerers Applied
        </Typography>
        {cart.offersDiscount &&
          (cart.appliedOffers || []).map((offers) => (
            <Typography variant="body2" color="textSecondary" component="p">
              {get(offers, "description")}
            </Typography>
          ))}

        {cart.offersDiscount && (
          <Typography
            variant="body2"
            color="primary"
            component="p"
            className={classes.discount}
          >
            Offers Amount : ${cart.offersDiscount || 0}
          </Typography>
        )}
        <Typography variant="h6" component="h2" className={classes.heading}>
          Coupon Applied {cart.discountCoupon && ` - ${cart.discountCoupon}`}
        </Typography>

        <Typography variant="body2" color="primary" component="p">
          {get(cart, "discountCouponDetails.description")}
        </Typography>

        {cart.couponDiscount && (
          <Typography
            variant="body2"
            color="primary"
            component="p"
            className={classes.discount}
          >
            Coupon Discount : ${cart.couponDiscount}
          </Typography>
        )}

        <Typography variant="subtitle2" color="textSecondary" component="p">
          <TextField
            id="outlined-search"
            label="Enter Coupon"
            type="search"
            variant="outlined"
            size="small"
            className={classes.couponInputBox}
            onChange={onCouponEnter}
            error={isInValidCoupon()}
            helperText={getHelperMessage()}
            value={couponValue}
          />
        </Typography>
        <Button
          size="small"
          variant="outlined"
          color="inherit"
          className={classes.margin}
          onClick={onCouponApply}
          disabled={isInValidCoupon() || !couponValue}
        >
          Apply Coupon
        </Button>
        <Button
          size="small"
          variant="outlined"
          className={classes.margin}
          onClick={onCouponRemove}
          disabled={discountCoupons.valid && cart.discountCoupon}
        >
          Clear Coupon
        </Button>
        <Typography
          variant="h6"
          component="h2"
          className={classes.heading}
        ></Typography>
        <Typography variant="h6" component="h2" className={classes.heading}>
          Total Paybel - ${cart.totalAmount || 0 - cart.totalDiscount || 0}
        </Typography>
      </CardContent>

      <CardActions disableSpacing>
        <Button
          variant="contained"
          color="primary"
          className={classes.checkout}
          startIcon={<ShoppingBasket color="secondary" />}
        >
          checkout
        </Button>
      </CardActions>
    </Card>
  );
}
