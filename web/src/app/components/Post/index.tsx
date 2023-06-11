import { useNavigate } from 'react-router-dom';
import { useEffect, useState } from 'react';

import { PostModal } from '@components/Post/Modal';
import { Avatar } from '@components/common/Avatar';
import { PostPhoto } from '@components/common/PostPhoto';
import commentIcon from '@static/img/post/commentIcon.png';
import favoriteIcon from '@static/img/post/favoriteIcon.png';
import isFavoriteIcon from '@static/img/post/isFavoriteIcon.png';
import { PostSlider } from './Slider';
import { CommentCreate, Post, PostLikedAction } from '@/post';
import { getPostsProfile, likeAndDislikePost, sendComment } from '@/app/store/actions/posts';
import { setCurrentId, setCurrentPost } from '@/app/store/reducers/posts';
import { useAppDispatch } from '@/app/hooks/useReduxToolkit';
import { getUserProfile } from '@/app/store/actions/users';


import './index.scss';

const ONE_LIKE = 1;
const AVATAR_SIZE = 40;

export const PostPage: React.FC<{ post: Post }> = ({ post }) => {
    const [isOpenModal, setIsOpenModal] = useState<boolean>(false);
    const [isFavoritePost, setIsFavoritePost] = useState<boolean>(false);
    const [currentComment, setCurrentComment] = useState<string>('');
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const onChangeAddComment = (e: any) => {
        setCurrentComment(e.target.value);
    };

    const handleNavigateToUser = () => {
        dispatch(getUserProfile(post.creatorProfile.creatorId));
        dispatch(setCurrentId(post.creatorProfile.creatorId));
        dispatch(getPostsProfile(post.creatorProfile.creatorId));
        navigate(`/user/${post.creatorProfile.creatorId}`);
    };

    const handleModalOpening = () => {
        dispatch(setCurrentPost(post));
        setIsOpenModal(true);
    };

    const setFavorite = async() => {
        await dispatch(likeAndDislikePost(new PostLikedAction(post.postId, !isFavoritePost)));
        setIsFavoritePost(!isFavoritePost);
    };

    const sendCommentText = () => {
        dispatch(sendComment(new CommentCreate(post.postId, currentComment)));
        setCurrentComment('');
    };

    const onKeyDownInput = (e: any) => {
        if (e.key === 'Enter') {
            sendCommentText();
        }
    };

    useEffect(() => {
        setIsFavoritePost(post.isLiked);
    }, [post]);

    return (
        <>
            <div className="post">
                <div onClick={() => handleNavigateToUser()} className="post__user">
                    <Avatar size={AVATAR_SIZE} userId={post.creatorProfile.creatorId} isAvatarExists={post.creatorProfile.isAvatarExists} />
                    <p className="post__user__text">
                        {post.creatorProfile.username}
                    </p>
                </div>
                {post.numOfImages ?
                    < PostSlider postId={post.postId} numOfImages={post.numOfImages} />
                    :
                    <PostPhoto isPostPhotoExist={false} width={400} height={500} />
                }
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
                    {post.numOfLikes} {post.numOfLikes === ONE_LIKE ? 'вподобайка' : 'вподобайки'}
                </div>
                {post.description &&
                    <div className="post__description">
                        <div
                            onClick={() => handleNavigateToUser()}
                            className="post__description__name">
                            {post.creatorProfile.username}
                        </div>
                        <p className="post__description__text">
                            {post.description}
                        </p>
                    </div>
                }
                <div className="post__show-all-comments"
                    onClick={() => handleModalOpening()}>
                    показати все коментарі ({post.comments.length})
                </div>
                <div className="post__add-comment">
                    <input
                        className="post__add-comment__input"
                        type="text"
                        placeholder="Додайте коментарій..."
                        onChange={onChangeAddComment}
                        value={currentComment}
                        onKeyDown={onKeyDownInput}
                    />
                    {currentComment &&
                        <button className="post__add-comment__send" onClick={() => sendCommentText()}>
                            Відправити
                        </button>
                    }
                </div>
            </div>
            {isOpenModal && <PostModal post={post} setIsOpenModal={setIsOpenModal} />}
        </>
    );
};
