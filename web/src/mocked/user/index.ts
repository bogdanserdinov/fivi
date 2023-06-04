import mockedAvatar from '@static/img/post/mockedAvatar.png';
import { post } from '../posts';

const subcscribes = {
    id: '00000000-0000-0000-0000-000000000000',
    avatar: mockedAvatar,
    nickname: 'ann_ann',
};

const subcscribers = {
    id: '00000000-0000-0000-0000-000000000000',
    avatar: mockedAvatar,
    nickname: 'ann_ann',
    isSubcribe: false,
};

const subcscribers1 = {
    id: '00000000-0000-0000-0000-000000000000',
    avatar: mockedAvatar,
    nickname: 'ann_ann',
    isSubcribe: true,
};

export const user = {
    id: '00000000-0000-0000-0000-000000000000',
    name: 'Ann',
    surname: 'Ann',
    phoneNumber: '0999999999',
    email: 'ann@gmail.com',
    avatar: mockedAvatar,
    nickname: 'ann_ann',
    posts: [
        post, post, post, post, post, post,
    ],
    subscribers: [
        subcscribers, subcscribers1, subcscribers1, subcscribers, subcscribers, subcscribers, subcscribers, subcscribers,
    ],
    subscribes: [
        subcscribes, subcscribes, subcscribes, subcscribes, subcscribes, subcscribes, subcscribes,
    ],
};
