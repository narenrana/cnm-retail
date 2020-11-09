import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
 

const useStyles = makeStyles((theme) => ({
  root: {
    '& > *': {
      marginTop: theme.spacing(2),
    },
  },
}));

export default function Checkout(props) {
  const classes = useStyles();
  const { count } = props;

  return (
    <div className={classes.root}>
       
    </div>
  );
}