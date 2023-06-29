import { configureStore } from '@reduxjs/toolkit';
import session from "./session"
import PathsReducer from "./paths"
import CohortsReducer from './cohorts';
import PostsReducer from './posts';


export const store = configureStore({
  reducer: {
    session,
    paths: PathsReducer,
    cohorts: CohortsReducer,
    posts: PostsReducer,
  },
});
