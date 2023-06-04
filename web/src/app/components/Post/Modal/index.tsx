import { Link } from 'react-router-dom';
import { useState } from 'react';

import commentIcon from '@static/img/post/commentIcon.png';
import favoriteIcon from '@static/img/post/favoriteIcon.png';
import closeIcon from '@static/img/User/Post/closeIcon.png';
import moreIcon from '@static/img/User/Post/moreIcon.png';
import isFavoriteIcon from '@static/img/post/isFavoriteIcon.png';

import { Avatar } from '@components/common/Avatar';
import { CommentPostModal } from '@components/Post/Modal/Comment';
import { PostSlider } from '@components/Post/Slider';

import './index.scss';
import { SettingsModal } from './SettingsModal';
import { EditingModal } from './EditingModal';

const ONE_PERSON_LIKE = 1;
const AVATAR_SIZE = 50;

export const PostModal: React.FC<{
    post: any;
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
}> = ({ post, setIsOpenModal }) => {
    const [currentComment, setCurrentComment] = useState<string>('');
    const [isModalSettings, setIsModalSettings] = useState<boolean>(false);
    const [isModalEditing, setIsModalEditing] = useState<boolean>(false);
    const [isFavoritePost, setIsFavoritePost] = useState<boolean>(false);

    const onClickOutOfModal = (e: any) => { };
    const showUserPhotoLikes = () => { };
    const onChangeAddComment = (e: any) => {
        setCurrentComment(e.target.value);
    };

    return (
        <div className="post-modal" id="modal" onClick={onClickOutOfModal}>
            <div className="post-modal__content">
                <div className="post-modal__photos">
                    <PostSlider sliderImages={post.photos} />
                </div>
                <div className="post-modal__info">
                    <div className="post-modal__top-side" >
                        <Link to={`user/${post.creator.id}`} className="post-modal__user-info" >
                            <Avatar size={AVATAR_SIZE} photo={post.creator.creatorAvatar} />
                            <p className="post-modal__user-info__nickname">
                                {post.creator.name}
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
                            {isModalSettings && <SettingsModal setIsModalEditing={setIsModalEditing} setIsOpenModal={setIsModalSettings} />}
                        </div>
                    </div>
                    <div className="post-modal__comments">
                        {post.description &&
                            <CommentPostModal
                                photo={post.creator.creatorAvatar}
                                nickname={post.creator.name}
                                text={post.description}
                                commentorId={post.creator.id}
                            />
                        }
                        {post.comments.map((comment: any) =>
                            <CommentPostModal
                                photo={comment.commentorAvatar}
                                nickname={comment.commentorName}
                                text={comment.commentText}
                                commentorId={comment.commentorId}
                            />)
                        }
                    </div>
                    <div className="post-modal__bottom-side">
                        <div className="post-modal__actions">
                            <button className="post-modal__button" onClick={() => setIsFavoritePost(!isFavoritePost)}>
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
                        {post.favorites ?
                            <div className="post-modal__likes" onClick={() => showUserPhotoLikes()}>
                                Подобається {post.favorites} {post.favorites === ONE_PERSON_LIKE ? 'людині' : 'людям'}
                            </div>
                            :
                            <p className="post-modal__no-likes">Нікому ще не сподобалось</p>
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
                                <button className="post-modal__add-comment__send">
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
