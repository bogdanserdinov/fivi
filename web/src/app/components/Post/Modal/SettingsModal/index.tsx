import { Modal } from "@components/common/Modal"

import deleteIcon from "@img/User/Post/Settings/deleteIcon.png"
import editIcon from "@img/User/Post/Settings/editIcon.png"

import './index.scss'

export const SettingsModal: React.FC<{
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
    setIsModalEditing: React.Dispatch<React.SetStateAction<boolean>>
}> = ({ setIsOpenModal, setIsModalEditing }) => {
    const deletePost = () => { }

    const editPost = () => {
        setIsOpenModal(false)
        setIsModalEditing(true);
    }
    return (
        <Modal classname="settings-modal" setIsOpenModal={setIsOpenModal}>
            <div>
                <button className="settings-modal__button" onClick={()=>editPost()}>
                    <img  className="settings-modal__button__image" src={editIcon} alt="edit icon" />
                    Редагувати
                </button>
                <button className="settings-modal__button">
                    <img className="settings-modal__button__image" src={deleteIcon} alt="delete icon" />
                    Видалити
                </button>
            </div>
        </Modal>
    )
}