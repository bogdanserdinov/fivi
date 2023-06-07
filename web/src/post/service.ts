import { Comment, CommentCreate, Post, PostAddData, PostUpdateData } from '.';
import { PostsClient } from '@/api/posts';

/**
 * Exposes all posts related logic.
 */
export class PostsService {
    private readonly posts: PostsClient;
    /** PostsService contains http implementation of posts API  */
    public constructor(posts: PostsClient) {
        this.posts = posts;
    }
    /** handles creating post */
    public async createPost(post: PostAddData): Promise<Post> {
        return await this.posts.createPost(post);
    }

    /** handles deleting post */
    public async delete(postId: string): Promise<void> {
        await this.posts.delete(postId);
    }
    /** gets post info */
    public async getPost(postId: string): Promise<Post> {
        return await this.posts.getPost(postId);
    }

    /** updates post info */
    public async update(post: PostUpdateData): Promise<void> {
        await this.posts.update(post);
    }

    /** Gets post in home page */
    public async getPostsHomePage(): Promise<Post[]> {
        return await this.posts.getPostsHomePage();
    }

    /** Gets post in user profile */
    public async getPostsProfile(userId: string): Promise<Post[]> {
        return await this.posts.getPostsProfile(userId);
    }

    /** Likes and dislikes post */
    public async likeAndDislike(postId: string): Promise<void> {
        await this.posts.likeAndDislike(postId);
    }

    /** Likes and dislikes post */
    public async sendComment(comment: CommentCreate): Promise<Comment> {
        return await this.posts.createComments(comment);
    }
}
