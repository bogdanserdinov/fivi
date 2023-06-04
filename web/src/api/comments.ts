import { APIClient } from '.';

import { CommentCreate } from '@/post';

/**
 * CommentsClient is a http implementation of comments API.
 * Exposes all comments-related functionality.
 */
export class CommentsClient extends APIClient {
    /** Creates comment */
    public async createComments(comment: CommentCreate): Promise<void> {
        const path = `${this.ROOT_PATH}/comments/v1`;
        const response = await this.http.post(path, JSON.stringify(comment));

        if (!response.ok) {
            await this.handleError(response);
        }
    }
}
