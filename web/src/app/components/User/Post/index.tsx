import { useState } from 'react';

import { PostModal } from '@components/Post/Modal';
import { PostPhoto } from '@components/common/PostPhoto';

import favoriteIcon from '@static/img/User/Post/favoriteIcon.png';
import commentIcon from '@static/img/User/Post/commentIcon.png';
import galleryIcon from '@static/img/User/Post/galleryIcon.png';
import { setCurrentPost } from '@/app/store/reducers/posts';
import { useAppDispatch } from '@/app/hooks/useReduxToolkit';
import { Post } from '@/post';

import './index.scss';

const ONE_PHOTO = 1;

export const UserPost: React.FC<{ post: Post }> = ({ post }) => {
    const [isOpenModal, setIsOpenModal] = useState(false);
    const dispatch = useAppDispatch();

    const handleModalOpening = () => {
        dispatch(setCurrentPost(post));
        setIsOpenModal(true);
    };

    return (
        <>
            <div className="user-post" onClick={() => handleModalOpening()}>
                {post.numOfImages > ONE_PHOTO &&
                    <img
                        src={galleryIcon}
                        alt="post"
                        className="user-post__lots-photos-icon"
                    />

                }
                {post.numOfImages ?
                    <PostPhoto postId={post.postId} postIndex={0} width={300} height={400} isPostPhotoExist={true} />
                    :
                    <PostPhoto isPostPhotoExist={false} width={300} height={400} />
                }
                <div className="user-post--hovered">
                    <p className="user-post--hovered__info">
                        <img
                            className="user-post--hovered__info__image"
                            src={favoriteIcon}
                            alt="favorite" />
                        {post.numOfLikes}
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
