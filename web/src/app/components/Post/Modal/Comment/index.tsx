import { Avatar } from '@components/common/Avatar';
import { Link, useNavigate } from 'react-router-dom';
import { useAppDispatch } from '@/app/hooks/useReduxToolkit';
import { getPostsProfile } from '@/app/store/actions/posts';
import { getUserProfile } from '@/app/store/actions/users';
import { setCurrentId } from '@/app/store/reducers/posts';

import './index.scss';

export const CommentPostModal: React.FC<{
    text: string;
    isAvatarExists: boolean;
    nickname: string;
    commentorId: string;
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
}> = ({ text, isAvatarExists, nickname, commentorId, setIsOpenModal }) => {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const handleNavigateToUser = async() => {
        await dispatch(getUserProfile(commentorId));
        await dispatch(setCurrentId(commentorId));
        await dispatch(getPostsProfile(commentorId));
        navigate(`/user/${commentorId}`);
        setIsOpenModal(false);
    };

    return (
        <div className="post-modal-comment" >

            <div onClick={() => handleNavigateToUser()} className="post-modal-comment__photo">
                <Avatar size={40} userId={commentorId} isAvatarExists={isAvatarExists} />
            </div>

            <div className="post-modal-comment__text">
                <div onClick={() => handleNavigateToUser()}
                    className="post-modal-comment__nickname">
                    {nickname}
                </div>
                <p className="post-modal-comment__comment">{text}</p>

            </div>

        </div>);
};


