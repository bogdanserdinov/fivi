import noUserPhoto from '@static/img/User/no-photo-profile.webp';

import './index.scss';

export const Avatar: React.FC<{ size: number; userId: string; isAvatarExists: boolean }> = ({ size, userId, isAvatarExists }) =>
    <>
        {
            isAvatarExists ?
                <div
                    className="avatar"
                    style={{ backgroundImage: `url(${window.location.origin}/images/users/${userId}.png)`, width: `${size}px`, height: `${size}px` }}
                />
                :
                <div
                    className="avatar"
                    style={{ backgroundImage: `url(${noUserPhoto})`, width: `${size}px`, height: `${size}px` }}
                />

        }
    </>;


