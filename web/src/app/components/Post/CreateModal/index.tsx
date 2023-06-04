import { ChangeEvent, useState } from 'react';

import { Modal } from '@components/common/Modal';

import addPhotoIcon from '@img/post/addPhotoIcon.png';
import closeIcon from '@img/User/Post/closeIcon.png';

import './index.scss';

export const PostCreateModal: React.FC<{
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
}>
    = ({ setIsOpenModal }) => {
        const [files, setFiles] = useState<string[]>();
        const removePhoto = (index: number) => {
            if (files) {
                const galleryImagesData = [...files];
                galleryImagesData?.splice(index, 1);
                setFiles(galleryImagesData);
            }
        };

        const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
            if (e.target.files?.length) {
                setFiles([URL.createObjectURL(e.target.files[0])]);
            }
        };

        return (
            <Modal setIsOpenModal={setIsOpenModal}>

                <form className="create-post-modal__form">
                    <div className="create-post-modal__add-photo">
                        <label
                            className="create-post-modal__add-photo__label"
                            htmlFor="create-post-add-photo">
                            <span className="create-post-modal__add-photo__label__icon">
                                <img
                                    className="create-post-modal__add-photo__label__icon__photo"
                                    src={addPhotoIcon}
                                    alt="add photo" />
                            </span>
                            <span className="create-post-modal__add-photo__label__button">
                                Оберіть фотографію
                            </span>
                        </label>
                        <input
                            className="create-post-modal__add-photo__input"
                            id="create-post-add-photo" type="file" accept="image/png, image/jpeg"
                            onChange={handleFileChange} />
                    </div>
                    <div className="create-post-modal__photos">{
                        files?.length && files.length !== 0 && files.map((file, index) =>
                            <div style={{ backgroundImage: `url(${file}` }}
                                className="create-post-modal__photos__item"
                                key={file}>
                                <button
                                    type="button"
                                    className="create-post-modal__photos__item__close"
                                    onClick={() => removePhoto(index)}
                                >
                                    <img
                                        src={closeIcon}
                                        alt="close"
                                        className="create-post-modal__photos__item__close__icon"
                                    />
                                </button>
                            </div>
                        )

                    }
                    </div>
                    <textarea className="create-post-modal__description" />
                    <button className="create-post-modal__button">
                        Create
                    </button>
                </form>
            </Modal>
        );
    };
