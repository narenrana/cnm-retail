import React from "react";
import PropTypes from "prop-types";
import Container from "@material-ui/core/Container";
import Typography from "@material-ui/core/Typography";
import { Grid } from "@material-ui/core";
import MenuItem from "@material-ui/core/MenuItem";
import MenuList from "@material-ui/core/MenuList";
import Copyright from "./Copyright";
import { footerMenu, appStoreLinks, footerSocialMediaLinks } from "./data";
import useStyles from "./style";
import Subscribe from "./Subscribe";

export default function Footer(props) {
  const classes = useStyles();

  return (
    <footer className={classes.footer}>
      <Container maxWidth="lg">
        <Subscribe />
        <Grid container spacing={2}>
          <Grid item xs={12} md={10} xl={10}>
            <Grid container spacing={0}>
              {footerMenu.map((footer, index) => (
                <Grid
                  item
                  xs={12}
                  md={2}
                  xl={2}
                  key={"foot-menu-link-" + footer.title + index}
                >
                  <Typography variant="subtitle1">{footer.title}</Typography>
                  <MenuList>
                    {footer.items.map((menu) => (
                      <MenuItem
                        key={"foot-items-link-" + menu.label + "-" + index}
                      >
                        {menu.label}
                      </MenuItem>
                    ))}
                  </MenuList>
                </Grid>
              ))}
            </Grid>
          </Grid>
          <Grid item xs={12} md={2} xl={2}>
            <MenuList>
              {appStoreLinks.map((menu, index) => (
                <MenuItem key={"app-store-link-" + index}>
                  <img alt={menu.image} src={menu.image} />
                </MenuItem>
              ))}
            </MenuList>
            <Grid container className={classes.socialLinks}>
              {footerSocialMediaLinks.map((links, index) => (
                <Grid item xs={3} key={"social-link-" + index}></Grid>
              ))}
            </Grid>
          </Grid>
        </Grid>

        <Copyright />
      </Container>
    </footer>
  );
}

Footer.propTypes = {
  description: PropTypes.string,
  title: PropTypes.string,
};
