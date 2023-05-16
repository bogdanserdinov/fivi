import mockedAvatar from '@static/img/post/mockedAvatar.png';
import mockedSlider from '@static/img/post/slider/mockedSlider.jpeg';
import mockedSlider1 from '@static/img/post/slider/mockedSlider1.jpg';

const MOCKED_PHOTOS = [mockedSlider1, mockedSlider, mockedSlider1, mockedSlider, mockedSlider1, mockedSlider, mockedSlider1, mockedSlider, mockedSlider1];

const comment = {
    commentId: 'csafasf',
    commentorId: 'cacscadqwdqwd',
    commentorName: 'acjnscnjnja',
    commentText: 'ascnjasnjcbajbwjcqbw',
    commentorAvatar: mockedAvatar,
};

export const post = {
    id: 'saadsdsa',
    favorites: 228,
    isFavorite: true,
    description: 'asfasf',
    photos: MOCKED_PHOTOS,
    creator: {
        id: 'csafasf',
        creatorAvatar: mockedAvatar,
        name: 'afwsfwf',
        nickname: 'scaskcinas'
    },
    comments: [comment, comment, comment, comment, comment, comment, comment, comment, comment, comment, comment],
}

export const posts = [
    post, post, post, post, post, post,
];
