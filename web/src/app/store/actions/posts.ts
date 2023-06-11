// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';
import { createAsyncThunk } from '@reduxjs/toolkit';

import { BadRequestError } from '@/api';
import { setErrorMessage } from '@/app/store/reducers/error';

import { PostsClient } from '@/api/posts';
import { PostsService } from '@/post/service';
import { CommentCreate, PostAddData, PostLikedAction, PostUpdateData } from '@/post';
import { postsSlice } from '@/app/store/reducers/posts';

const postsClient = new PostsClient();
export const postsService = new PostsService(postsClient);

export const createPost = createAsyncThunk(
    '/post/create',
    async function (post: PostAddData) {
        const postData = await postsService.createPost(post);

        return postData;
    }
);

export const updatePost = createAsyncThunk(
    '/post/update',
    async function (post: PostUpdateData) {
        const postData = await postsService.update(post);

        return postData;
    }
);

export const deletePost = createAsyncThunk(
    '/post/delete',
    async function (postId: string) {
        await postsService.delete(postId);

        return postId;
    }
);

export const getPost = createAsyncThunk(
    '/post',
    async (postId: string) => {
        const response = await postsService.getPost(postId);

        return response;
    }
);

export const getPostsHomePage = () => async function (dispatch: Dispatch) {
    try {
        const posts = await postsService.getPostsHomePage();

        dispatch(postsSlice.actions.setPostsHomePage(posts));
    } catch (error: any) {
        if (error instanceof BadRequestError) {
            dispatch(setErrorMessage('No valid posts info'));
        }
    }
};

export const getPostsProfile = createAsyncThunk(
    '/posts/profile',
    async (userId: string) => {
        const response = await postsService.getPostsProfile(userId);

        return response;
    }
);

export const sendComment = createAsyncThunk(
    '/comment',
    async (comment: CommentCreate) => {
        const commentData = await postsService.sendComment(comment);

        return commentData;
    }
);

export const likeAndDislikePost = createAsyncThunk(
    '/likeAndDislike/post',
    async function (likedPostActionData: PostLikedAction) {
        await postsService.likeAndDislike(likedPostActionData.postId);

        return { post_id: likedPostActionData.postId, isLiked: likedPostActionData.isLiked };
    }
);
