import { makeStyles } from "@material-ui/core/styles";
import { red } from "@material-ui/core/colors";
export default makeStyles((theme) => ({
  root: {
    maxWidth: 345,
    float: "right",
  },
  media: {
    height: 0,
  },
  expand: {
    transform: "rotate(0deg)",
    marginLeft: "auto",
    transition: theme.transitions.create("transform", {
      duration: theme.transitions.duration.shortest,
    }),
  },
  expandOpen: {
    transform: "rotate(180deg)",
  },
  avatar: {
    backgroundColor: red[500],
  },
  heading: {
    margin: "10px 0",
    padding: "10px 0",
    borderBottom: "1px solid #ccc",
    textTransform: " uppercase",
    //borderTop: "1px solid #ccc",
  },
  deleteIcon: {},
  couponInputBox: {
    margin: "20px 0px",
  },
  margin: {
    margin: "0 5px",
  },
  checkout: {
    width: "100%",
    color: "#fff",
  },
  discount: {
    margin: "5px 0",
  },
}));
