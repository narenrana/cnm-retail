import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import * as _ from "lodash";
import { httpClient } from "../../core";

const initialOtpState = {
  isLoading: false,
  isValid: null,
  discountCoupons: {},
};

/**
 * Thunks Actions
 */

const loadCoupon = createAsyncThunk(
  "checkout/loadCoupons",
  async (request, thunkAPI) => {
    const options = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    };
    let response =
      (await httpClient("/coupons/v1/find?coupon=" + request, options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  }
);

/**
 * Reducers
 */
const Reducer = createSlice({
  name: "checkout",
  initialState: initialOtpState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(loadCoupon.fulfilled, (state, { payload }) => {
      state.discountCoupons = { ...payload };
      state.isLoading = false;
    });

    builder.addCase(loadCoupon.rejected, (state, action) => {
      state.isLoading = false;
    });
  },
});

/*Actions export*/
const {} = Reducer.actions;
export { loadCoupon };
/*Reducer export*/
export default Reducer.reducer;
