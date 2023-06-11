import { ChangeEvent, useEffect, useState } from 'react';

import { Modal } from '@components/common/Modal';
import addPhotoIcon from '@img/post/addPhotoIcon.png';
import closeIcon from '@img/User/Post/closeIcon.png';
import { convertToBase64 } from '@/app/internal/convertImage';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { PostAddData } from '@/post';
import { addPostPhotos, deletePostPhoto, deletePostPhotos, setCurrentId } from '@/app/store/reducers/posts';
import { RootState } from '@/app/store';
import { createPost } from '@/app/store/actions/posts';


import './index.scss';
import { User } from '@/users';

const SECOND_INDEX = 1;

export const PostCreateModal: React.FC<{
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
}>
    = ({ setIsOpenModal }) => {
        const [files, setFiles] = useState<string[]>();
        const [description, setDescription] = useState<string>('');
        const dispatch = useAppDispatch();

        const user: User = useAppSelector((state: RootState) => state.usersReducer.user);
        const postPhotos: string[] | [] = useAppSelector((state: RootState) => state.postsReducer.postPhotos);

        const handleFileChange = async(e: ChangeEvent<HTMLInputElement>) => {
            if (e.target.files?.length) {
                const photosData = [];
                const filesData = [];

                const uploadedFiles = Array.from(e.target.files);

                for await (const uploadedFile of uploadedFiles) {
                    photosData.push(URL.createObjectURL(uploadedFile));

                    const convertedFile: string = await convertToBase64(uploadedFile);
                    filesData.push(convertedFile.split(',')[SECOND_INDEX]);
                }

                dispatch(addPostPhotos(photosData));
                setFiles(filesData);
            }
        };

        const create = async() => {
            try {
                dispatch(setCurrentId(user.userId));
                await dispatch(createPost(new PostAddData(
                    description,
                    files
                )));

                setIsOpenModal(false);
            }
            catch (e) {
                // error
            }
        };

        const deletePhoto = (index: number) => {
            dispatch(deletePostPhoto(postPhotos[index]));
        };

        useEffect(() => {
            dispatch(deletePostPhotos());
        }, []);


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
                            id="create-post-add-photo"
                            type="file"
                            accept="image/png, image/jpeg"
                            onChange={handleFileChange}
                            multiple
                            hidden
                        />
                    </div>
                    <div className="create-post-modal__photos">{
                        postPhotos && postPhotos.map((photo, index) =>
                            <div style={{ backgroundImage: `url(${photo}` }}
                                className="create-post-modal__photos__item"
                                key={photo}>
                                <button
                                    type="button"
                                    className="create-post-modal__photos__item__close"
                                    onClick={() => deletePhoto(index)}
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
                    <textarea
                        className="create-post-modal__description"
                        onChange={e => setDescription(e.target.value)}
                    />
                    <button
                        className="create-post-modal__button"
                        type="button"
                        onClick={() => create()}>
                        Create
                    </button>
                </form>
            </Modal>
        );
    };
