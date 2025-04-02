# Blog Api 


# Database 
- __Permissions__ – User permissions. 
- __Users__ – Authors who write posts.
- __Posts__ – Blog posts.
- __Categories__ – Categorization for posts.
- __Tags__ – Optional tagging system.
- __Comments__ – User comments on posts.



# To Do 
### Users Endpoint
- [X] `GET /users` – List all users x.
- [X] `GET /users/:id` – Get user by ID.
- [X] `POST /users` – Register new user .
- [ ] `PATCH|PUT /users/:id` – Update user .
- [ ] `DELETE /users/:id` – Delete user .
- [ ] implement test cases 

### Error Handling
- [X] Standardized JSON errors 
    ```json 
    { 
        "message": "error message", 
        "status":"error", 
        "errors": { // [optional] with validation errors
            "key": ["error message"]
        } 
    }
    ```

### Pagination
- [ ] GET /posts?page=1&limit=10 – Paginated responses.
- [ ] Metadata: totalPages, currentPage, nextPage.

### Authentication
- [ ] implement jwt token generation
- [ ] implement middleware for authentication
- [ ] `POST /auth/login` – JWT token generation.
- [ ] `POST /auth/refresh` – Refresh expired tokens.
- [ ] `POST /auth/logout` – Revoke current token.
- [ ] `POST /auth/register` – User registration.
- [ ] Filter by `tags`, `categories`.

### Profile 
- [ ] `GET /profile` – List all users x.
- [ ] `PUT /profile` – update user profile.
- [ ] `PUT /profile/password` – Update user password .
- [ ] `DELETE /profile/` – Delete user profile .
- [ ] implement test cases 


### Authorization
- [ ] implement permissions
- [ ] implement helper functions for authorization
- [ ] implement test cases

### Posts Endpoint
- [ ] `GET /posts` – List posts .
- [ ] `POST /posts` – Create post (auth required).
- [ ] `GET /posts/:id` – Get post .
- [ ] `PATCH /posts/:id` – Update post (auth required).
- [ ] `DELETE /posts/:id` – Delete post (auth required).
- [ ] implement test cases



### Categories Endpoint
- [ ] `GET /categories` – List Categories .
- [ ] `POST /categories` – Create categories (auth required).
- [ ] `GET /categories/:id` – Get categories .
- [ ] `PATCH /categories/:id` – Update categories (auth required).
- [ ] `DELETE /categories/:id` – Delete categories (auth required).
- [ ] implement test cases

### Tags Endpoint
- [ ] `GET /tags` – List tags .
- [ ] `POST /tags` – Create tags (auth required).
- [ ] `GET /tags/:id` – Get tags .
- [ ] `PATCH /tags/:id` – Update tags (auth required).
- [ ] `DELETE /tags/:id` – Delete tags (auth required).
- [ ] implement test cases

### Comments Endpoint
- [ ] `GET /comments` – List comments .
- [ ] `POST /comments` – Create comments (auth required).
- [ ] `GET /comments/:id` – Get comments .
- [ ] `PATCH /comments/:id` – Update comments (auth required).
- [ ] `DELETE /comments/:id` – Delete comments (auth required).
- [ ] implement test cases



# 📌 Notes
- Prioritize Core Features first.
- Use [X] for done, [ ] for pending, [/] for partial.