import { makeStyles } from "@material-ui/core/styles";
export default makeStyles({
  root: {
    maxWidth: 345,
  },
  cardTitleContainer: {
    display: "flex",
    justifyContent: "center",
    padding: "5px",
  },
  title: {
    fontSize: "1em",
  },
  extendedIcon: {
    fontSize: "1.5em",
    margin: "0 15px",
  },
  extendedIconText: {
    fontSize: "1em",
    padding: "0",
  },
  storeButton: {
    padding: "0 8px",
    width: "100% !important",
    borderRadius: "8px",
    boxShadow: "none",
    background: "#e0e0e0",
    fontSize: "0.8em",
  },
  cardActions: {
    alignItems: "center",
    justifyContent: "center",
  },
  countShoppingCartIcon: {
    padding: "0",
    fontSize: "1rem",
  },
  countActionContainer: {
    display: "flex",
    background: "#4caf50",
    borderRadius: "26px",
    width: "100%",
    justifyContent: "space-around",
  },
  countActionButton: {
    padding: "0 10px",
  },
  countActionButtonIcon: {
    color: "#fff",
    fontSize: "1.1rem",
  },
  countTypography: {
    color: "#fff",
    padding: "5px",
  },
  counterCardFabButton: {
    boxShadow: "none",
  },
  countActionsRoot: {
    justifyContent: "center",
  },
  listHeader: {
    paddingTop: "0",
  },
  filterSection: {
    display: "flex",
    justifyContent: "space-between",
    margin: "5px 0",
  },
  cardPriceContainer: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    padding: 0,
  },
  price: {
    color: "#707070",
    fontSize: "0.8rem",
    fontWeight: 500,
  },
});
