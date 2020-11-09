import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Pagination from "@material-ui/lab/Pagination";

const useStyles = makeStyles((theme) => ({
  root: {
    "& > *": {
      marginTop: theme.spacing(2),
    },
  },
}));

export default function PaginationButtons(props) {
  const classes = useStyles();
  const { count } = props;

  return (
    <div className={classes.root}>
      <Pagination count={count} showFirstButton showLastButton />
    </div>
  );
}
