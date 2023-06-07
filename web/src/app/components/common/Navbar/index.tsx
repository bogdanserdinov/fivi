import { useEffect, useState } from 'react';
import { NavLink } from 'react-router-dom';

import { Avatar } from '@components/common/Avatar';
import { PostCreateModal } from '@components/Post/CreateModal';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { getUser } from '@/app/store/actions/users';
import { RootState } from '@/app/store';
import { User } from '@/users';

import searchIcon from '@img/Navbar/searchIcon.png';
import settingsIcon from '@img/Navbar/settingsIcon.png';
import addPostIcon from '@img/Navbar/addPostIcon.png';

import './index.scss';

const AVATAR_SIZE = 50;
export const Navbar = () => {
    const dispatch = useAppDispatch();
    const [isOpenModal, setIsOpenModal] = useState(false);

    const user: User = useAppSelector((state: RootState) => state.usersReducer.user);

    useEffect(() => {
        dispatch(getUser());
    }, []);

    return (
        <>
            <header className="header">
                <div className="header__content">
                    <NavLink className="header__logo" to="/">
                        fivi
                    </NavLink>
                    <div className="header__search">
                        <label htmlFor="search" className="header__search__label">
                            <img className="header__search__icon" src={searchIcon} alt="search icon" />
                        </label>
                        <input className="header__search__input"
                            id="search" type="text" placeholder="Я шукаю..." />
                        <button className="header__search__button" >
                            Знайти
                        </button>
                    </div>
                    <div className="header__navigation">
                        <div className="header__icon" onClick={() => setIsOpenModal(true)}>
                            <img src={addPostIcon}
                                alt="favorite"
                                className="header__icon__image" />
                        </div>
                        <NavLink to={`/user/${user.userId}/edit`} className="header__icon">
                            <img src={settingsIcon}
                                alt="settings"
                                className="header__icon__image" />
                        </NavLink>
                        <NavLink className="header__user" to={`/user/${user.userId}`}>
                            <Avatar size={AVATAR_SIZE} photo={`${window.location.origin}/images/users/${user.userId}.png`} isAvatarExists={user.isAvatarExists} />
                            <p className="header__user__text">{user.username}</p>
                        </NavLink>
                    </div>
                </div>
            </header>
            {isOpenModal && <PostCreateModal setIsOpenModal={setIsOpenModal} />}
        </>);
};
