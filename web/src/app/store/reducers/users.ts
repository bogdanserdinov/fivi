import { createSlice } from '@reduxjs/toolkit';
import type { PayloadAction } from '@reduxjs/toolkit';

import { getUserProfile, updateUser } from '../actions/users';
import { User, UserProfile } from '@/users';

/** Exposes channels state */
class UsersState {
    /** class implementation */
    constructor(
        public user: User = new User(),
        public mnemonicPhrases: string[] = [],
        public userProfile: UserProfile = new UserProfile(),
    ) { }
}

const initialState: UsersState = {
    user: new User(),
    mnemonicPhrases: [],
    userProfile: new UserProfile(),
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
    },
});

// Action creators are generated for each case reducer function
export const { setUser, setMnemonicPhrases } = userSlice.actions;

export default userSlice.reducer;
