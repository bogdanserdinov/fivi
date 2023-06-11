import { useEffect, useState } from 'react';
import { NavLink, useNavigate } from 'react-router-dom';

import { Avatar } from '@components/common/Avatar';
import { PostCreateModal } from '@components/Post/CreateModal';
import searchIcon from '@img/Navbar/searchIcon.png';
import settingsIcon from '@img/Navbar/settingsIcon.png';
import addPostIcon from '@img/Navbar/addPostIcon.png';
import { SearchingModal } from './SearchingModal';
import { User, UserProfile } from '@/users';
import { RootState } from '@/app/store';
import { getUser, getUserProfile, searchUsers } from '@/app/store/actions/users';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { getPostsProfile } from '@/app/store/actions/posts';
import { setCurrentId } from '@/app/store/reducers/posts';


import './index.scss';

const AVATAR_SIZE = 50;
export const Navbar = () => {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();
    const [isOpenModal, setIsOpenModal] = useState(false);
    const [isSearching, setIsSearching] = useState(false);

    const user: User = useAppSelector((state: RootState) => state.usersReducer.user);
    const foundedUsers: UserProfile[] | [] = useAppSelector((state: RootState) => state.usersReducer.foundedUsers);

    const handleSearching = (e: any) => {
        if (e.target.value) {
            setIsSearching(true);
            dispatch(searchUsers(e.target.value));
        }
        else {
            setIsSearching(false);
        }
    };

    const handleRedirectingUserPage = () => {
        navigate(`/user/${user.userId}`);
        dispatch(getUserProfile(user.userId));
        dispatch(getPostsProfile(user.userId));
        dispatch(setCurrentId(user.userId));
    };

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
                            id="search"
                            type="text"
                            placeholder="Я шукаю..."
                            onChange={handleSearching}
                            autoComplete="off"
                        />
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
                        <div className="header__user" onClick={() => handleRedirectingUserPage()}>
                            <Avatar size={AVATAR_SIZE} userId={user.userId} isAvatarExists={user.isAvatarExists} />
                            <p className="header__user__text">{user.username}</p>
                        </div>
                    </div>
                </div>
                {isOpenModal && <PostCreateModal setIsOpenModal={setIsOpenModal} />}
            </header>
            {isSearching && <SearchingModal setIsSearching={setIsSearching} foundedUsers={foundedUsers} />}
        </>);
};
