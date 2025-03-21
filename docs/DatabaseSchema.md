## 📌 Database Schema

### 🧑 Users
- `id` → Primary Key (uint)
- `name` → User's full name (string)
- `email` → Unique user email (string)
- `password` → Hashed password (string)
- `created_at` → Timestamp
- `updated_at` → Timestamp

---

### 📝 Posts
- `id` → Primary Key (uint)
- `title` → Post title (string)
- `content` → Post content (text)
- `user_id` → Foreign Key → Users(id) (uint)
- `category_id` → Foreign Key → Categories(id) (uint)
- `created_at` → Timestamp
- `updated_at` → Timestamp

---

### 📂 Categories
- `id` → Primary Key (uint)
- `name` → Unique category name (string)

---

### 🏷️ Tags
- `id` → Primary Key (uint)
- `name` → Tag name (string)

---

### 🔗 Post_Tags (Many-to-Many)
- `post_id` → Foreign Key → Posts(id) (uint)
- `tag_id` → Foreign Key → Tags(id) (uint)

---

### 💬 Comments
- `id` → Primary Key (uint)
- `post_id` → Foreign Key → Posts(id) (uint)
- `user_id` → Foreign Key → Users(id) (uint)
- `content` → Comment content (text)
- `created_at` → Timestamp
- `updated_at` → Timestamp

---

### 🔗 Relationships
- **Users → Posts** → One user can create multiple posts.  
- **Categories → Posts** → One category can contain multiple posts.  
- **Posts ↔ Tags** → Many-to-Many relationship (via `Post_Tags`).  
- **Posts → Comments** → A post can have multiple comments.  
- **Users → Comments** → A user can comment on multiple posts.  

---
