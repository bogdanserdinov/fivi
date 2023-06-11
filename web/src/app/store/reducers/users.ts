import { createSlice, current } from '@reduxjs/toolkit';
import type { PayloadAction } from '@reduxjs/toolkit';

import {
    followUser,
    getUserProfile,
    searchUsers,
    unFollowUser,
    updateUser,
} from '../actions/users';
import { User, UserProfile } from '@/users';

/** Exposes channels state */
class UsersState {
    /** class implementation */
    constructor(
        public user: User = new User(),
        public mnemonicPhrases: string[] = [],
        public userProfile: UserProfile = new UserProfile(),
        public foundedUsers: UserProfile[] = []
    ) { }
}

const initialState: UsersState = {
    user: new User(),
    mnemonicPhrases: [],
    userProfile: new UserProfile(),
    foundedUsers: [],

};

export const userSlice = createSlice({
    name: 'usersReducer',
    initialState,
    reducers: {
        setUser: (state, action: PayloadAction<User>) => {
            state.user = action.payload;
        },
        setMnemonicPhrases: (state, action: PayloadAction<string[]>) => {
            state.mnemonicPhrases = action.payload;
        },
    },
    extraReducers: (builder) => {
        builder.addCase(getUserProfile.fulfilled, (state, action) => {
            state.userProfile = action.payload;
        });
        builder.addCase(updateUser.fulfilled, (state, action) => {
            state.user = action.payload;
        });
        builder.addCase(searchUsers.fulfilled, (state, action) => {
            state.foundedUsers = action.payload;
        });

        builder.addCase(followUser.fulfilled, (state, action) => {
            const subscribers =
                state.userProfile.userId !== state.user.userId ?
                    state.userProfile.subscribers.concat(action.payload)
                    :
                    state.userProfile.subscribers;

            const subscriptions = state.userProfile.userId === state.user.userId ?
                state.userProfile.subscriptions.concat(action.payload)
                :
                state.userProfile.subscriptions;

            return {
                ...state,
                userProfile: {
                    ...state.userProfile,
                    subscribers: subscribers,
                    subscriptions: subscriptions,
                    isFollowed: true,
                },
            };
        });

        builder.addCase(unFollowUser.fulfilled, (state, action) => {
            const subscribers =
                state.userProfile.subscribers.filter((subscriber) => subscriber.subscriptionId !== action.payload);
            const subscriptions =
                state.userProfile.subscriptions.filter((subscribe) => subscribe.subscriptionId !== action.payload);

            return {
                ...state,
                userProfile: {
                    ...state.userProfile,
                    subscribers: subscribers,
                    subscriptions: subscriptions,
                    isFollowed: false,
                },
            };
        });
    },
});

// Action creators are generated for each case reducer function
export const { setUser, setMnemonicPhrases } = userSlice.actions;

export default userSlice.reducer;
