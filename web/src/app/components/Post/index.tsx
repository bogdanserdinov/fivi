import { Link } from 'react-router-dom';
import { useState } from 'react';

import commentIcon from '@static/img/post/commentIcon.png';
import favoriteIcon from '@static/img/post/favoriteIcon.png';
import isFavoriteIcon from '@static/img/post/isFavoriteIcon.png';

import { PostSlider } from '@components/Post/Slider';
import { PostModal } from '@components/Post//Modal';
import { Avatar } from '@components/common/Avatar';

import './index.scss';

const ONE_LIKE = 1;
const AVATAR_SIZE = 40;

export const Post: React.FC<{ post: any }> = ({ post }) => {
    const [isOpenModal, setIsOpenModal] = useState<boolean>(false);
    const [isFavoritePost, setIsFavoritePost] = useState<boolean>(false);
    const [currentComment, setCurrentComment] = useState<string>('');

    const onChangeAddComment = (e: any) => {
        setCurrentComment(e.target.value);
    };

    return (
        <>
            <div className="post">
                <Link to={`/user/${post.creator.id}`} className="post__user">
                    <Avatar size={AVATAR_SIZE} photo={post.creator.creatorAvatar} />
                    <p className="post__user__text">
                        {post.creator.name}
                    </p>
                </Link>
                <PostSlider sliderImages={post.photos} />
                <div className="post__bottom-side">
                    <button className="post__button" onClick={()=>setIsFavoritePost(!isFavoritePost)}>
                        {isFavoritePost ?
                        <img
                            className="post__button__image"
                            src={isFavoriteIcon}
                                alt="favorite" />
                            :
                        <img
                            className="post__button__image"
                            src={favoriteIcon}
                                alt="favorite" />
                        }
                        
                    </button>
                    <button className="post__button">
                        <img
                            className="post__button__image"
                            src={commentIcon}
                            alt="comment" />
                    </button >
                </div>
                <div className="post__likes">
                    {post.favorites} {post.favorites === ONE_LIKE ? 'вподобайка' : 'вподобайки'}
                </div>
                {post.description &&
                <div className="post__description">
                    <Link
                        to={`/user/${post.creator.id}`}
                        className="post__description__name">
                        {post.creator.name}
                    </Link>
                    <p className="post__description__text">
                        {post.description}
                    </p>
                </div>
                }
                <div className="post__show-all-comments"
                    onClick={() => setIsOpenModal(true)}>
                показати все коментарі ({post.comments.length})
                </div>
                <div className="post__add-comment">
                    <input
                        className="post__add-comment__input"
                        type="text"
                        placeholder="Додайте коментарій..."
                        onChange={onChangeAddComment}
                        value={currentComment}
                    />
                    {currentComment &&
                    <button className="post__add-comment__send">Відправити</button>
                    }
                </div>
            </div>
            {isOpenModal && <PostModal post={post} setIsOpenModal={setIsOpenModal} />}
        </>
    );
};
