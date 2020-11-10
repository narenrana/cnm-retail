import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import * as _ from "lodash";
import { httpClient, getAuthState, storeAuth } from "../../core";

const initialOtpState = {
  isLoading: false,
  cart: [],
  auth: getAuthState() || {},
};

/**
 * Thunks Actions
 */
const login = createAsyncThunk("auth/login", async (request, thunkAPI) => {
  const options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(request),
  };
  let response = (await httpClient("/auth/v1/login", options)) || [];

  response = _.isEmpty(response) ? initialOtpState.auth : response;
  storeAuth({ ...response, isLogin: true });
  return response;
});

const logout = createAsyncThunk("auth/logout", async (request, thunkAPI) => {
  storeAuth({ isLogin: false });
  const options = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(request),
  };
  let response = (await httpClient("/auth/v1/logout", options)) || [];
  response = _.isEmpty(response) ? initialOtpState.auth : response;
  return response;
});

const refreshToken = createAsyncThunk(
  "auth/refreshToken",
  async (request, thunkAPI) => {
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(request),
    };
    let response = (await httpClient("/auth/v1/login", options)) || [];
    response = _.isEmpty(response) ? initialOtpState.auth : response;
    storeAuth({ ...response, isLogin: true });
    return response;
  }
);

/**
 * Reducers
 */
const Reducer = createSlice({
  name: "common",
  initialState: initialOtpState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(login.fulfilled, (state, { payload }) => {
      if (payload.error) {
        state.error = payload.error;
        state.auth.isLogin = false;
        return;
      } else {
        state.auth = { ...payload, isLogin: true };
        state.isLoading = false;
      }
    });

    builder.addCase(login.rejected, (state, action) => {
      state.isLoading = false;
    });
    builder.addCase(logout.fulfilled, (state, { payload }) => {
      state.auth = { ...payload };
      state.isLoading = false;
    });

    builder.addCase(logout.rejected, (state, action) => {
      state.isLoading = false;
    });
    builder.addCase(refreshToken.fulfilled, (state, { payload }) => {
      state.auth = { ...payload };
      state.isLoading = false;
    });

    builder.addCase(refreshToken.rejected, (state, action) => {
      state.isLoading = false;
    });
  },
});

/*Actions export*/
const {} = Reducer.actions;
export { login };
/*Reducer export*/
export default Reducer.reducer;
