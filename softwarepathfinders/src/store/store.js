import { configureStore } from '@reduxjs/toolkit';
import session from "./session"
import PathsReducer from "./paths"


export const store = configureStore({
  reducer: {
    session,
    paths: PathsReducer
  },
});
