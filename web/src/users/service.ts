import { Subscribers, User, UserLoginData, UserProfile, UserRegisterData, UserUpdate } from '.';
import { UsersClient } from '@/api/users';
import { FollowData } from '@/followers';

/**
 * Exposes all users related logic.
 */
export class UsersService {
    private readonly users: UsersClient;
    /** UsersService contains http implementation of users API  */
    public constructor(users: UsersClient) {
        this.users = users;
    }
    /** handles user registration */
    public async register(user: UserRegisterData): Promise<string> {
        return await this.users.register(user);
    }
    /** handles user login */
    public async login(user: UserLoginData): Promise<string> {
        return await this.users.login(user);
    }
    /** gets user info */
    public async getUser(): Promise<User> {
        return await this.users.getUser();
    }

    /** gets user info */
    public async getUserProfile(userId: string): Promise<UserProfile> {
        return await this.users.getUserProfile(userId);
    }


    /** updates user info */
    public async update(user: UserUpdate): Promise<User> {
        return await this.users.update(user);
    }

    /** logouts */
    public async getMnemonicPhrases(): Promise<string[]> {
        return await this.users.getMnemonicPhrases();
    }

    /** Searches users */
    public async searchUsers(text: string): Promise<UserProfile[]> {
        return await this.users.searchUsers(text);
    }

    /** Searches users */
    public async followUser(followData: FollowData): Promise<Subscribers> {
        return await this.users.followUser(followData);
    }

    /** Searches users */
    public async unFollowUser(userId: string): Promise<void> {
        await this.users.unFollowUser(userId);
    }
}
