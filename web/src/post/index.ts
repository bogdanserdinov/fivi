export class Comment {
    public constructor(
        public commentId: string = '',
        public text: string = '',
        public postId: string = '',
        public username: string = '',
        public userId: string = '',
        public userImage: string = '',
        public isAvatarExists: boolean = false,
    ) { }
}

export class CommentCreate {
    public constructor(
        public postId: string = '',
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
        public postId: string = '',
        public description: string = '',
        public creatorId: string = '',
        public creatorUsername: string = '',
        public images: string[] = [],
        public numOfImages: number = 0,
        public numOfLikes: number = 0,
        public numOfComments: number = 0,
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
        public postId: string = '',
        public text: string = '',
        public images: string[] = [],
    ) { };
}

export class PostLikedAction {
    public constructor(
        public postId: string = '',
        public isLiked: boolean = false,
    ) { };
}
