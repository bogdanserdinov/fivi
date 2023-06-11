import { useEffect, useState } from 'react';
import { Link, useParams } from 'react-router-dom';

import { UserPost } from '@components/User/Post';
import { Avatar } from '@components/common/Avatar';
import { UserSubscribersModal } from '@components/User/SubscibersModal';
import { UserSubscribesModal } from '@components/User/SubscribesModal';
import { PostCreateModal } from '@components/Post/CreateModal';
import { followUser, getUser, getUserProfile, unFollowUser } from '@/app/store/actions/users';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { RootState } from '@/app/store';
import { Subscribers, User, UserProfile } from '@/users';
import { getPostsProfile } from '@/app/store/actions/posts';
import { Post } from '@/post';
import { FollowData } from '@/followers';

import './index.scss';

const AVATAR_SIZE = 135;

const UserPage = () => {
    const [isOpenSubscribesModal, setIsOpenSubscribesModal] = useState(false);
    const [isOpenSubscribersModal, setIsOpenSubscribersModal] = useState(false);
    const [isOpenModal, setIsOpenModal] = useState(false);
    const { id } = useParams();

    const dispatch = useAppDispatch();

    const posts: Post[] = useAppSelector((state: RootState) => state.postsReducer.userProfilePosts);
    const userProfile: UserProfile = useAppSelector((state: RootState) => state.usersReducer.userProfile);
    const user: User = useAppSelector((state: RootState) => state.usersReducer.user);

    const openSubscribers = () => {
        setIsOpenSubscribersModal(true);
        setIsOpenSubscribesModal(false);
    };
    const openSubscribes = () => {
        setIsOpenSubscribesModal(true);
        setIsOpenSubscribersModal(false);
    };

    const handleSubscribing = () => {
        if (userProfile.isFollowed) {
            const subscription = userProfile.subscribers.filter((subscription: Subscribers) =>
                subscription.userId === user.userId);

            if (subscription[0]) {
                dispatch(unFollowUser(subscription[0].subscriptionId));
            }
        }
        else {
            dispatch(followUser(new FollowData(user.userId, userProfile.userId)));
        }
    };

    const onLoadPage = () => {
        if (id) {
            dispatch(getUser());
            dispatch(getUserProfile(id));
            dispatch(getPostsProfile(id));
        }
    };

    useEffect(() => {
        onLoadPage();
    }, []);

    return (
        <div className="user">
            <div className="user__top-side">
                <Avatar size={AVATAR_SIZE} userId={userProfile.userId} isAvatarExists={userProfile.isAvatarExists} />

                <div className="user__right-side">
                    <div className="user__info">
                        <p className="user__nick-name">{userProfile.username}</p>
                        {
                            user.userId === userProfile.userId ?
                                <Link className="user__info__button"
                                    to={`/user/${user.userId}/edit`}
                                >
                                    редагувати профіль
                                </Link>
                                :
                                userProfile.isFollowed ?
                                    <button className="user__info__button"
                                        onClick={() => handleSubscribing()}>
                                        відписатися
                                    </button>
                                    :
                                    <button className="user__info__button"
                                        onClick={() => handleSubscribing()}>
                                        Підписатися
                                    </button>
                        }

                    </div>
                    <div className="user__info">
                        <p className="user__info__text">{posts.length} публікацій</p>
                        <div className="user__info__text user__info__text--clickable"
                            onClick={() => openSubscribers()}
                        >
                            {userProfile.subscribers.length} підписники
                        </div>
                        <div className="user__info__text user__info__text--clickable"
                            onClick={() => openSubscribes()}
                        >
                            {userProfile.subscriptions.length} підписки
                        </div>
                    </div>
                </div>
            </div>
            {
                posts.length ?
                    <div className="user__posts">
                        {
                            posts.map((post: Post) =>
                                <UserPost post={post} key={post.postId} />
                            )

                        }
                    </div>
                    : <div className="user__posts__no-item">
                        <h2 className="user__posts__no-item__title"> У користувача немає публікацій</h2>
                        {
                            user.userId === userProfile.userId &&
                            <button
                                onClick={() => setIsOpenModal(true)}
                                className="user__posts__no-item__add-post"
                            >
                                Додати пост
                            </button>
                        }
                    </div>
            }
            {isOpenModal && <PostCreateModal setIsOpenModal={setIsOpenModal} />}
            {isOpenSubscribersModal && <UserSubscribersModal setIsOpenModal={setIsOpenSubscribersModal} />}
            {isOpenSubscribesModal && <UserSubscribesModal setIsOpenModal={setIsOpenSubscribesModal} />}
        </div>
    );
};

export default UserPage;
