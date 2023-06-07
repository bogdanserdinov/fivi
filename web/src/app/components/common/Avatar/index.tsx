import noUserPhoto from '@static/img/User/no-photo-profile.webp';

import './index.scss';

export const Avatar: React.FC<{ size: number; photo: string; isAvatarExists: boolean }> = ({ size, photo, isAvatarExists }) =>
    <>
        {
            isAvatarExists ?
                <div
                    className="avatar"
                    style={{ backgroundImage: `url(${photo})`, width: `${size}px`, height: `${size}px` }}
                />
                :
                <div
                    className="avatar"
                    style={{ backgroundImage: `url(${noUserPhoto})`, width: `${size}px`, height: `${size}px` }}
                />

        }
    </>;


