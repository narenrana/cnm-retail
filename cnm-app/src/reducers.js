/*
 * Module Description :  Configure Redux App Store
 *
 * Revision History
 *
 * Date            Author             SR No          Description of Change
 * ----------      ----------         ------         ----------------------
 * 9-June-2020     Narender Kumar     Release1       Initial commit
 *
 * Copyright (c) aposailor.com.
 * All rights reserved.
 *
 * This software is the confidential and proprietary information of
 * aposailor.com ("Confidential Information").You shall not
 * disclose such Confidential Information and shall use it only in accordance
 * with the terms of the license agreement you entered into with aposailor.com
 */

import { combineReducers } from "redux";
import commonStore from "./common/redux";
import checkoutStore from "./checkout/redux";
import footerStore from "./footer/redux";
import headerStore from "./header/redux";
import paymentStore from "./header/redux";
import productsStore from "./products/redux";
import sidebarStore from "./sidebar/redux";
import signInStore from "./sign-in/redux";
//import signUpStore from "./sign-up/redux";

const store = combineReducers({
  commonStore,
  checkoutStore,
  footerStore,
  headerStore,
  paymentStore,
  productsStore,
  sidebarStore,
  signInStore,
  //signUpStore,
});

export default store;
