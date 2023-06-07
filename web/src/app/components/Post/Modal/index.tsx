import { Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { useAppDispatch } from '@/app/hooks/useReduxToolkit';

import { Avatar } from '@components/common/Avatar';
import { CommentPostModal } from '@components/Post/Modal/Comment';
import { PostSlider } from '@components/Post/Slider';

import isFavoriteIcon from '@static/img/post/isFavoriteIcon.png';
import { SettingsModal } from '@/app/components/Post/Modal/SettingsModal';
import { EditingModal } from '@/app/components/Post/Modal/EditingModal';
import { Comment, CommentCreate, Post, PostLikedAction } from '@/post';
import { likeAndDislikePost, sendComment } from '@/app/store/actions/posts';

import favoriteIcon from '@static/img/post/favoriteIcon.png';
import closeIcon from '@static/img/User/Post/closeIcon.png';
import moreIcon from '@static/img/User/Post/moreIcon.png';

import './index.scss';

const ONE_PERSON_LIKE = 1;
const AVATAR_SIZE = 50;

export const PostModal: React.FC<{
    post: Post;
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
}> = ({ post, setIsOpenModal }) => {
    const dispatch = useAppDispatch();

    const [currentComment, setCurrentComment] = useState<string>('');
    const [isModalSettings, setIsModalSettings] = useState<boolean>(false);
    const [isModalEditing, setIsModalEditing] = useState<boolean>(false);
    const [isFavoritePost, setIsFavoritePost] = useState<boolean>(false);

    const onClickOutOfModal = (e: any) => { };

    const showUserPhotoLikes = () => { };

    const onChangeAddComment = (e: any) => {
        setCurrentComment(e.target.value);
    };

    const setFavorite = async () => {
        await dispatch(likeAndDislikePost(new PostLikedAction(post.postId, !isFavoritePost)));
        setIsFavoritePost(!isFavoritePost);
    };

    const sendCommentText = () => {
        dispatch(sendComment(new CommentCreate(post.postId, currentComment)));
    };


    useEffect(() => {
        setIsFavoritePost(post.isLiked);
    }, [post]);

    return (
        <div className="post-modal" id="modal" onClick={onClickOutOfModal}>
            <div className="post-modal__content">
                <div className="post-modal__photos">
                    <PostSlider postId={post.postId} numOfImages={post.num_of_images} />
                </div>
                <div className="post-modal__info">
                    <div className="post-modal__top-side" >
                        <Link to={`user/${post.creatorId}`} className="post-modal__user-info" >
                            <Avatar size={AVATAR_SIZE} photo={`${window.location.origin}/images/users/${post.creatorProfile.creatorId}.png`} isAvatarExists={post.creatorProfile.isAvatarExists} />
                            <p className="post-modal__user-info__nickname">
                                {post.creatorProfile.username}
                            </p>
                        </Link>
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
                            {isModalSettings && <SettingsModal setIsModalEditing={setIsModalEditing} setIsOpenModal={setIsModalSettings} postId={post.postId} />}
                        </div>
                    </div>
                    <div className="post-modal__comments">
                        {post.description.length !== 0 &&
                            <CommentPostModal
                                isAvatarExists={post.creatorProfile.isAvatarExists}
                                nickname={post.creatorProfile.username}
                                text={post.description}
                                commentorId={post.creatorProfile.creatorId}
                            />
                        }
                        {post.comments.map((comment: Comment) =>
                            <CommentPostModal
                                // TODO : add real data
                                isAvatarExists={true}
                                nickname={comment.username}
                                text={comment.text}
                                commentorId={comment.userId}
                                key={comment.username}
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
                        {post.num_of_likes ?
                            <div className="post-modal__likes" onClick={() => showUserPhotoLikes()}>
                                Подобається {post.num_of_likes} {post.num_of_likes === ONE_PERSON_LIKE ? 'людині' : 'людям'}
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
                            />
                            {currentComment &&
                                <button className="post-modal__add-comment__send"
                                    onClick={() => sendCommentText()}>
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
