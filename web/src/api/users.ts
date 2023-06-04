import { APIClient } from '.';

import { User, UserLoginData, UserRegisterData, UserUpdate } from '@/users';

/**
 * UsersClient is a http implementation of users API.
 * Exposes all users-related functionality.
 */
export class UsersClient extends APIClient {
    /** exposes user registration logic */
    public async register(user: UserRegisterData): Promise<void> {
        const path = `${this.ROOT_PATH}/auth/v1/register`;
        const response = await this.http.post(path, JSON.stringify(user));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** exposes user login logic */
    public async login(user: UserLoginData): Promise<void> {
        const path = `${this.ROOT_PATH}/auth/v1/login`;
        const response = await this.http.post(
            path,
            JSON.stringify(user)
        );

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** Gets user */
    public async update(user: UserUpdate): Promise<void> {
        const path = `${this.ROOT_PATH}/profile/v1`;
        const response = await this.http.put(path, JSON.stringify(user));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

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

        return mnemonicPhrases;
    }
}
