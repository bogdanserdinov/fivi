import { Avatar } from '@components/common/Avatar';
import { Modal } from '@components/common/Modal';
import { Link, useNavigate } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { RootState } from '@/app/store';
import { followUser, getUserProfile, unFollowUser } from '@/app/store/actions/users';
import { FollowData } from '@/followers';
import { Subscribers, User, UserProfile } from '@/users';

import './index.scss';
import { setCurrentId } from '@/app/store/reducers/posts';
import { getPostsProfile } from '@/app/store/actions/posts';

export const UserSubscribesModal: React.FC<{
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

    const rejectSubscribe = (subscriptionId: string) => {
        dispatch(unFollowUser(subscriptionId));
    };

    const handleSubscribing = (userId: string) => {
        dispatch(followUser(new FollowData(user.userId, userId)));
    };


    return (
        <Modal setIsOpenModal={setIsOpenModal}>
            <div className="subscribes">
                {userProfile.subscriptions.length ?
                    userProfile.subscriptions.map((subscribe: Subscribers) =>
                        <div className="subscribes__item" key={subscribe.userId}>
                            <div className="subscribes__item__info" onClick={() => handleNavigateToUser(subscribe.userId)}>
                                <Avatar size={50} userId={subscribe.userId} isAvatarExists={subscribe.isAvatarExists} />
                                <p className="subscribes__item__nickname">{subscribe.username}</p>
                            </div>
                            {subscribe.isSubscribed ?
                                <button
                                    className="subscribes__item__reject-subscribes"
                                    onClick={() => rejectSubscribe(subscribe.subscriptionId)}
                                >
                                    Підписки
                                </button> :
                                <button
                                    className="subscribes__item__reject-subscribes"
                                    onClick={() => handleSubscribing(subscribe.userId)}
                                >
                                    Підписатися
                                </button>
                            }

                        </div>
                    ) :
                    <p>Ще нема підписок</p>
                }
            </div>
        </Modal>
    );
};
