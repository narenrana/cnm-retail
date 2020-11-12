import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import * as _ from "lodash";
import { httpClient } from "../../core";

const initialOtpState = {
  isLoading: false,
  coupons: {},
};

const listCoupons = createAsyncThunk(
  "coupons/loadCoupons",
  async (request, thunkAPI) => {
    const options = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    };
    let response =
      (await httpClient("/coupons/v1/list" + request, options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  }
);

const generateCoupons = createAsyncThunk(
  "coupons/generate",
  async (request, thunkAPI) => {
    const options = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    };
    let response =
      (await httpClient("/coupons/v1/generate" + request, options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  }
);
/**
 * Reducers
 */
const Reducer = createSlice({
  name: "coupons",
  initialState: initialOtpState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(listCoupons.fulfilled, (state, { payload }) => {
      state.discountCoupons = { ...payload };
      state.isLoading = false;
    });

    builder.addCase(listCoupons.rejected, (state, action) => {
      state.isLoading = false;
    });
    builder.addCase(generateCoupons.fulfilled, (state, { payload }) => {
      state.discountCoupons = { ...payload };
      state.isLoading = false;
    });

    builder.addCase(generateCoupons.rejected, (state, action) => {
      state.isLoading = false;
    });
  },
});

/*Actions export*/
const {} = Reducer.actions;
export {};
/*Reducer export*/
export default Reducer.reducer;
