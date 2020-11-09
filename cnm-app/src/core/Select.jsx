import React from 'react';
import NativeSelect from '@material-ui/core/NativeSelect';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles((theme) => ({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
    selectEmpty: {
      
    },
    root: {
        display: 'flex'
    },
    name: {
        fontSize: '0.8rem',
        padding: '5px 20px',
        fontWeight: '500px'
    }
  }));

export default function Select(props) {
    const {name,value,placeHolder, label,options=[],handleChange}= props;
    const classes = useStyles();
    const onChange=() => {
        if(handleChange) {
            handleChange()
        }
    }
    return (
        <div className={classes.root}>
         <Typography variant="caption"   className={classes.name}>
              {name}
        </Typography>
        <NativeSelect
        className={classes.selectEmpty}
        value={value}
        name="select"
        onChange={onChange}
        inputProps={{ 'aria-label': label}}
      >

        <option value="" disabled>
          {placeHolder}
        </option>
        {
            options.map((option, index)=> <option value={option.value} key={option.value+"-"+index}>{option.label}</option>)
        }
      </NativeSelect>
      </div>
    )
}