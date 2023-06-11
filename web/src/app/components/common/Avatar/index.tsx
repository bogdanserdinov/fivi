import noUserPhoto from '@static/img/User/no-photo-profile.webp';

import './index.scss';

export const Avatar: React.FC<{ size: number; userId?: string; isAvatarExists?: boolean; urlPhoto?: string }> = ({ size, userId, isAvatarExists, urlPhoto }) =>
    <>
        {
            isAvatarExists ?
                <div
                    className="avatar"

                    style={{
                        backgroundImage: `url(${userId ?
                            `${window.location.origin}/images/users/${userId}.png`
                            : urlPhoto
                        })`, width: `${size}px`, height: `${size}px`,
                    }}
                />
                :
                <div
                    className="avatar"
                    style={{ backgroundImage: `url(${noUserPhoto})`, width: `${size}px`, height: `${size}px` }}
                />

        }
    </>;


