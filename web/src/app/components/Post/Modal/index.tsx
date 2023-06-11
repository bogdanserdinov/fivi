import { Link, useNavigate } from 'react-router-dom';
import { useEffect, useState } from 'react';

import { Avatar } from '@components/common/Avatar';
import { CommentPostModal } from '@components/Post/Modal/Comment';
import { PostSlider } from '@components/Post/Slider';
import isFavoriteIcon from '@static/img/post/isFavoriteIcon.png';
import favoriteIcon from '@static/img/post/favoriteIcon.png';
import closeIcon from '@static/img/User/Post/closeIcon.png';
import moreIcon from '@static/img/User/Post/moreIcon.png';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { SettingsModal } from '@/app/components/Post/Modal/SettingsModal';
import { EditingModal } from '@/app/components/Post/Modal/EditingModal';
import { Comment, CommentCreate, Post, PostLikedAction } from '@/post';
import { getPostsProfile, likeAndDislikePost, sendComment } from '@/app/store/actions/posts';
import { User } from '@/users';
import { RootState } from '@/app/store';
import { getUserProfile } from '@/app/store/actions/users';
import { setCurrentId } from '@/app/store/reducers/posts';


import './index.scss';

const ONE_PERSON_LIKE = 1;
const AVATAR_SIZE = 50;
const NO_DESCRIPTION = 0;

export const PostModal: React.FC<{
    post: Post;
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
}> = ({ post, setIsOpenModal }) => {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const [currentComment, setCurrentComment] = useState<string>('');
    const [isModalSettings, setIsModalSettings] = useState<boolean>(false);
    const [isModalEditing, setIsModalEditing] = useState<boolean>(false);
    const [isFavoritePost, setIsFavoritePost] = useState<boolean>(false);

    const user: User = useAppSelector((state: RootState) => state.usersReducer.user);

    const onClickOutOfModal = (e: any) => { };

    const showUserPhotoLikes = () => { };

    const onChangeAddComment = (e: any) => {
        setCurrentComment(e.target.value);
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

    const setFavorite = async() => {
        await dispatch(likeAndDislikePost(new PostLikedAction(post.postId, !isFavoritePost)));
        setIsFavoritePost(!isFavoritePost);
    };

    const handleNavigateToUser = async() => {
        await dispatch(getUserProfile(post.creatorId));
        await dispatch(setCurrentId(post.creatorId));
        await dispatch(getPostsProfile(post.creatorId));
        navigate(`/user/${post.creatorId}`);
    };


    useEffect(() => {
        setIsFavoritePost(post.isLiked);
    }, [post]);

    return (
        <div className="post-modal" id="modal" onClick={onClickOutOfModal}>
            <div className="post-modal__content">
                <div className="post-modal__photos">
                    <PostSlider postId={post.postId} numOfImages={post.numOfImages} />
                </div>
                <div className="post-modal__info">
                    <div className="post-modal__top-side" >
                        <div onClick={() => handleNavigateToUser()} className="post-modal__user-info" >
                            <Avatar size={AVATAR_SIZE} userId={post.creatorProfile.creatorId} isAvatarExists={post.creatorProfile.isAvatarExists} />
                            <p className="post-modal__user-info__nickname">
                                {post.creatorProfile.username}
                            </p>
                        </div>
                        {post.creatorId === user.userId
                            ?
                            <div className="post-modal__buttons">

                                <button
                                    className="post-modal__button"
                                    onClick={() => setIsModalSettings(true)}
                                >
                                    <img src={moreIcon} alt="more icon" className="post-modal__button__more-image" />
                                </button>
                                <button
                                    className="post-modal__button"
                                    onClick={() => setIsOpenModal(false)}
                                >
                                    <img src={closeIcon} alt="close icon" className="post-modal__button__image" />
                                </button>
                                {isModalSettings && <SettingsModal
                                    setIsModalEditing={setIsModalEditing}
                                    setIsModalSettings={setIsModalSettings}
                                    setIsOpenModal={setIsOpenModal}
                                    postId={post.postId} />}
                            </div>
                            :
                            <button
                                className="post-modal__button"
                                onClick={() => setIsOpenModal(false)}
                            >
                                <img src={closeIcon} alt="close icon" className="post-modal__button__image" />
                            </button>
                        }

                    </div>
                    <div className="post-modal__comments">
                        {post.description.length !== NO_DESCRIPTION &&
                            <CommentPostModal
                                isAvatarExists={post.creatorProfile.isAvatarExists}
                                nickname={post.creatorProfile.username}
                                text={post.description}
                                commentorId={post.creatorProfile.creatorId}
                                setIsOpenModal={setIsOpenModal}
                            />
                        }
                        {post.comments.map((comment: Comment) =>
                            <CommentPostModal
                                isAvatarExists={comment.isAvatarExists}
                                nickname={comment.username}
                                text={comment.text}
                                commentorId={comment.userId}
                                key={comment.commentId}
                                setIsOpenModal={setIsOpenModal}
                            />)
                        }
                    </div>
                    <div className="post-modal__bottom-side">
                        <div className="post-modal__actions">
                            <button className="post-modal__button" onClick={() => setFavorite()}>
                                {isFavoritePost ?
                                    <img
                                        className="post-modal__button__image"
                                        src={isFavoriteIcon}
                                        alt="favorite" />
                                    :
                                    <img
                                        className="post-modal__button__image"
                                        src={favoriteIcon}
                                        alt="favorite" />
                                }

                            </button>

                        </div>
                        {post.numOfLikes ?
                            <div className="post-modal__likes" onClick={() => showUserPhotoLikes()}>
                                Подобається {post.numOfLikes} {post.numOfLikes === ONE_PERSON_LIKE ? 'людині' : 'людям'}
                            </div>
                            :
                            <p className="post-modal__likes">Нікому ще не сподобалось</p>
                        }
                        <div className="post-modal__add-comment">
                            <input
                                className="post-modal__add-comment__input"
                                type="text"
                                placeholder="Додайте коментарій..."
                                onChange={onChangeAddComment}
                                value={currentComment}
                                onKeyDown={onKeyDownInput}
                            />
                            {currentComment &&
                                <button className="post-modal__add-comment__send"
                                    onClick={() => sendCommentText()}
                                >
                                    Відправити
                                </button>
                            }
                        </div>
                    </div>
                </div>
            </div>
            {isModalEditing && <EditingModal setIsOpenModal={setIsModalEditing} post={post} />}
        </div >
    );
};
