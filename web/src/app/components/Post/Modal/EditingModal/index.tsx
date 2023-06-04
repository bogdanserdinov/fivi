import { Avatar } from '@components/common/Avatar';
import { Modal } from '@components/common/Modal';
import { useState } from 'react';
import { Link } from 'react-router-dom';

import addPhotoIcon from '@img/post/addPhotoIcon.png';
import closeIcon from '@img/User/Post/closeIcon.png';

import './index.scss';

export const EditingModal: React.FC<{
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
    post: any;
}> = ({ setIsOpenModal, post }) => {
    const [files, setFiles] = useState<string[]>(post.photos);
    const [description, setDescription] = useState(post.description);
    const [currentPhoto, setCurrentPhoto] = useState(post.photos[0]);

    const onChangeDescription = (e: any) => {
        setDescription(e.target.value);
    };

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files?.length) {
            const photos = [...files];
            photos.concat([URL.createObjectURL(e.target.files[0])]);
            setFiles(photos);
        }
    };

    const removePhoto = (e: any, index: number) => {
        e.stopPropagation();
        if (files) {
            const galleryImagesData = [...files];
            galleryImagesData?.splice(index, 1);
            setFiles(galleryImagesData);
        }
    };

    return (
        <Modal classname="editing-modal" setIsOpenModal={setIsOpenModal}>
            <div className="editing-modal__content">
                <div
                    className="editing-modal__photos"
                >
                    {post.photos.length &&
                        <>
                            <div style={{ backgroundImage: `url(${currentPhoto})` }}
                                className="editing-modal__current-photo"
                            />
                            <div className="editing-modal__mini-photos">
                                {post.photos.map((photo: string, index: number) =>
                                    <div
                                        key={`${photo}-${index}`}
                                        style={{ backgroundImage: `url(${photo})` }}
                                        onClick={() => setCurrentPhoto(photo)}
                                        className="editing-modal__mini-photos__item"
                                    >
                                        <img src={closeIcon} alt="close icon" className="editing-modal__mini-photos__item__close" onClick={(e) => removePhoto(e, index)} />
                                    </div>

                                )}
                            </div>
                        </>
                    }
                    <div className="editing-modal__add-photo">
                        <label
                            className="editing-modal__add-photo__label"
                            htmlFor="editing-modal-add-photo">
                            <span className="editing-modal__add-photo__label__icon">
                                <img
                                    className="editing-modal__add-photo__label__icon__photo"
                                    src={addPhotoIcon}
                                    alt="add photo" />
                            </span>
                            <span className="editing-modal__add-photo__label__button">
                                Оберіть фотографію
                            </span>
                        </label>
                        <input
                            className="editing-modal__add-photo__input"
                            id="editing-modal-add-photo"
                            type="file"
                            accept="image/png, image/jpeg"
                            onChange={handleFileChange} />
                    </div>
                </div>
                <div className="editing-modal__info">
                    <Link className="editing-modal__user-info" to={`user/${post.creator.id}`}>
                        <Avatar size={40} photo={post.creator.creatorAvatar} />
                        <p>{post.creator.nickname}</p>
                    </Link>
                    <textarea value={description} onChange={onChangeDescription} />
                </div>
            </div>
        </Modal>
    );
};
