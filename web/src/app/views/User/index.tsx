import { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';

import { UserPost } from '@components/User/Post';
import { Avatar } from '@components/common/Avatar';
import { UserSubscribersModal } from '@components/User/SubscibersModal';
import { UserSubscribesModal } from '@components/User/SubscribesModal';
import { user } from '@/mocked/user';

import './index.scss';

const User = () => {
    const [isOpenSubscribesModal, setIsOpenSubscribesModal]=useState(false);
    const [isOpenSubscribersModal, setIsOpenSubscribersModal] = useState(false);
    const location=useLocation();

    const openSubscribers = () => {
        setIsOpenSubscribersModal(true);
        setIsOpenSubscribesModal(false);
    };
    const openSubscribes = () => {
        setIsOpenSubscribesModal(true);
        setIsOpenSubscribersModal(false);
    };

    return (
        <div className="user">
            <div className="user__top-side">
                <Avatar size={135} photo={user.avatar} />
                <div className="user__right-side">
                    <div className="user__info">
                        <p className="user__nick-name">{user.nickname}</p>
                        <Link className="user__info__button"
                            to={`/user/${user.id}/edit`}>редагувати профіль</Link>
                    </div>
                    <div className="user__info">
                        <p className="user__info__text">{user.posts.length} публікацій</p>
                        <div className="user__info__text user__info__text--clickable"
                            onClick={() => openSubscribers()}
                        >
                            {user.subscribers.length} підписники
                        </div>
                        <div className="user__info__text user__info__text--clickable"
                            onClick={() => openSubscribes()}
                        >
                            {user.subscribes.length} підписки
                        </div>
                    </div>
                </div>
            </div>

            <div className="user__posts">
                {user.posts.map((post) =>
                    <UserPost post={post} />
                )}
            </div>
            {isOpenSubscribersModal && <UserSubscribersModal setIsOpenModal={setIsOpenSubscribersModal} subscribers={user.subscribers} />}
            {isOpenSubscribesModal && <UserSubscribesModal setIsOpenModal={setIsOpenSubscribesModal} subscribes={user.subscribes} />}
        </div>
    );
};

export default User;
