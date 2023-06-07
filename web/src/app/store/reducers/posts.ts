import { createSlice } from '@reduxjs/toolkit';
import type { PayloadAction } from '@reduxjs/toolkit';

import { createPost, deletePost, getPost, getPostsProfile, likeAndDislikePost, sendComment, updatePost } from '../actions/posts';
import { Post } from '@/post';

/** Exposes channels state */
class PostsState {
    /** class implementation */
    constructor(
        public currentPost: Post = new Post(),
        public userProfilePosts: Post[] = [],
        public homePosts: Post[] = [],
        public postPhotos: string[] = []
    ) { }
}

const initialState: PostsState = {
    currentPost: new Post(),
    userProfilePosts: [],
    homePosts: [],
    postPhotos: [],
};

export const postsSlice = createSlice({
    name: 'postsReducer',
    initialState,
    reducers: {
        setPostsHomePage: (state, action: PayloadAction<Post[]>) => {
            state.homePosts = action.payload;
        },
        setPostPhotos: (state, action: PayloadAction<string[] | []>) => {
            state.postPhotos =
                action.payload;
        },

        deletePostPhoto: (state, action: PayloadAction<string>) => {
            state.postPhotos =
                state.postPhotos.filter((photo) => photo !== action.payload);
        },
        addPostPhotos: (state, action: PayloadAction<string[]>) => {
            state.postPhotos =
                state.postPhotos.concat(action.payload);
        },
        deletePostPhotos: (state) => {
            state.postPhotos = [];
        },
    },
    extraReducers: (builder) => {
        builder.addCase(createPost.fulfilled, (state, action) => {
            state.homePosts = state.userProfilePosts.concat(action.payload);
            state.userProfilePosts = state.homePosts.concat(action.payload);
        });

        builder.addCase(getPost.fulfilled, (state, action) => {
            state.currentPost = action.payload;
        });

        builder.addCase(deletePost.fulfilled, (state, action) => {
            state.userProfilePosts = state.userProfilePosts.filter((post) => post.postId !== action.payload);
        });

        builder.addCase(getPostsProfile.fulfilled, (state, action) => {
            state.userProfilePosts = action.payload;
        });

        builder.addCase(sendComment.fulfilled, (state, action) => {
            state.userProfilePosts = state.userProfilePosts.map((post) => {
                if (post.postId === action.payload.postId) {
                    post.comments.concat(action.payload);
                }

                return post;
            }

            );
            state.homePosts = state.homePosts.map((post) => {
                if (post.postId === action.payload.postId) {
                    post.comments.concat(action.payload);
                }

                return post;
            }
            );
        });

        builder.addCase(likeAndDislikePost.fulfilled, (state, action) => {
            state.userProfilePosts = state.userProfilePosts.map((post) => {
                if (post.postId === action.payload.post_id) {
                    action.payload.isLiked ?
                        post.num_of_likes += 1
                        :
                        post.num_of_likes -= 1;
                };

                return post;
            });
            state.homePosts = state.homePosts.map((post) => {
                if (post.postId === action.payload.post_id) {
                    action.payload.isLiked ?
                        post.num_of_likes += 1
                        :
                        post.num_of_likes -= 1;
                };

                return post;
            });
        });
    },
});

// Action creators are generated for each case reducer function
export const { setPostsHomePage, setPostPhotos, deletePostPhoto, addPostPhotos, deletePostPhotos } = postsSlice.actions;

export default postsSlice.reducer;
