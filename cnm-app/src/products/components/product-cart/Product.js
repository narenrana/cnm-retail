import React, { useEffect } from "react";
import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Typography from "@material-ui/core/Typography";
import AddShoppingCart from "@material-ui/icons/AddShoppingCart";
import Fab from "@material-ui/core/Fab";
import IconButton from "@material-ui/core/IconButton";
import Add from "@material-ui/icons/Add";
import Remove from "@material-ui/icons/Remove";
import useStyles from "../../style";

export default function ProductCart(props) {
  const {
    product,
    onAddToCart,
    onIncreaseQuantity,
    onReduceQuantity,
    isExpanded,
    cartQuantity,
  } = props;
  const classes = useStyles();
  const [expanded, setExpanded] = React.useState(isExpanded);
  const [count, setCount] = React.useState(cartQuantity);

  const addToCart = (e) => {
    setExpanded(true);
    onAddToCart(product);
  };

  const onAdd = () => {
    const newCount = count + 1;
    setCount(newCount);
    onIncreaseQuantity({ ...product, quantity: newCount });
  };

  const onRemove = () => {
    const newCount = count > 0 ? count - 1 : 0;
    setCount(newCount);
    onReduceQuantity({ ...product, quantity: newCount });
  };

  useEffect(() => {
    setExpanded(isExpanded);
    setCount(cartQuantity);
    return () => {};
  }, [isExpanded, cartQuantity]);

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
      {!expanded && (
        <CardActions disableSpacing className={classes.cardActions}>
          <Fab
            variant="extended"
            size="small"
            color="secondary"
            aria-label="add"
            className={classes.storeButton}
            onClick={addToCart}
          >
            <AddShoppingCart className={classes.extendedIcon} />
            <div className={classes.extendedIconText}>Add to cart</div>
          </Fab>
        </CardActions>
      )}
      {expanded && (
        <CardActions disableSpacing className={classes.countActionsRoot}>
          <div className={classes.countActionContainer}>
            <Fab size="small" className={classes.counterCardFabButton}>
              <AddShoppingCart
                className={classes.countShoppingCartIcon}
                color="primary"
              />
            </Fab>
            <div className={classes.countActionContainer}>
              <IconButton
                onClick={onRemove}
                className={classes.countActionButton}
              >
                <Remove
                  color="secondary"
                  className={classes.countActionButtonIcon}
                />
              </IconButton>
              <Typography className={classes.countTypography}>
                {count}
              </Typography>
              <IconButton onClick={onAdd} className={classes.countActionButton}>
                <Add
                  color="secondary"
                  className={classes.countActionButtonIcon}
                />
              </IconButton>
            </div>
          </div>
        </CardActions>
      )}
    </Card>
  );
}
