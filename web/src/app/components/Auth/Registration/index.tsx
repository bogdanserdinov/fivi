import { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';

import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { RootState } from '@/app/store';

import { getMnemonicPhrases, register } from '@/app/store/actions/users';
import { UserRegisterData } from '@/users';
import { RoutesConfig } from '@/app/routes';

import copyIcon from '@img/Auth/copy.png';

import './index.scss';
import '../index.scss';

export const RegistrationPage = () => {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');

    const mnemonicPhrases: string[] | null = useAppSelector((state: RootState) => state.usersReducer.mnemonicPhrases);

    const registerUser = async () => {
        await dispatch(register(new UserRegisterData(email, username, mnemonicPhrases)));

        navigate(RoutesConfig.Home.path);
    };

    useEffect(() => {
        dispatch(getMnemonicPhrases());
    }, []);

    return (
        <div className="registration">
            <div className="authentication__switcher">
                <Link className="authentication__switcher__item " to="/login">
                    Увійти
                </Link>
                <Link className="authentication__switcher__item  authentication__switcher__item--active" to="/registration">Реєстрація</Link>
            </div>
            <h1 className="authentication__title">Registration</h1>

            <h2 className="authentication__subtitle">Secret recovery phrase</h2>
            <div className="registration__phrases">
                {mnemonicPhrases.map((phrase, index) =>
                    <div className="registration__phrases__item" key={phrase}>
                        {index + 1}
                        <div className="registration__phrases__item__block">
                            {phrase}
                        </div>
                    </div>
                )}
                <button className="registration__copy">
                    <img src={copyIcon} alt="copy icon" className="registration__copy__icon" /> Copy
                </button>
            </div>

            <h2 className="authentication__subtitle">Add your data</h2>
            <form className="authentication__form">
                <input
                    type="text"
                    placeholder="Username"
                    className="authentication__input"
                    onChange={e => setUsername(e.target.value)}
                />
                <input
                    type="email"
                    placeholder="Email"
                    className="authentication__input"
                    onChange={e => setEmail(e.target.value)}
                />
            </form>
            <button
                type="button"
                className="authentication__submit"
                onClick={() => registerUser()}
            >
                Register
            </button>
        </div >
    );
};
