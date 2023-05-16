import { useState } from 'react';

import favoriteIcon from '@static/img/User/Post/favoriteIcon.png';
import commentIcon from '@static/img/User/Post/commentIcon.png';
import galleryIcon from '@static/img/User/Post/galleryIcon.png';

import { PostModal } from '@components/Post/Modal';

import './index.scss';

const ONE_PHOTO = 1;
const FIRST_PHOTO = 0;

export const UserPost: React.FC<{ post: any }> = ({ post }) => {
    const [isOpenModal, setIsOpenModal] = useState(false);

    return (
        <>
            <div className="user-post" onClick={() => setIsOpenModal(true)}>
                {post.photos.length > ONE_PHOTO &&
                    <img
                        src={galleryIcon}
                        alt="post"
                        className="user-post__lots-photos-icon"
                    />
                }
                <img src={post.photos[FIRST_PHOTO]} alt="post" className="user-post__image" />
                <div className="user-post--hovered">
                    <p className="user-post--hovered__info">
                        <img
                            className="user-post--hovered__info__image"
                            src={favoriteIcon}
                            alt="favorite" />
                        {post.favorites}
                    </p>
                    <p className="user-post--hovered__info">
                        <img
                            src={commentIcon}
                            alt="comment"
                            className="user-post--hovered__info__image"
                        />
                        {post.comments.length}
                    </p>
                </div>
            </div>
            {isOpenModal && <PostModal post={post} setIsOpenModal={setIsOpenModal} />}
        </>
    );
};
