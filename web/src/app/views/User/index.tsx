import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

import { UserPost } from '@components/User/Post';
import { Avatar } from '@components/common/Avatar';
import { UserSubscribersModal } from '@components/User/SubscibersModal';
import { UserSubscribesModal } from '@components/User/SubscribesModal';
import { getUserProfile } from '@/app/store/actions/users';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { RootState } from '@/app/store';
import { UserProfile } from '@/users';
import { getPostsProfile } from '@/app/store/actions/posts';
import { Post } from '@/post';

import './index.scss';

const AVATAR_SIZE = 135;
const LAST_ITEM_PATH_INCREMENT = 1;

const UserPage = () => {
    const [isOpenSubscribesModal, setIsOpenSubscribesModal] = useState(false);
    const [isOpenSubscribersModal, setIsOpenSubscribersModal] = useState(false);
    const getLastItem = (thePath: string) => thePath.substring(thePath.lastIndexOf('/') + LAST_ITEM_PATH_INCREMENT);
    const dispatch = useAppDispatch();

    const posts: Post[] = useAppSelector((state: RootState) => state.postsReducer.userProfilePosts);
    const user: UserProfile = useAppSelector((state: RootState) => state.usersReducer.userProfile);

    const openSubscribers = () => {
        setIsOpenSubscribersModal(true);
        setIsOpenSubscribesModal(false);
    };
    const openSubscribes = () => {
        setIsOpenSubscribesModal(true);
        setIsOpenSubscribersModal(false);
    };

    console.log(posts);
    useEffect(() => {
        const userId = getLastItem(window.location.pathname);
        dispatch(getUserProfile(userId));
        dispatch(getPostsProfile(userId));
    }, []);

    return (
        <div className="user">
            <div className="user__top-side">
                <Avatar size={AVATAR_SIZE} photo={`${window.location.origin}/images/users/${user.userId}.png`} isAvatarExists={user.isAvatarExists} />

                <div className="user__right-side">
                    <div className="user__info">
                        <p className="user__nick-name">{user.username}</p>
                        <Link className="user__info__button"
                            to={`/user/${user.userId}/edit`}>редагувати профіль</Link>
                    </div>
                    <div className="user__info">
                        <p className="user__info__text">{posts.length} публікацій</p>
                        <div className="user__info__text user__info__text--clickable"
                            onClick={() => openSubscribers()}
                        >
                            {user.subscribers.length} підписники
                        </div>
                        <div className="user__info__text user__info__text--clickable"
                            onClick={() => openSubscribes()}
                        >
                            {user.subscriptions.length} підписки
                        </div>
                    </div>
                </div>
            </div>

            <div className="user__posts">
                {posts.map((post: Post) =>
                    <UserPost post={post} key={post.postId} />
                )}
            </div>
            {isOpenSubscribersModal && <UserSubscribersModal setIsOpenModal={setIsOpenSubscribersModal} subscribers={user.subscribers} />}
            {isOpenSubscribesModal && <UserSubscribesModal setIsOpenModal={setIsOpenSubscribesModal} subscribes={user.subscriptions} />}
        </div>
    );
};

export default UserPage;
