import { APIClient } from '.';

import { Comment, CommentCreate, Creator, Post, PostAddData, PostUpdateData } from '@/post';

const NO_ITEMS_ARRAY = 0;
/**
 * PostsClient is a http implementation of posts API.
 * Exposes all posts-related functionality.
 */
export class PostsClient extends APIClient {
    /** Creates post */
    public async createPost(post: PostAddData): Promise<Post> {
        const path = `${this.ROOT_PATH}/posts/v1/posts`;
        const response = await this.http.post(path, JSON.stringify(post));

        if (!response.ok) {
            await this.handleError(response);
        }

        const postData = await response.json();

        const comments = postData.post.comments.length > NO_ITEMS_ARRAY ?
            postData.map((comment: any) =>
                new Comment(
                    comment.identifier,
                    comment.text,
                    comment.post_id,
                    comment.username,
                    comment.user_id,
                    comment.user_image
                )
            ) : [];

        const creator = new Creator(
            postData.post.creator_profile.id,
            postData.post.creator_profile.email,
            postData.post.creator_profile.username,
            postData.post.creator_profile.numbOfPosts,
            postData.post.creator_profile.subscribers,
            postData.post.creator_profile.subscribtions,
            postData.post.creator_profile.isAvatarExists,
        );

        return new Post(
            postData.post.identifier,
            postData.post.text,
            postData.post.creator_id,
            postData.post.creator_username,
            postData.post.images,
            postData.post.num_of_images,
            postData.post.num_of_likes,
            postData.post.num_of_comments,
            comments,
            postData.post.is_liked,
            creator
        );
    }

    /** Gets posts for user profile */
    public async getPostsProfile(userId: string): Promise<Post[]> {
        const path = `${this.ROOT_PATH}/posts/v1/creator/${userId}`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const posts = await response.json();

        return posts.posts.map((post: any) => {
            const creator = new Creator(
                post.creator_profile.id,
                post.creator_profile.email,
                post.creator_profile.username,
                post.creator_profile.numbOfPosts,
                post.creator_profile.subscribers,
                post.creator_profile.subscribtions,
                post.creator_profile.isAvatarExists,
            );
            const comments = post.comments.length > NO_ITEMS_ARRAY ?
                post.comments.map((comment: any) =>
                    new Comment(
                        comment.identifier,
                        comment.text,
                        comment.post_id,
                        comment.username,
                        comment.user_id,
                        comment.user_image,
                        comment.is_avatar_exists
                    )
                ) : [];

            return (
                new Post(
                    post.identifier,
                    post.text,
                    post.creator_id,
                    post.creator_username,
                    post.images,
                    post.num_of_images,
                    post.num_of_likes,
                    post.num_of_comments,
                    comments,
                    post.is_liked,
                    creator
                ));
        }
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
    public async getPostsHomePage(): Promise<any> {
        const path = `${this.ROOT_PATH}/posts/v1/posts`;
        const response = await this.http.get(path);

        if (!response.ok) {
            await this.handleError(response);
        }
        const posts = await response.json();

        return posts.post.map((post: any) => {
            const creator = new Creator(
                post.creator_profile.id,
                post.creator_profile.email,
                post.creator_profile.username,
                post.creator_profile.numbOfPosts,
                post.creator_profile.subscribers,
                post.creator_profile.subscribtions,
                post.creator_profile.isAvatarExists,
            );
            const comments = post.comments.length > NO_ITEMS_ARRAY ?
                post.comments.map((comment: any) =>
                    new Comment(
                        comment.identifier,
                        comment.text,
                        comment.post_id,
                        comment.username,
                        comment.user_id,
                        comment.user_image
                    )
                ) : [];

            return (
                new Post(
                    post.identifier,
                    post.text,
                    post.creator_id,
                    post.creator_username,
                    post.images,
                    post.num_of_images,
                    post.num_of_likes,
                    post.num_of_comments,
                    comments,
                    post.is_liked,
                    creator
                ));
        }
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
    public async update(post: PostUpdateData): Promise<Post> {
        const path = `${this.ROOT_PATH}/v1/posts/post/${post.postId}`;
        const response = await this.http.put(path, JSON.stringify({ text: post.text, images: post.images }));


        if (!response.ok) {
            await this.handleError(response);
        }

        const postData = await response.json();

        const comments: Comment[] = [];

        postData.post.comments.length > NO_ITEMS_ARRAY &&
            postData.post.comments.map((comment: any) => {
                comments.push(new Comment(
                    comment.identifier,
                    comment.text,
                    comment.post_id,
                    comment.username,
                    comment.user_id,
                    comment.user_image
                ));
            });

        const creator = new Creator(
            postData.post.creator_profile.id,
            postData.post.creator_profile.email,
            postData.post.creator_profile.username,
            postData.post.creator_profile.numbOfPosts,
            postData.post.creator_profile.subscribers,
            postData.post.creator_profile.subscribtions,
            postData.post.creator_profile.isAvatarExists,
        );

        return new Post(
            postData.post.identifier,
            postData.post.text,
            postData.post.creator_id,
            postData.post.creator_username,
            postData.post.images,
            postData.post.num_of_images,
            postData.post.num_of_likes,
            postData.post.num_of_comments,
            comments,
            postData.post.is_liked,
            creator
        );
    }

    /** Likes/dislikes post */
    public async likeAndDislike(postId: string): Promise<void> {
        const path = `${this.ROOT_PATH}/likes/v1`;
        const response = await this.http.post(path, JSON.stringify({ 'post_id': postId }));

        if (!response.ok) {
            await this.handleError(response);
        }
    }

    /** Creates comment */
    public async createComments(comment: CommentCreate): Promise<Comment> {
        const path = `${this.ROOT_PATH}/comments/v1`;
        const response = await this.http.post(path, JSON.stringify(comment));

        if (!response.ok) {
            await this.handleError(response);
        }

        const commentData = await response.json();

        return new Comment(
            commentData.comment.identifier,
            commentData.comment.text,
            commentData.comment.post_id,
            commentData.comment.username,
            commentData.comment.user_id,
            commentData.comment.user_image,
            commentData.comment.is_avatar_exists
        );
    }
}
