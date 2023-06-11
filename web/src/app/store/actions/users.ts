// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch } from 'redux';
import { createAsyncThunk } from '@reduxjs/toolkit';

import { UsersClient } from '@/api/users';
import { BadRequestError } from '@/api';
import { UsersService } from '@/users/service';
import { userSlice } from '@/app/store/reducers/users';
import { setErrorMessage } from '@/app/store/reducers/error';
import { UserLoginData, UserRegisterData, UserUpdate } from '@/users';
import { setLocalStorageItem } from '@/app/utils/localStorage';
import { FollowData } from '@/followers';

const usersClient = new UsersClient();
export const usersService = new UsersService(usersClient);

export const register = createAsyncThunk(
    '/auth/register',
    async function(user: UserRegisterData) {
        const token = await usersService.register(user);
        setLocalStorageItem('AUTH_TOKEN', token);
    }
);

export const login = createAsyncThunk(
    '/auth/login',
    async function(user: UserLoginData) {
        const token = await usersService.login(user);
        setLocalStorageItem('AUTH_TOKEN', token);
    }
);

export const updateUser = createAsyncThunk(
    '/users/update',
    async function(user: UserUpdate) {
        const userData = await usersService.update(user);

        return userData;
    }
);

export const getUser = () => async function(dispatch: Dispatch) {
    try {
        const user = await usersService.getUser();
        dispatch(userSlice.actions.setUser(user));
    } catch (error: any) {
        if (error instanceof BadRequestError) {
            dispatch(setErrorMessage('No valid user info'));
        }
    }
};

export const getUserProfile = createAsyncThunk(
    '/get/users',
    async function(userId: string) {
        const user = await usersService.getUserProfile(userId);

        return user;
    }
);

export const getMnemonicPhrases = () => async function(dispatch: Dispatch) {
    try {
        const phrases = await usersService.getMnemonicPhrases();
        dispatch(userSlice.actions.setMnemonicPhrases(phrases));
    } catch (error: any) {
        if (error instanceof BadRequestError) {
            dispatch(setErrorMessage('No valid mnemonic info'));
        }
    }
};

export const searchUsers = createAsyncThunk(
    '/search/users',
    async function(text: string) {
        const users = await usersService.searchUsers(text);

        return users;
    }
);

export const followUser = createAsyncThunk(
    '/follow/user',
    async function(followData: FollowData) {
        const followedUser = await usersService.followUser(followData);

        return followedUser;
    }
);

export const unFollowUser = createAsyncThunk(
    '/unfollow/user',
    async function(subscriptionId: string) {
        await usersService.unFollowUser(subscriptionId);

        return subscriptionId;
    }
);
