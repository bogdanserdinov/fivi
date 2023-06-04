export class Comment {
    public constructor(
        public commentId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public text: string = '',
        public creatorId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public creatorNickName: string = '',
    ) { }
}

export class CommentCreate {
    public constructor(
        public postId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public text: string = '',
    ) { }
}

export class Post {
    public constructor(
        public postId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public favorites: number = 0,
        public isFavorite: boolean = false,
        public description: string = '',
        public comments: Comment[] = [],
    ) { };
}

export class PostAddData {
    public constructor(
        public text: string = '',
        public images: string[] = [],
    ) { };
}


export class PostUpdateData {
    public constructor(
        public postId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public text: string = '',
        public images: string[] = [],
    ) { };
}


