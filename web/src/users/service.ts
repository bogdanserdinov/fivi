import { User, UserLoginData, UserRegisterData, UserUpdate } from '.';
import { UsersClient } from '@/api/users';

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
    public async register(user: UserRegisterData): Promise<void> {
        await this.users.register(user);
    }
    /** handles user login */
    public async login(user: UserLoginData): Promise<void> {
        await this.users.login(user);
    }
    /** gets user info */
    public async getUser(): Promise<User> {
        return await this.users.getUser();
    }

    /** updates user info */
    public async update(user: UserUpdate): Promise<void> {
        await this.users.update(user);
    }

    /** logouts */
    public async getMnemonicPhrases(): Promise<string[]> {
        return await this.users.getMnemonicPhrases();
    }
}
