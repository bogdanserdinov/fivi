export class User {
    public constructor(
        public userId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public username: string = '',
        public email: string = '',
        public isAvatarExists: boolean = false,
    ) { };
}

export class UserProfile {
    public constructor(
        public userId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public username: string = '',
        public email: string = '',
        public subscribers: [] = [],
        public subscriptions: [] = [],
        public isAvatarExists: boolean = false,
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
