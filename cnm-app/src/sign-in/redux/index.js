import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import * as _ from "lodash";
import { httpClient } from "../../core";

const initialOtpState = {
  isLoading: false,
  auth: {},
};

/**
 * Thunks Actions
 */

const loginUser = createAsyncThunk(
  "signIn/login",
  async (request, thunkAPI) => {
    console.log({ request });
    const options = {
      method: "GET",
      headers: {
        //"Content-Type": "application/json",
      },
      body: JSON.stringify(request),
    };
    let response =
      (await httpClient("/coupons/v1/find?coupon=naren", options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  }
);

/**
 * Reducers
 */
const Reducer = createSlice({
  name: "get/loginUser",
  initialState: initialOtpState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(loginUser.fulfilled, (state, { payload }) => {
      state.auth = { ...payload };
      state.isLoading = false;
    });

    builder.addCase(loginUser.rejected, (state, action) => {
      console.log({ state, action });
      state.isLoading = false;
    });

    builder.addCase(loginUser.pending, (state, action) => {
      console.log({ state, action });
      state.isLoading = false;
    });
  },
});

/*Actions export*/
const {} = Reducer.actions;
export { loginUser };
/*Reducer export*/
export default Reducer.reducer;
