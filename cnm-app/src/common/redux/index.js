import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import * as _ from "lodash";
import { httpClient } from "../../core";

const initialOtpState = {
  isLoading: false,
  cart: [],
  auth: {},
};

/**
 * Thunks Actions
 */
const login = (data) =>
  createAsyncThunk("auth/login", async (request, thunkAPI) => {
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      data,
    };
    let response = (await httpClient("/auth/login", options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  });

const logout = (data) =>
  createAsyncThunk("auth/logout", async (request, thunkAPI) => {
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      data,
    };
    let response = (await httpClient("/auth/login", options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  });

const refreshToken = (data) =>
  createAsyncThunk("auth/refreshToken", async (request, thunkAPI) => {
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      data,
    };
    let response = (await httpClient("/auth/login", options)) || [];
    response = _.isEmpty(response) ? initialOtpState.defaultTabList : response;
    return response;
  });

/**
 * Reducers
 */
const Reducer = createSlice({
  name: "footer",
  initialState: initialOtpState,
  reducers: {},
  extraReducers: (builder) => {},
});

/*Actions export*/
const {} = Reducer.actions;
export { login };
/*Reducer export*/
export default Reducer.reducer;
