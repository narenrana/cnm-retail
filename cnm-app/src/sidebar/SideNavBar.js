import React from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Collapse from '@material-ui/core/Collapse';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import StarBorder from '@material-ui/icons/StarBorder';
import Divider from '@material-ui/core/Divider';
import options from './data';
import useStyles from './style';

export default function SideBarList() {
  const classes = useStyles();
  const [open, setOpen] = React.useState(false);

  const handleClick = () => {
    setOpen(!open);
  };

  if (options) {
    return (
      <div className={classes.root}>
        <List>
          {options.map((option, index) =>
            <React.Fragment key={'list-button' + option.label + index}>
              <ListItem button onClick={handleClick} >
                <ListItemText primary={option.label} />
                {option.children && (open ? <ExpandLess /> : <ExpandMore />)}
              </ListItem>

              {option.children && <Collapse in={open} timeout="auto" unmountOnExit>
                <List component="div" disablePadding>
                  {option.children.map(child => <ListItem button className={classes.nested} key={'list-item' + (child.label) + index}>
                    <ListItemIcon>
                      <StarBorder />
                    </ListItemIcon>
                    <ListItemText primary={child.label} />
                  </ListItem>)}
                </List>
              </Collapse>}
              <Divider />
            </React.Fragment>
          )}
        </List>
      </div>
    )
  }
}
