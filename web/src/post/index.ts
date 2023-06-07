export class Comment {
    public constructor(
        public commentId: string = '',
        public text: string = '',
        public postId: string = '',
        public username: string = '',
        public userId: string = '',
        public userImage: string = '',
    ) { }
}

export class CommentCreate {
    public constructor(
        public postId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public text: string = '',
    ) { }
}

export class Creator {
    public constructor(
        public creatorId: string = '',
        public email: string = '',
        public username: string = '',
        public numOfPosts: number = 0,
        public subscribers: any[] = [],
        public subscribtions: any[] = [],
        public isAvatarExists: boolean = false,
    ) { }
}

export class Post {
    public constructor(
        public postId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public description: string = '',
        public creatorId: string = '',
        public creatorUsername: string = '',
        public images: string[] = [],
        public num_of_images: number = 0,
        public num_of_likes: number = 0,
        public num_of_comments: number = 0,
        public comments: Comment[] = [],
        public isLiked: boolean = false,
        public creatorProfile: Creator = new Creator()
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

export class PostLikedAction {
    public constructor(
        public postId: string = 'did:ion:test:EiAH4mOt_BJthhMkoizi9538NYHtP4-ai49hzQU9WSIJjA',
        public isLiked: boolean = false,
    ) { };
}
