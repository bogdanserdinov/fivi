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

const usersClient = new UsersClient();
export const usersService = new UsersService(usersClient);

export const register = createAsyncThunk(
    '/auth/register',
    async function(user: UserRegisterData) {
        await usersService.register(user);
    }
);


export const login = createAsyncThunk(
    '/auth/login',
    async function(user: UserLoginData) {
        await usersService.login(user);
    }
);

export const updateUser = createAsyncThunk(
    '/users',
    async function(user: UserUpdate) {
        await usersService.update(user);
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
