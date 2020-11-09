import { createSlice } from "@reduxjs/toolkit";
import * as _ from "lodash";

const initialOtpState = {
  isLoading: false,
};

/**
 * Reducers
 */
const Reducer = createSlice({
  name: "header",
  initialState: initialOtpState,
  reducers: {},
  extraReducers: (builder) => {},
});

/*Actions export*/
const {} = Reducer.actions;
export {};
/*Reducer export*/
export default Reducer.reducer;
