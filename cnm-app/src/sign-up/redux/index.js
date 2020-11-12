import { createSlice } from "@reduxjs/toolkit";
import SignIn from "./components/SignIn";
export { SignIn };

const initialOtpState = {
  isLoading: false,
  auth: {},
};

/**
 * Thunks Actions
 */

/**
 * Reducers
 */
const Reducer = createSlice({
  name: "get/loginUser",
  initialState: initialOtpState,
  reducers: {},
  extraReducers: (builder) => {},
});

/*Actions export*/
const {} = Reducer.actions;
export {};
/*Reducer export*/
export default Reducer.reducer;
