# Articles API reference


## Routes

<br/>

### Fetch all articles
#### (GET) `/articles`

<br/>
 
Parameters: -  
Returns: [[]Article](#Article) `(Object)`  
Authentication needed: `false`

<br/>

### Fetch an article by ID
#### (GET) `/articles/{id}`

<br/>
 
Parameters: Article ID `(Int)`  
Returns: [Article](#Article) `(Object)`  
Authentication needed: `false`

<br/>

### Create an article
#### (POST) `/articles`

<br/>

Parameters: [ArticleForm](#ArticleForm) `(Object)`  
Returns: [Article](#Article) `(Object)`  
Authentication needed: `true`

<br/>

### Edit an existing article by ID
#### (POST) `/articles/{id}`

<br/>

Parameters: [ArticleForm](#ArticleForm) `(Object)`  
Returns: [Article](#Article) `(Object)`  
Authentication needed: `true`

<br/>

### Create a comment on an article.  
#### (POST) `/articles/{id}/comment`

<br/>

Parameters: [CommentForm](#CommentForm) `(Object)`  
Returns: [Comment](#Comment) `(Object)`  
Authentication needed: `true`

<br/>

### Toggle like on a certain article.  
#### (POST) `/articles/{id}/like`

<br/>

Parameters: `true` or `false` `(Bool)`  
Authentication needed: `true`

<br/>

### Toggle dislike on a certain article.  
#### (POST) `/articles/{id}/dislike`

<br/>

Parameters: `true` or `false` `(Bool)`  
Authentication needed: `true`

<br/>

## Models

<br/>

### Article

Type: `Object`

Parameters:

• `ID` (Int) - Article ID  
• `username` (String) - Username of the article's author.  
• `title` (String) - Title of the article  
• `content` (String) - Content of the article  
• `create_date` (String) - Article creation date timestamp  
• `comments`  [([]Comment)](#Comment) - Array of comments linked to the article  
• `likes` (Int) - Number of likes on the Article  
• `dislike` (Int) - Number of dislikes on the Article  

<br/>

### ArticleForm

Type: `Object`

Parameters:

• `username` (String) - Username of the article's author.  
• `title` (String) - Title of the article  
• `content` (String) - Content of the article  

<br/>

### Comment

Type: `Object`  

Parameters:

• `ID` (Int) - Comment ID  
• `username` (String) - Username of the comment's author.  
• `title` (String) - Title of the comment  
• `content` (String) - Content of the comment  
• `create_date` (String) - Comment creation date timestamp  
• `article_id` (String) - ID of the corresponding [Article](#Article)  
• `likes` (Int) - Number of likes on the Comment  
• `dislike` (Int) - Number of dislikes on the Comment  

<br/>

### CommentForm

Type: `Object`

Parameters:

• `username` (String) - Username of the comment's author.  
• `title` (String) - Title of the comment  
• `content` (String) - Content of the comment  
• `article_id` (String) - ID of the corresponding [Article](#Article)  