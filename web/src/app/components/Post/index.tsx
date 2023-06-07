import { Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { useAppDispatch } from '@/app/hooks/useReduxToolkit';

import commentIcon from '@static/img/post/commentIcon.png';
import favoriteIcon from '@static/img/post/favoriteIcon.png';
import isFavoriteIcon from '@static/img/post/isFavoriteIcon.png';

import { PostModal } from '@components/Post/Modal';
import { Avatar } from '@components/common/Avatar';
import { PostSlider } from './Slider';
import { Post, PostLikedAction } from '@/post';
import { likeAndDislikePost } from '@/app/store/actions/posts';

import './index.scss';

const ONE_LIKE = 1;
const AVATAR_SIZE = 40;

export const PostPage: React.FC<{ post: Post }> = ({ post }) => {
    const [isOpenModal, setIsOpenModal] = useState<boolean>(false);
    const [isFavoritePost, setIsFavoritePost] = useState<boolean>(false);
    const [currentComment, setCurrentComment] = useState<string>('');
    const dispatch = useAppDispatch();

    const onChangeAddComment = (e: any) => {
        setCurrentComment(e.target.value);
    };

    const setFavorite = async () => {
        await dispatch(likeAndDislikePost(new PostLikedAction(post.postId, !isFavoritePost)));
        setIsFavoritePost(!isFavoritePost);
    };

    useEffect(() => {
        setIsFavoritePost(post.isLiked);
    }, [post]);

    return (
        <>
            <div className="post">
                <Link to={`/user/${post.creatorProfile.creatorId}`} className="post__user">
                    <Avatar size={AVATAR_SIZE} photo={`${window.location.origin}/images/users/${post.creatorProfile.creatorId}.png`} isAvatarExists={post.creatorProfile.isAvatarExists} />
                    <p className="post__user__text">
                        {post.creatorProfile.username}
                    </p>
                </Link>
                <PostSlider postId={post.postId} numOfImages={post.num_of_images} />
                <div className="post__bottom-side">
                    <button className="post__button" onClick={() => setFavorite()}>
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
                    {post.num_of_likes} {post.num_of_likes === ONE_LIKE ? 'вподобайка' : 'вподобайки'}
                </div>
                {post.description &&
                    <div className="post__description">
                        <Link
                            to={`/user/${post.creatorProfile.creatorId}`}
                            className="post__description__name">
                            {post.creatorProfile.username}
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
