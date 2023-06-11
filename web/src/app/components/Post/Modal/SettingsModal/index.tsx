import { Modal } from '@components/common/Modal';

import deleteIcon from '@img/User/Post/Settings/deleteIcon.png';
import editIcon from '@img/User/Post/Settings/editIcon.png';
import { deletePost } from '@/app/store/actions/posts';
import { useAppDispatch } from '@/app/hooks/useReduxToolkit';


import './index.scss';

export const SettingsModal: React.FC<{
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
    setIsModalEditing: React.Dispatch<React.SetStateAction<boolean>>;
    setIsModalSettings: React.Dispatch<React.SetStateAction<boolean>>;
    postId: string;
}> = ({ setIsOpenModal, setIsModalEditing, postId, setIsModalSettings }) => {
    const dispatch = useAppDispatch();
    const deletePostss = () => {
        dispatch(deletePost(postId));
        setIsOpenModal(false);
    };

    const editPost = () => {
        setIsModalSettings(false);
        setIsModalEditing(true);
    };

    return (
        <Modal classname="settings-modal" setIsOpenModal={setIsOpenModal}>
            <div>
                <button className="settings-modal__button" onClick={() => editPost()}
                    type="button">
                    <img className="settings-modal__button__image" src={editIcon} alt="edit icon" />
                    Редагувати
                </button>
                <button className="settings-modal__button" type="button"
                    onClick={() => deletePostss()}>
                    <img className="settings-modal__button__image" src={deleteIcon} alt="delete icon" />
                    Видалити
                </button>
            </div>
        </Modal>
    );
};
