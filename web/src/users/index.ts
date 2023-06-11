export class User {
    public constructor(
        public userId: string = '',
        public username: string = '',
        public email: string = '',
        public isAvatarExists: boolean = false,
    ) { };
}

export class Subscribers {
    public constructor(
        public subscriptionId: string = '',
        public userId: string = '',
        public username: string = '',
        public isAvatarExists: boolean = false,
        public isSubscribed: boolean = false
    ) { };
}


export class UserProfile {
    public constructor(
        public userId: string = '',
        public username: string = '',
        public email: string = '',
        public subscribers: Subscribers[] = [],
        public subscriptions: Subscribers[] = [],
        public isAvatarExists: boolean = false,
        public isFollowed: boolean = false,
    ) { };
}

export class UserRegisterData {
    public constructor(
        public email: string = '',
        public name: string = '',
        public mnemonic: string[] = [],

    ) { }
}

export class UserLoginData {
    public constructor(
        public username: string = '',
        public mnemonic: string[] = [],
    ) { }
}

export class UserUpdate {
    public constructor(
        public name: string = '',
        public username: string = '',
        public email: string = '',
        public image: string = ''
    ) { }
}
