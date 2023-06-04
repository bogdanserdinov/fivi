import { configureStore } from '@reduxjs/toolkit';

import { handleErrorMiddleware } from './middleaware';

import userSlice from '@/app/store/reducers/users';
import postsSlice from '@/app/store/reducers/posts';

const reducer = {
    usersReducer: userSlice,
    postsReducer: postsSlice,
};

export const store = configureStore({
    reducer,
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware({ serializableCheck: false }).concat(handleErrorMiddleware),
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
