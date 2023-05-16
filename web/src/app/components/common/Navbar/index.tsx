import { Link, NavLink } from 'react-router-dom';

import searchIcon from '@img/Navbar/searchIcon.png';
import settingsIcon from '@img/Navbar/settingsIcon.png';
import addPostIcon from '@img/Navbar/addPostIcon.png';

import { Avatar } from '@components/common/Avatar';
import { PostCreateModal } from '@components/Post/CreateModal';

import { user } from '@/mocked/user';

import './index.scss';
import { useState } from 'react';

export const Navbar = () => {
      const [isOpenModal,setIsOpenModal]=useState(false)
    return(
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
                    <div  className="header__icon" onClick={()=>setIsOpenModal(true)}>
                        <img src={addPostIcon}
                            alt="favorite"
                            className="header__icon__image" />
                    </div>
                    <NavLink to="/settings" className="header__icon">
                        <img src={settingsIcon}
                            alt="settings"
                            className="header__icon__image" />
                    </NavLink>
                    <NavLink className="header__user" to={`/user/${user.id}`}>
                        <Avatar size={40} photo={user.avatar} />
                        <p className="header__user__text">{user.nickname}</p>
                    </NavLink>
                </div>
            </div>
        </header>
        {isOpenModal && <PostCreateModal setIsOpenModal={setIsOpenModal} />}
    </>)
}
