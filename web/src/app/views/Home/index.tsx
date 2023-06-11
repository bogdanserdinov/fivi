import { useEffect } from 'react';
import { PostPage } from '@components/Post';
import { useAppDispatch, useAppSelector } from '@/app/hooks/useReduxToolkit';

import { getPostsHomePage } from '@/app/store/actions/posts';
import { RootState } from '@/app/store';
import { Post } from '@/post';

import './index.scss';

const Home = () => {
    const dispatch = useAppDispatch();

    const posts: Post[] | [] = useAppSelector((state: RootState) => state.postsReducer.homePosts);

    const getPosts = async() => {
        await dispatch(getPostsHomePage());
    };

    useEffect(() => {
        getPosts();
    }, []);

    return (
        <div className="home">
            {posts.length ?
                <div className="home__posts">
                    {posts.map((post) =>
                        <PostPage post={post} key={post.postId} />)
                    }
                </div> :
                <div className="home__no-posts">
                    There is no posts yet
                </div>
            }
        </div>);
};

export default Home;
