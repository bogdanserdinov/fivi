import { useState } from 'react';

import arrowIcon from '@static/img/post/slider/arrow.png';

import './index.scss';

const SLIDER_STEP = 1;
const FIRST_SLIDE = 0;
const CHECK_SLIDER_PHOTO_INDEX = 1;
const ONE_PHOTO = 1;

export const PostSlider: React.FC<{ sliderImages: string[];classname?:string }> = ({ sliderImages,classname }) => {
    const [current, setCurrent] = useState<number>(FIRST_SLIDE);
    const sliderImagesLength = sliderImages.length;

    const nextSlide = () => {
        setCurrent(current === sliderImagesLength - CHECK_SLIDER_PHOTO_INDEX ? FIRST_SLIDE : current + SLIDER_STEP);
    };

    const prevSlide = () => {
        setCurrent(current === FIRST_SLIDE ? sliderImagesLength - CHECK_SLIDER_PHOTO_INDEX : current - SLIDER_STEP);
    };

    return (
        <div className={`post-slider ${classname?classname:''}` }>
            {sliderImagesLength !== ONE_PHOTO ?
                <>
                    {current !== FIRST_SLIDE &&
                        <div className="post-slider__arrow post-slider__arrow__prev" onClick={() => prevSlide()}>
                            <img src={arrowIcon} alt="arrow-left" className="post-slider__arrow__prev__image" />
                        </div>
                    }

                    <div className="post-slider__container">
                        {sliderImages.map((sliderImage, index) =>
                            <div key={`${sliderImage}`}
                                className={` post-slider__item ${index === current ? 'active' : ''}`}
                            >
                                {index === current &&
                                   <div style={{ backgroundImage: `url(${ sliderImages[FIRST_SLIDE]} )`}} className="post-slider__item__image" />
                                }
                            </div>
                        )}
                    </div>
                    {sliderImagesLength - CHECK_SLIDER_PHOTO_INDEX !== current &&
                        <div className="post-slider__arrow post-slider__arrow__next" onClick={() => nextSlide()}>
                            <img src={arrowIcon} alt="arrow-left" className="post-slider__arrow__next__image" />
                        </div>
                    }
                </>
                :
                <div style={{ backgroundImage: `url(${ sliderImages[FIRST_SLIDE]} )`}} className="post-slider__item__image" />}
        </div>
    );
};
