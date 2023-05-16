import './index.scss'
import '../index.scss'
import { Link } from 'react-router-dom'

import copyIcon from '@img/Auth/copy.png'
const MOCKED__PHRASES = [
    'tool', 'tell', 'boy', 'mom', 'ball', 'rate', 'strike', 'lady', 'trial', 'banner', 'oppose', 'uphold'
]

export const RegistrationPage = () => {

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
                {MOCKED__PHRASES.map((phrase, index) =>
                    <div className="registration__phrases__item" key={phrase}>
                        {index}
                        <div className="registration__phrases__item__block">
                            {phrase}
                        </div>
                    </div>
                )}
                 <button className="registration__copy">
                <img src={copyIcon} alt='copy icon' className="registration__copy__icon" /> Copy
            </button>
            </div>
           
            <h2 className="authentication__subtitle">Add your data</h2>
            <form className="authentication__form">
                <input type='text' placeholder='Username' className="authentication__input" />
                <input type='email' placeholder='Email' className="authentication__input" />
            </form>
            <button type="button"className="authentication__submit" >Register</button>
        </div>
    )
}
