export class FollowData {
    public constructor(
        public userId: string = '',
        public userToFollowId: string = '',
    ) { }
}

export class FollowersData {
    public constructor(
        public userId: string = '',
        public userName: string = '',
    ) { }
}
