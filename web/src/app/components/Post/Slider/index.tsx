import { useEffect, useState } from 'react';

import arrowIcon from '@static/img/post/slider/arrow.png';

import './index.scss';
import { PostPhoto } from '@components/common/PostPhoto';

const SLIDER_STEP = 1;
const FIRST_SLIDE = 0;
const CHECK_SLIDER_PHOTO_INDEX = 1;
const ONE_PHOTO = 1;

export const PostSlider: React.FC<{ postId: string; numOfImages: number; classname?: string }> = ({ postId, numOfImages, classname }) => {
    const [postPhotos, setPostPhotos] = useState<string[]>([]);

    const [current, setCurrent] = useState<number>(FIRST_SLIDE);


    const nextSlide = () => {
        setCurrent(current === numOfImages - CHECK_SLIDER_PHOTO_INDEX ? FIRST_SLIDE : current + SLIDER_STEP);
    };

    const prevSlide = () => {
        setCurrent(current === FIRST_SLIDE ? numOfImages - CHECK_SLIDER_PHOTO_INDEX : current - SLIDER_STEP);
    };

    const getPhotosArray = () => {
        const sliderPhotos: string[] = [];

        for (let index = 0; index < numOfImages; index++) {
            sliderPhotos.push(`${window.location.origin}/images/posts/${postId}/${index}.png`);
        }

        return sliderPhotos;
    };

    useEffect(() => {
        const slides = getPhotosArray();
        setPostPhotos(slides);
    }, []);

    return (
        <div className={`post-slider ${classname ? classname : ''}`}>
            {numOfImages !== ONE_PHOTO ?
                <>
                    {current !== FIRST_SLIDE &&
                        <div className="post-slider__arrow post-slider__arrow__prev" onClick={() => prevSlide()}>
                            <img src={arrowIcon} alt="arrow-left" className="post-slider__arrow__prev__image" />
                        </div>
                    }

                    <div className="post-slider__container">
                        {postPhotos.map((sliderImage, index) =>
                            <div key={`${sliderImage}`}
                                className={` post-slider__item ${index === current ? 'active' : ''}`}
                            >
                                {index === current &&
                                    <PostPhoto urlPhoto={postPhotos[index]} width={400} height={500} isPostPhotoExist={true} />
                                }
                            </div>
                        )}
                    </div>
                    {numOfImages - CHECK_SLIDER_PHOTO_INDEX !== current &&
                        <div className="post-slider__arrow post-slider__arrow__next" onClick={() => nextSlide()}>
                            <img src={arrowIcon} alt="arrow-left" className="post-slider__arrow__next__image" />
                        </div>
                    }
                </>
                :
                <PostPhoto urlPhoto={postPhotos[FIRST_SLIDE]} width={400} height={500} isPostPhotoExist={true} />

            }
        </div>
    );
};
