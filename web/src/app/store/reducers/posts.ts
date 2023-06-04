import { createSlice } from '@reduxjs/toolkit';
import type { PayloadAction } from '@reduxjs/toolkit';

import { deletePost, getPost, getPostsProfile, updatePost } from '../actions/posts';
import { Post } from '@/post';

/** Exposes channels state */
class PostsState {
    /** class implementation */
    constructor(
        public currentPost: Post = new Post(),
        public userProfilePosts: Post[] = [],
        public homePosts: Post[] = []
    ) { }
}

const initialState: PostsState = {
    currentPost: new Post(),
    userProfilePosts: [],
    homePosts: [],
};

export const postsSlice = createSlice({
    name: 'postsReducer',
    initialState,
    reducers: {
        setPostsHomePage: (state, action: PayloadAction<Post[]>) => {
            state.homePosts = action.payload;
        },
    },
    extraReducers: (builder) => {
        builder.addCase(getPost.fulfilled, (state, action) => {
            state.currentPost = action.payload;
        });

        builder.addCase(deletePost.fulfilled, (state, action) => {
            state.userProfilePosts = state.userProfilePosts.filter((post) => post.postId !== action.payload);
        });

        builder.addCase(getPostsProfile.fulfilled, (state, action) => {
            state.userProfilePosts = action.payload;
        });
    },
});

// Action creators are generated for each case reducer function
export const { setPostsHomePage } = postsSlice.actions;

export default postsSlice.reducer;
