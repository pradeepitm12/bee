#1 Create
grpcurl -plaintext -d '{ 
"title": "My First Blog",
"content": "This is a gRPC-powered blog post!",
"author": "Alice",
"tags": ["golang", "grpc"]
}' localhost:50051 blog.BlogPostService/CreatePost

{
"post": {
"postID": "c2b11eac-9c0d-48ae-b823-f54e5586e8e0",
"title": "My First Blog",
"content": "This is a gRPC-powered blog post!",
"author": "Alice",
"publicationDate": "2025-06-26T18:25:53+05:30"
}
}
#2 Create
grpcurl -plaintext -d '{
"title": "My Second Blog",
"content": "This is a gRPC-powered blog post!",
"author": "Alice",
"tags": ["golang", "grpc"]
}' localhost:50051 blog.BlogPostService/CreatePost

{
"post": {
"postID": "e4bf68c9-b3cc-49b2-b1e9-62e3d470ca72",
"title": "My Second Blog",
"content": "This is a gRPC-powered blog post!",
"author": "Alice",
"publicationDate": "2025-06-26T18:28:32+05:30"
}
}

#3 Create
grpcurl -plaintext -d '{
"title": "My Second Blog",
"content": "This is a gRPC-powered blog post!",
"author": "Alice",
"tags": ["golang", "grpc"]
}' localhost:50051 blog.BlogPostService/CreatePost

#4 Get
grpcurl -plaintext -d '{
"postID": "e4bf68c9-b3cc-49b2-b1e9-62e3d470ca72"
}' localhost:50051 blog.BlogPostService/ReadPost

{
"post": {
"postID": "e4bf68c9-b3cc-49b2-b1e9-62e3d470ca72",
"title": "My Second Blog",
"content": "This is a gRPC-powered blog post!",
"author": "Alice",
"publicationDate": "2025-06-26T18:28:32+05:30"
}
}

#5 Delete

grpcurl -plaintext -d '{
"postID": "e4bf68c9-b3cc-49b2-b1e9-62e3d470ca72"
}' localhost:50051 blog.BlogPostService/DeletePost

{
"message": "Post e4bf68c9-b3cc-49b2-b1e9-62e3d470ca72 Deleted"
}

#6 Read Deleted Post

grpcurl -plaintext -d '{
"postID": "e4bf68c9-b3cc-49b2-b1e9-62e3d470ca72"
}' localhost:50051 blog.BlogPostService/ReadPost
ERROR:
Code: NotFound
Message: post not found

#7 Update Post

grpcurl -plaintext -d '{
"id": "c2b11eac-9c0d-48ae-b823-f54e5586e8e0",
"title": "My Second Blog",
"content": "This is a gRPC-powered blog post!",
"author": "Pradeep",
"tags": ["golang", "grpc"]
}' localhost:50051 blog.BlogPostService/UpdatePost

{
"post": {
"postID": "c2b11eac-9c0d-48ae-b823-f54e5586e8e0",
"title": "My Second Blog",
"content": "This is a gRPC-powered blog post!",
"author": "Pradeep",
"publicationDate": "2025-06-26T18:25:53+05:30"
}
}

# Get Updated POst
grpcurl -plaintext -d '{
"postID": "c2b11eac-9c0d-48ae-b823-f54e5586e8e0"
}' localhost:50051 blog.BlogPostService/ReadPost







