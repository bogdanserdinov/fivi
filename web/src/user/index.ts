export class User {
    public constructor(
        public userId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public name: string = '',
        public surname: string = '',
        public phoneNumber: string = '',
        public email: string = '',
        public avatar: string = '',
        public nickName: string = '',
        public posts: [] = []
    ) { };
}
