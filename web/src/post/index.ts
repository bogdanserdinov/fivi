export class Comment {
    public constructor(
        public commentId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public text: string = '',
        public creatorId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public creatorNickName: string = '',
    ) { }
}

export class Post {
    public constructor(
        public postId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public favorites: number = 0,
        public isFavorite: boolean = false,
        public description: string = '',
        public comments: Comment[] = [],
        public posts: [] = []
    ) { };
}

