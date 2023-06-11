import closeIcon from '@static/img/User/Post/closeIcon.png';

import './index.scss';

export const Modal: React.FC<{
    children: JSX.Element;
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
    classname?: string;
}> = ({ children, setIsOpenModal, classname }) => <div className={`modal ${classname ? classname : ''}`} onClick={() => setIsOpenModal(false)}>
    <div className="modal__content"
        onClick={e => {
            e.stopPropagation();
        }}>
        <div className="modal__close" onClick={() => setIsOpenModal(false)} >
            <img className="modal__close__icon" src={closeIcon} alt="close" />
        </div>
        <div className="modal__content__line" />
        {children}
    </div>
</div>;
