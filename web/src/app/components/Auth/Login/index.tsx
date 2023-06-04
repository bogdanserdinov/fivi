import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';

import { useAppDispatch } from '@/app/hooks/useReduxToolkit';
import { RoutesConfig } from '@/app/routes';
import { login } from '@/app/store/actions/users';
import { UserLoginData } from '@/users';

import '../index.scss';
import './index.scss';

export const LoginPage = () => {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const [username, setUsername] = useState<string>('');
    const [mnemonicPhrases, setMnemonicPhrases] = useState<string[]>();

    const convertMnemonicPhrases = (phrases: string) => {
        setMnemonicPhrases(phrases.split(','));
    };

    const loginUser = async() => {
        await dispatch(login(new UserLoginData(username, mnemonicPhrases)));

        await window.localStorage.setItem('IS_LOGGEDIN', JSON.stringify(true));

        navigate(RoutesConfig.Home.path);
    };

    return (
        <div className="login">
            <div className="authentication__switcher">
                <Link className="authentication__switcher__item authentication__switcher__item--active" to="/login">
                    Увійти
                </Link>
                <Link className="authentication__switcher__item" to="/registration">Реєстрація</Link>
            </div>
            <h1 className="authentication__title">Login</h1>
            <h2 className="authentication__subtitle">Add your data</h2>
            <form className="authentication__form">
                <input
                    type="text"
                    placeholder="Username"
                    className="authentication__input"
                    onChange={e => setUsername(e.target.value)}
                />
                <textarea
                    className="login__textarea"
                    placeholder="Add secret phrase"
                    onChange={e => convertMnemonicPhrases(e.target.value)} />
            </form>
            <button
                type="button"
                className="authentication__submit"
                onClick={() => loginUser()}
            >Login</button>
        </div>);
};
