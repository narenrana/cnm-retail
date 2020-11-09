import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit';
import * as _ from 'lodash';
import { httpClient } from '../../core'
 

const initialOtpState = {
  isLoading: false,
};

 


/**
 * Reducers
 */
const Reducer = createSlice({
  name: 'footer',
  initialState: initialOtpState,
  reducers: {
    
 
  },
  extraReducers: (builder) => {
 
  },
});

/*Actions export*/
const {
  
} = Reducer.actions;
export {
  
};
/*Reducer export*/
export default Reducer.reducer;
