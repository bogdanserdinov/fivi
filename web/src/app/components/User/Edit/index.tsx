import { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Avatar } from '@components/common/Avatar';
import backButtonIcon from '@img/backButtonIcon.png';
import noUserPhoto from '@static/img/User/no-photo-profile.webp';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';

import { RootState } from '@/app/store';
import { User, UserUpdate } from '@/users';
import { getUserProfile, updateUser } from '@/app/store/actions/users';
import { convertToBase64 } from '@/app/internal/convertImage';
import { removeLocalStorageItem } from '@/app/utils/localStorage';


import './index.scss';
import { setCurrentId } from '@/app/store/reducers/posts';
import { getPostsProfile } from '@/app/store/actions/posts';

const AVATAR_SIZE = 160;
const PHOTO_INDEX = 0;
const SECOND_INDEX = 1;

export const UserEdit = () => {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();
    const user: User = useAppSelector((state: RootState) => state.usersReducer.user);

    const [photo, setPhoto] = useState<string>('');
    const [nickName, setNickName] = useState(user.username);
    const [file, setFile] = useState<string>('');
    const [email, setEmail] = useState<string>('');

    const handleFileChange = async(e: any) => {
        if (e.target.files?.length) {
            setPhoto(URL.createObjectURL(e.target.files[PHOTO_INDEX]));
            const convertedFile: string = await convertToBase64(e.target.files[PHOTO_INDEX]);
            setFile(convertedFile.split(',')[SECOND_INDEX]);
        }
    };

    const handleNicknameChange = (e: any) => {
        setNickName(e.target.value);
    };

    const handleEmailChange = (e: any) => {
        setEmail(e.target.value);
    };

    const sendChanges = async() => {
        await dispatch(updateUser(new UserUpdate(
            '',
            nickName,
            email,
            file
        )));
        navigate(`/user/${user.userId}`);
    };

    const setAvatar = () => {
        user.isAvatarExists ? setPhoto(`${window.location.origin}/images/users/${user.userId}.png`) : setPhoto(noUserPhoto);
    };

    const logoutUser = () => {
        removeLocalStorageItem('AUTH_TOKEN');
        navigate('/registration');
    };

    const handleNavigateToUser = () => {
        dispatch(getUserProfile(user.userId));
        dispatch(setCurrentId(user.userId));
        dispatch(getPostsProfile(user.userId));
        navigate(`/user/${user.userId}`);
    };

    useEffect(() => {
        setAvatar();
        setNickName(user.username);
        setEmail(user.email);
    }, [user]);

    return (
        <div className="user-edit">
            <div className="user-edit__top-side">
                <div onClick={() => handleNavigateToUser()} className="user-edit__back-button">
                    <img src={backButtonIcon} alt="back button" className="user-edit__back-button__image" />
                </div>

                <button
                    className="user-edit__logout"
                    type="button"
                    onClick={() => logoutUser()}
                >
                    Вийти
                </button>
            </div>

            <div className="user-edit__profile">
                <label htmlFor="user-edit" className="user-edit__file-label">
                    <Avatar size={AVATAR_SIZE} urlPhoto={photo} isAvatarExists={true} />

                    <p className="user-edit__file-label__text">Змінити фото</p>
                </label>
                <input
                    type="file"
                    accept="image/png, image/jpeg"
                    onChange={handleFileChange}
                    id="user-edit"
                    className="user-edit__file-input"
                />
            </div>
            <div className="user-edit__field">
                <label className="user-edit__field__label">Змінити нік</label>
                <input className="user-edit__field__input" type="text" value={nickName} onChange={handleNicknameChange} />
            </div>
            <div className="user-edit__field">
                <label className="user-edit__field__label">Змінити пошту</label>
                <input className="user-edit__field__input" type="text" value={email} onChange={handleEmailChange} />
            </div>
            <button
                className="user-edit__save-changes"
                onClick={() => sendChanges()}>
                Примінити зміни
            </button>
        </div>
    );
};
