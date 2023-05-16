import { Post } from '@components/Post';
import { posts } from '@/mocked/posts';

import './index.scss';

const Home = () =>
    <div className="home">
        <div className="home__posts">
            {posts.map((post) =>
                <Post post={post} />)
            }
        </div>
    </div>;

export default Home;
