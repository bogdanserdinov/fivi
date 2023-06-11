import { APIClient } from '.';

import { Subscribers, User, UserLoginData, UserProfile, UserRegisterData, UserUpdate } from '@/users';
import { FollowData } from '@/followers';

/**
 * UsersClient is a http implementation of users API.
 * Exposes all users-related functionality.
 */
export class UsersClient extends APIClient {
    /** exposes user registration logic */
    public async register(user: UserRegisterData): Promise<string> {
        const path = `${this.ROOT_PATH}/auth/v1/register`;
        const response = await this.http.post(path, JSON.stringify(user));

        if (!response.ok) {
            await this.handleError(response);
        }
        const token = await response.json();

        return token.jwt;
    }

    /** exposes user login logic */
    public async login(user: UserLoginData): Promise<string> {
        const path = `${this.ROOT_PATH}/auth/v1/login`;
        const response = await this.http.post(
            path,
            JSON.stringify(user)
        );

        if (!response.ok) {
            await this.handleError(response);
        }
        const token = await response.json();

        return token.jwt;
    }

    /** Gets user */
    public async update(user: UserUpdate): Promise<User> {
        const path = `${this.ROOT_PATH}/profile/v1`;
        const response = await this.http.put(path, JSON.stringify(user));

        if (!response.ok) {
            await this.handleError(response);
        }
        const userData = await response.json();

        return new UserProfile(
            userData.id,
            userData.username,
            userData.email,
            userData.subscribers,
            userData.subscribtions,
            userData.isAvatarExists,
        );
    }

    /** Gets user */
    public async getUserProfile(userId: string): Promise<UserProfile> {
        const path = `${this.ROOT_PATH}/profile/v1?userDid=${userId}`;

        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const user = await response.json();

        const subscribers = user.subscribers.map((subscriber: any) =>
            new Subscribers(subscriber.id, subscriber.userId, subscriber.username, subscriber.isAvatarExists, subscriber.isSubscribed));

        const subscribtions = user.subscriptions.map((subscribtion: any) =>
            new Subscribers(subscribtion.id, subscribtion.userId, subscribtion.username, subscribtion.isAvatarExists, subscribtion.isSubscribed));

        return new UserProfile(
            user.id,
            user.username,
            user.email,
            subscribers,
            subscribtions,
            user.isAvatarExists,
            user.isFollowed
        );
    };


    /** Gets user */
    public async getUser(): Promise<User> {
        const path = `${this.ROOT_PATH}/profile/v1`;

        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const user = await response.json();

        return new User(
            user.id,
            user.username,
            user.email,
            user.isAvatarExists,
        );
    }


    /** Gets user */
    public async getMnemonicPhrases(): Promise<string[]> {
        const path = `${this.ROOT_PATH}/auth/v1/generate`;
        const response = await this.http.post(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const mnemonicPhrases = await response.json();

        return mnemonicPhrases.mnemonic;
    }

    /** Searches users */
    public async searchUsers(text: string): Promise<UserProfile[]> {
        const path = `${this.ROOT_PATH}/profile/v1/search?username=${text}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const users = await response.json();


        return users.profiles.map((user: any) => {
            const subscribers = user.subscribers.map((subscriber: any) =>
                new Subscribers(subscriber.id, subscriber.username));

            const subscribtions = user.subscribers.map((subscribtion: any) =>
                new Subscribers(subscribtion.id, subscribtion.username));

            return new UserProfile(
                user.id,
                user.username,
                user.email,
                subscribers,
                subscribtions,
                user.isAvatarExists,
                user.isFollowed
            );
        });
    };

    /** exposes user following logic */
    public async followUser(followData: FollowData): Promise<Subscribers> {
        const path = `${this.ROOT_PATH}/followers/v1`;
        const response = await this.http.post(path, JSON.stringify(followData));

        if (!response.ok) {
            await this.handleError(response);
        }
        const followedUser = await response.json();

        return new Subscribers(
            followedUser.id,
            followedUser.userId,
            followedUser.username,
            followedUser.isAvatarExists,
            followedUser.isSubscribed,
        );
    }

    /** exposes user unfollowing logic */
    public async unFollowUser(userId: string): Promise<void> {
        const path = `${this.ROOT_PATH}/followers/v1/follower/${userId}`;
        const response = await this.http.delete(path);

        if (!response.ok) {
            await this.handleError(response);
        }
    }
}
