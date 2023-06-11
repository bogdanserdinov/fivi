import { Avatar } from '@components/common/Avatar';
import { useNavigate } from 'react-router-dom';
import { useAppDispatch } from '@/app/hooks/useReduxToolkit';
import { getPostsProfile } from '@/app/store/actions/posts';
import { getUserProfile } from '@/app/store/actions/users';
import { UserProfile } from '@/users';

import './index.scss';
import { setCurrentId } from '@/app/store/reducers/posts';

export const SearchingModal: React.FC<{
    setIsSearching: React.Dispatch<React.SetStateAction<boolean>>;
    foundedUsers: UserProfile[];
    classname?: string;
}> = ({ setIsSearching, foundedUsers, classname }) => {
    const navigate = useNavigate();
    const dispatch = useAppDispatch();

    const handleRedirectUserProfile = (userId: string) => {
        dispatch(getUserProfile(userId));
        dispatch(getPostsProfile(userId));
        dispatch(setCurrentId(userId));
        navigate(`/user/${userId}`);
        setIsSearching(false);
    };

    return (
        <div className={`searching-modal ${classname && classname}`}>
            <div className="searching-modal__content">
                {foundedUsers.map((user: UserProfile) =>
                    <div className="searching-modal__info" key={user.userId} onClick={() => handleRedirectUserProfile(user.userId)}>
                        <Avatar isAvatarExists={user.isAvatarExists} size={40} userId={user.userId} />
                        <p className="searching-modal__username">{user.username}</p>
                    </div>
                )}
            </div>

        </div>);
};
