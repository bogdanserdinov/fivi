import noPostPhoto from '@static/img/post/NoImageYet.jpg';

import './index.scss';

export const PostPhoto: React.FC<{
    height: number;
    width: number;
    postId?: string;
    postIndex?: number;
    urlPhoto?: string;
    isPostPhotoExist?: boolean;
}>
    = ({ height,
        width,
        postId,
        postIndex,
        isPostPhotoExist,
        urlPhoto,
    }) =>
        <>
            {
                isPostPhotoExist ?
                    <div
                        className="post-photo"
                        style={{
                            backgroundImage: `url(${postId ?
                                `${window.location.origin}/images/posts/${postId}/${postIndex}.png`
                                : urlPhoto
                            })`,
                            width: `${width}px`,
                            height: `${height}px`,
                        }}
                    />
                    :
                    <div
                        className="post-photo"
                        style={{ backgroundImage: `url(${noPostPhoto})`, width: `${width}px`, height: `${height}px` }}
                    />
            }
        </>;


