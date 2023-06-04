import { APIClient } from '.';

import { Post, PostAddData, PostUpdateData } from '@/post';

/**
 * PostsClient is a http implementation of posts API.
 * Exposes all posts-related functionality.
 */
export class PostsClient extends APIClient {
    /** Creates post */
    public async createPost(post: PostAddData): Promise<void> {
        const path = `${this.ROOT_PATH}/posts/v1/posts`;
        const response = await this.http.post(path, JSON.stringify(post));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** Gets posts for user profile */
    public async getPostsProfile(userId: string): Promise<Post[]> {
        const path = `${this.ROOT_PATH}/posts/v1/creator/${userId}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const posts = await response.json();

        return posts.map((post: any) =>
            new Post(
                post.identifier,
                post.num_of_likes,
                post.isFavorite,
                post.text,
                post.comments,
            )
        );
    }

    /** Gets post */
    public async getPost(postId: string): Promise<Post> {
        const path = `${this.ROOT_PATH}/posts/v1/post/${postId}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const post = await response.json();

        return new Post(
            post.identifier,
            post.num_of_likes,
            post.isFavorite,
            post.text,
            post.comments,
        );
    }

    /** Gets posts for home page */
    public async getPostsHomePage(): Promise<Post[]> {
        const path = `${this.ROOT_PATH}/posts/v1/posts`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const posts = await response.json();

        return posts.map((post: any) =>
            new Post(
                post.identifier,
                post.num_of_likes,
                post.isFavorite,
                post.text,
                post.comments,
            )
        );
    }

    /** Deletes post */
    public async delete(postId: string): Promise<void> {
        const path = `${this.ROOT_PATH}/v1/posts/post/${postId}`;
        const response = await this.http.delete(path);

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** Updates post */
    public async update(post: PostUpdateData): Promise<void> {
        const path = `${this.ROOT_PATH}/v1/posts/post/${post.postId}`;
        const response = await this.http.put(path, JSON.stringify(post.text, post.images));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** Likes/dislikes post */
    public async likeAndDislike(postId: string): Promise<void> {
        const path = `${this.ROOT_PATH}/likes/v1`;
        const response = await this.http.post(path, JSON.stringify(postId));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
}
