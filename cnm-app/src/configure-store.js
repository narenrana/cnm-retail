import { configureStore } from "@reduxjs/toolkit";
import thunkMiddleware from "redux-thunk";
import { createLogger } from "redux-logger";
import rootReducer from "./reducers.js";

export default function configureAppStore() {
  const logger = createLogger();
  const store = configureStore({
    reducer: rootReducer,
    middleware: [logger, thunkMiddleware],
  });

  if (process.env.NODE_ENV !== "production" && module.hot) {
    module.hot.accept("./reducers", () => store.replaceReducer(rootReducer));
  }

  return store;
}
