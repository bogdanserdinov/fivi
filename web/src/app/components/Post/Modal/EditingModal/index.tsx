import { Avatar } from '@components/common/Avatar';
import { Modal } from '@components/common/Modal';
import { useEffect, useState } from 'react';

import addPhotoIcon from '@img/post/addPhotoIcon.png';
import closeIcon from '@img/User/Post/closeIcon.png';
import { Post, PostUpdateData } from '@/post';
import { addPostPhotos, deletePostPhoto, deletePostPhotos, setPostPhotos } from '@/app/store/reducers/posts';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';
import { RootState } from '@/app/store';
import { convertToBase64 } from '@/app/internal/convertImage';
import { updatePost } from '@/app/store/actions/posts';

import './index.scss';

const SECOND_INDEX = 1;

export const EditingModal: React.FC<{
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
    post: Post;
}> = ({ setIsOpenModal, post }) => {
    const [files, setFiles] = useState<string[]>([]);
    const [description, setDescription] = useState(post.description);
    const [currentPhoto, setCurrentPhoto] = useState('');
    const dispatch = useAppDispatch();

    const postPhotos: string[] | null = useAppSelector((state: RootState) => state.postsReducer.postPhotos);

    const onChangeDescription = (e: any) => {
        setDescription(e.target.value);
    };

    const handleFileChange = async(e: React.ChangeEvent<HTMLInputElement>) => {
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

    const deletePhoto = (index: number) => {
        if (postPhotos[index] === currentPhoto && postPhotos.length > 0) {
            setCurrentPhoto(postPhotos[0]);
        }
        dispatch(deletePostPhoto(postPhotos[index]));
    };

    const sendChanges = async() => {
        await dispatch(updatePost(new PostUpdateData(
            post.postId,
            description,
            files
        )));


        setIsOpenModal(false);
    };

    const getPhotosArray = () => {
        const sliderPhotos: string[] = [];

        for (let index = 0; index < post.numOfImages; index++) {
            sliderPhotos.push(`${window.location.origin}/images/posts/${post.postId}/${index}.png`);
        }

        return sliderPhotos;
    };

    useEffect(() => {
        const images = getPhotosArray();
        dispatch(deletePostPhotos());
        dispatch(setPostPhotos(images));
        setCurrentPhoto(`${window.location.origin}/images/posts/${post.postId}/0.png`);
        setDescription(post.description);
    }, [post]);

    return (
        <Modal classname="editing-modal" setIsOpenModal={setIsOpenModal}>
            <div className="editing-modal__content">
                <div
                    className="editing-modal__photos"
                >
                    {postPhotos.length ?
                        <>
                            <div style={{ backgroundImage: `url(${currentPhoto})` }}
                                className="editing-modal__current-photo"
                            />
                            <div className="editing-modal__mini-photos">
                                {postPhotos.map((photo: string, index: number) =>
                                    <div
                                        key={`${photo}-${index}`}
                                        style={{ backgroundImage: `url(${photo})` }}
                                        onClick={() => setCurrentPhoto(photo)}
                                        className="editing-modal__mini-photos__item"
                                    >
                                        <img src={closeIcon} alt="close icon" className="editing-modal__mini-photos__item__close" onClick={(e) => deletePhoto(index)} />
                                    </div>

                                )}
                            </div>
                        </>
                        :
                        <div className="editing-modal__no-photos">
                            No photos
                        </div>
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
                            onChange={handleFileChange}
                            multiple
                            hidden />

                    </div>
                </div>
                <div className="editing-modal__info">
                    <div>
                        <div className="editing-modal__user-info" >
                            <Avatar size={40} userId={post.creatorProfile.creatorId} isAvatarExists={post.creatorProfile.isAvatarExists} />
                            <p className="editing-modal__user-info__username">{post.creatorProfile.username}</p>
                        </div>
                        <textarea
                            className="editing-modal__comment"
                            value={description}
                            onChange={onChangeDescription} />

                    </div>

                    <button
                        onClick={() => sendChanges()}
                        className="editing-modal__send-changes"
                    >
                        Застосувати зміни
                    </button>
                </div>
            </div>
        </Modal >
    );
};
