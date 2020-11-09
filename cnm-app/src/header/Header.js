import React from "react";
import PropTypes from "prop-types";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import Link from "@material-ui/core/Link";
import AddShoppingCart from "@material-ui/icons/AddShoppingCart";
import UserIcon from "@material-ui/icons/PermIdentity";

import Badge from "@material-ui/core/Badge";
import InputBase from "@material-ui/core/InputBase";
import SearchIcon from "@material-ui/icons/Search";
import useStyles from "./style";
import { useSelector } from "react-redux";
import * as _ from "lodash";

export default function Header(props) {
  const classes = useStyles();
  const { sections } = props;

  const { cart } = useSelector((state) => state.productsStore);
  const { cartItems = [] } = cart;
  const itemsCount = _.reduce(
    cartItems,
    (reduced, item) => {
      return reduced + item.quantity;
    },
    0
  );

  return (
    <React.Fragment>
      <Toolbar className={classes.toolbar}>
        <div className={classes.rightMenuPanel}>
          <a href="/checkout">
            <IconButton>
              <Badge badgeContent={itemsCount} color="primary">
                <AddShoppingCart />
              </Badge>
            </IconButton>
            <IconButton>
              <UserIcon />
            </IconButton>
          </a>
        </div>
      </Toolbar>
      <Toolbar
        component="nav"
        variant="dense"
        className={classes.toolbarSecondary}
      >
        {sections.map((section) => (
          <Link
            key={section.title}
            variant="body2"
            href={section.url}
            className={classes.toolbarLink}
            color="primary"
          >
            {section.title}
          </Link>
        ))}
        <div className={classes.search}>
          <div className={classes.searchIcon}>
            <SearchIcon />
          </div>
          <InputBase
            placeholder="Search…"
            classes={{
              root: classes.inputRoot,
              input: classes.inputInput,
            }}
            inputProps={{ "aria-label": "search" }}
          />
        </div>
      </Toolbar>
    </React.Fragment>
  );
}

Header.propTypes = {
  sections: PropTypes.array,
  title: PropTypes.string,
};
