import React, { useEffect } from "react";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Typography from "@material-ui/core/Typography";

import useStyles from "../../products/style";

export default function ProductCart(props) {
  const { product, cartQuantity } = props;
  const classes = useStyles();

  return (
    <Card className={classes.root}>
      <CardMedia
        component="img"
        alt="Contemplative Reptile"
        height="140"
        image={product.imageUrl}
        title={product.productTitle}
      />
      <CardContent className={classes.cardTitleContainer}>
        <Typography variant="subtitle2" className={classes.title}>
          {product.productTitle}
        </Typography>
      </CardContent>
      <CardContent className={classes.cardPriceContainer}>
        <Typography variant="body1" className={classes.price}>
          {`${product.productPrice} ${product.baseCurrency}`}
        </Typography>
      </CardContent>
      <CardContent className={classes.cardPriceContainer}>
        <Typography variant="body1" className={classes.price}>
          {`Qantity - ${cartQuantity}`}
        </Typography>
      </CardContent>
    </Card>
  );
}
