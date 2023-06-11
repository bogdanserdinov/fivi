import { createSlice } from '@reduxjs/toolkit';
import type { PayloadAction } from '@reduxjs/toolkit';

import { createPost, deletePost, getPost, getPostsProfile, likeAndDislikePost, sendComment } from '../actions/posts';
import { Post } from '@/post';

/** Exposes channels state */
class PostsState {
    /** class implementation */
    constructor(
        public currentPost: Post = new Post(),
        public userProfilePosts: Post[] = [],
        public homePosts: Post[] = [],
        public postPhotos: string[] = [],
        public currentUserProfileId: string = '',
    ) { }
}

const initialState: PostsState = {
    currentPost: new Post(),
    userProfilePosts: [],
    homePosts: [],
    postPhotos: [],
    currentUserProfileId: '',
};

export const postsSlice = createSlice({
    name: 'postsReducer',
    initialState,
    reducers: {
        setCurrentPost: (state, action: PayloadAction<Post>) => {
            state.currentPost = action.payload;
        },
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

        setCurrentId: (state, action: PayloadAction<string>) => {
            state.currentUserProfileId = action.payload;
        },

        deletePostPhotos: (state) => {
            state.postPhotos = [];
        },
    },
    extraReducers: (builder) => {
        builder.addCase(createPost.fulfilled, (state, action) => {
            const homePosts = state.homePosts.concat(action.payload);

            const userProfilePosts = action.payload.creatorId === state.currentUserProfileId ?
                state.userProfilePosts.concat(action.payload)
                :
                state.userProfilePosts;

            return {
                ...state,
                userProfilePosts: userProfilePosts,
                homePosts: homePosts,
            };
        });

        builder.addCase(getPost.fulfilled, (state, action) => {
            state.currentPost = action.payload;
        });

        builder.addCase(deletePost.fulfilled, (state, action) => {
            const userProfilePosts = state.userProfilePosts.filter((post) => post.postId !== action.payload);
            const homePosts = state.homePosts.filter((post) => post.postId !== action.payload);

            return {
                ...state,
                userProfilePosts: userProfilePosts,
                homePosts: homePosts,
            };
        });

        builder.addCase(getPostsProfile.fulfilled, (state, action) => {
            state.userProfilePosts = action.payload;
        });

        builder.addCase(sendComment.fulfilled, (state, action) => {
            const userProfilePosts = state.userProfilePosts.map((post) => {
                if (post.postId === action.payload.postId) {
                    post.comments.push(action.payload);
                    post.numOfComments += 1;
                }

                return post;
            }

            );
            const homePosts = state.homePosts.map((post) => {
                if (post.postId === action.payload.postId) {
                    post.comments.push(action.payload);
                    post.numOfComments += 1;
                }

                return post;
            }

            );
            const comments = state.currentPost.comments.concat(action.payload);
            const numOfComments = state.currentPost.numOfComments = +1;

            return {
                ...state,
                userProfilePosts: userProfilePosts,
                homePosts: homePosts,
                currentPost: {
                    ...state.currentPost,
                    comments: comments,
                    numOfComments: numOfComments,
                },
            };
        });

        builder.addCase(likeAndDislikePost.fulfilled, (state, action) => {
            state.userProfilePosts = state.userProfilePosts.map((post) => {
                if (post.postId === action.payload.post_id) {
                    action.payload.isLiked ?
                        post.numOfLikes += 1
                        :
                        post.numOfLikes -= 1;
                };

                return post;
            });
            state.homePosts = state.homePosts.map((post) => {
                if (post.postId === action.payload.post_id) {
                    action.payload.isLiked ?
                        post.numOfLikes += 1
                        :
                        post.numOfLikes -= 1;
                };

                return post;
            });
        });
    },
});

// Action creators are generated for each case reducer function
export const { setPostsHomePage, setCurrentId, setPostPhotos, deletePostPhoto, addPostPhotos, deletePostPhotos, setCurrentPost } = postsSlice.actions;

export default postsSlice.reducer;
