import React from "react";
import CssBaseline from "@material-ui/core/CssBaseline";
import Grid from "@material-ui/core/Grid";
import Container from "@material-ui/core/Container";
import Header from "./header/Header";
import Footer from "./footer/Footer";
import { ThemeProvider } from "@material-ui/core/styles";
import SideNavBar from "./sidebar/SideNavBar";
import ProductList from "./products/components/product-home/ProductHome";
import { Checkout } from "./checkout";
import theme from "./theme";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import { Provider } from "react-redux";
import configureStore from "./configure-store";
const store = configureStore();

const sections = [
  { title: "Home", url: "/home" },
  { title: "Orders", url: "/orders" },
  { title: "Paymemnts", url: "/paymemnts" },
];

export default function Cart() {
  return (
    <Provider store={store}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Container maxWidth="lg">
          <Header title="Cart" sections={sections} />
          <Router>
            <Switch>
              <Route path="/payment">
                <div>Payment</div>
              </Route>
              <Route path="/checkout">
                <main>
                  <Checkout />
                </main>
              </Route>
              <Route path="/">
                <main>
                  <Grid container>
                    <Grid item xs={12} md={3} xl={3}>
                      <SideNavBar />
                    </Grid>
                    <Grid item xs={12} md={9} xl={9}>
                      <ProductList />
                    </Grid>
                  </Grid>
                </main>
              </Route>
            </Switch>
          </Router>
        </Container>
        <Footer
          title="Footer"
          description="Something here to give the footer a purpose!"
        />
      </ThemeProvider>
    </Provider>
  );
}
