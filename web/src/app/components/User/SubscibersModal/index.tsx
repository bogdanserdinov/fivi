import { Avatar } from '@components/common/Avatar';
import { Modal } from '@components/common/Modal';
import { useNavigate } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { RootState } from '@/app/store';
import { Subscribers, User, UserProfile } from '@/users';
import { followUser, getUserProfile } from '@/app/store/actions/users';
import { setCurrentId } from '@/app/store/reducers/posts';
import { getPostsProfile } from '@/app/store/actions/posts';
import { FollowData } from '@/followers';

import './index.scss';

export const UserSubscribersModal: React.FC<{
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
}> = ({ setIsOpenModal }) => {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const user: User = useAppSelector((state: RootState) => state.usersReducer.user);
    const userProfile: UserProfile = useAppSelector((state: RootState) => state.usersReducer.userProfile);

    const handleNavigateToUser = async(userId: string) => {
        await dispatch(getUserProfile(userId));
        await dispatch(setCurrentId(userId));
        await dispatch(getPostsProfile(userId));
        navigate(`/user/${userId}`);
        setIsOpenModal(false);
    };

    const handleSubscribing = (userId: string) => {
        dispatch(followUser(new FollowData(user.userId, userId)));
    };

    return (
        <Modal setIsOpenModal={setIsOpenModal}>
            <div className="subscribers">
                {userProfile.subscribers.length ?
                    userProfile.subscribers.map((subscriber: Subscribers) =>
                        <div className="subscribers__item">
                            <div className="subscribes__item__info">
                                <div className="subscribes__item__info" onClick={() => handleNavigateToUser(subscriber.userId)}>
                                    <Avatar size={50} userId={subscriber.userId} isAvatarExists={subscriber.isAvatarExists} />
                                    <p className="subscribers__item__nickname">{subscriber.username}</p>
                                </div>
                                {!subscriber.isSubscribed &&
                                    <>
                                        &#8226;
                                        <button
                                            className="subscribers__item__subscribe"
                                            onClick={() => handleSubscribing(subscriber.userId)}
                                        >
                                            Підписатися
                                        </button>
                                    </>
                                }
                            </div>
                            {userProfile.userId === user.userId &&
                                <button className="subscribers__item__delete">
                                    Видалити
                                </button>
                            }
                        </div>
                    ) :
                    <p>Ще нема підписніків</p>
                }
            </div>

        </Modal>);
};
