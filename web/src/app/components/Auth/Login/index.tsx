import { Link } from 'react-router-dom';
import '../index.scss'
import './index.scss'

export const LoginPage = () =>
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
                <input type='text' placeholder='Username' className="authentication__input" />
                <textarea className="login__textarea"  placeholder='Add secret phrase'/>
            </form>
            <button type="button"className="authentication__submit" >Login</button>
    </div>;
