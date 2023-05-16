import { Avatar } from '@components/common/Avatar';
import { Link } from 'react-router-dom';

import './index.scss';

export const CommentPostModal: React.FC<{
    text: string;
    photo: string;
    nickname: string;
    commentorId: string;
}> = ({ text, photo, nickname, commentorId }) =>
    <div className="post-modal-comment" >
        <Link to={`user/${commentorId}`}>
            <Avatar size={40} photo={photo} />
        </Link>
        <p className="post-modal-comment__text">
            <Link to={`user/${commentorId}`}
                className="post-modal-comment__nickname">
                {nickname}
            </Link>
            {text}
        </p>
    </div>;


