import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import * as _ from "lodash";
import { httpClient } from "../../core";

const initialOtpState = {
  isLoading: false,
  orders: [],
};

/**
 * Thunks Actions
 */

const getOrders = createAsyncThunk("get/orders", async (request, thunkAPI) => {
  const options = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };
  let response = (await httpClient("/orders/v1/list", options)) || [];
  response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
  return response;
});

const placeOrder = createAsyncThunk(
  "orders/place",
  async (request, thunkAPI) => {
    console.log({ request });
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(request),
    };
    let response = (await httpClient("/orders/v1/placeOrder", options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  }
);

/**
 * Reducers
 */
const Reducer = createSlice({
  name: "orders/get",
  initialState: initialOtpState,
  reducers: {
    hideSucessMessage: (state) => {
      state.showSucessMessage = false;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(getOrders.fulfilled, (state, { payload = {} }) => {
      let orders = [];
      //orders = orders.concat(state.orders);
      orders = orders.concat(payload.orders);
      state.orders = orders;
      state.showSucessMessage = true;
      state.isLoading = false;
    });

    builder.addCase(getOrders.rejected, (state) => {
      state.isLoading = false;
    });

    builder.addCase(placeOrder.fulfilled, (state, { payload }) => {
      const orders = { ...payload };
      state.orders.push(orders);
      state.showSucessMessage = true;
      state.isLoading = false;
    });

    builder.addCase(placeOrder.rejected, (state) => {
      state.isLoading = false;
    });
  },
});

/*Actions export*/
const {} = Reducer.actions;
export { placeOrder, getOrders };
/*Reducer export*/
export default Reducer.reducer;
