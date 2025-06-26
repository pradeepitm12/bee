# bee

Bee is a go grpc based post management service.

Current available services.

- CreatePost
  - takes given input of title, content, author, tags and create an entry of post associated with id which is an uui response with the created post.
- ReadPost
  - takes  id and response associated post
- UpdatePost
  - takes id, update info title, content, author, tags and updates the post response with the updated post. 
- DeletePost
  - takes id as input and deletes the entry of the id and associated post response success message


