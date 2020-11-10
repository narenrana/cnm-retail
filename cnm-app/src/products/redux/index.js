import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import * as _ from "lodash";
import { httpClient } from "../../core";

const initialOtpState = {
  isLoading: false,
  products: [],
  cart: {
    cartName: "",
    userId: "",
    cartItems: [],
    discountCoupon: null,
  },
};

/**
 * Thunks Actions
 */

const getProducts = createAsyncThunk(
  "get/products",
  async (request, thunkAPI) => {
    const options = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    };
    let response = (await httpClient("/products/v1/list", options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  }
);

const updateCart = createAsyncThunk(
  "cart/updateCart",
  async (request, thunkAPI) => {
    console.log({ request });
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ cart: request }),
    };
    let response = (await httpClient("/carts/v1/add", options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  }
);

const getCart = createAsyncThunk("cart/get", async (request, thunkAPI) => {
  const options = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };
  let response =
    (await httpClient("/carts/v1/list?userId=1201", options)) || [];
  response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
  return response;
});

/**
 * Reducers
 */
const Reducer = createSlice({
  name: "products/getProducts",
  initialState: initialOtpState,
  reducers: {
    addCartItem: (state, { payload }) => {
      const cartItems = state.cart.cartItems || [];
      cartItems.push({ ...payload });
      state.cart.cartItems = cartItems;
    },
    updateCartItem: (state, { payload }) => {
      const cartItems = state.cart.cartItems.map((product) => {
        if (payload.productId === product.productId) {
          product.quantity = payload.quantity;
        }
        return product;
      });
      state.cart.cartItems = cartItems.filter((item) => item.quantity > 0);
    },
  },
  extraReducers: (builder) => {
    builder.addCase(getProducts.fulfilled, (state, { payload = {} }) => {
      state.products = payload.products;
      state.isLoading = false;
    });

    builder.addCase(getProducts.rejected, (state) => {
      state.isLoading = false;
    });

    builder.addCase(updateCart.fulfilled, (state, { payload }) => {
      state.cart = { ...payload };
      state.isLoading = false;
    });

    builder.addCase(updateCart.rejected, (state) => {
      state.isLoading = false;
    });

    builder.addCase(getCart.fulfilled, (state, { payload }) => {
      state.cart = { ...payload };
      state.isLoading = false;
    });

    builder.addCase(getCart.rejected, (state) => {
      state.isLoading = false;
    });
  },
});

/*Actions export*/
const { addCartItem, updateCartItem, removeCartItem } = Reducer.actions;
export {
  getProducts,
  updateCart,
  addCartItem,
  updateCartItem,
  removeCartItem,
  getCart,
};
/*Reducer export*/
export default Reducer.reducer;
