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




Some runs  
go run cmd/server/main.go
2025/06/26 22:11:46 gRPC server is running on port :50051
cd2025/06/26 22:12:13 Received CreatePost request - Title: My First Blog, Author: ABC
2025/06/26 22:12:13 Created new post with ID: 7d6614c2-eb84-4253-8f24-f4ce17c9aaa3
2025/06/26 22:12:13 Received ReadPost request for ID: 7d6614c2-eb84-4253-8f24-f4ce17c9aaa3
2025/06/26 22:12:13 The post: &{ID:7d6614c2-eb84-4253-8f24-f4ce17c9aaa3 Title:My First Blog Content:Some content Author:ABC PublicationDate:2025-06-26 22:12:13.489836 +0530 IST m=+26.840109334 LastModified:0001-01-01 00:00:00 +0000 UTC Tags:[golang grpc]}
2025/06/26 22:12:13 Received UpdatePost request for ID: 7d6614c2-eb84-4253-8f24-f4ce17c9aaa3
2025/06/26 22:12:13 Updated post with ID: 7d6614c2-eb84-4253-8f24-f4ce17c9aaa3
2025/06/26 22:12:13 Received DeletePost request for ID: 7d6614c2-eb84-4253-8f24-f4ce17c9aaa3
2025/06/26 22:12:13 Deleted post with ID: 7d6614c2-eb84-4253-8f24-f4ce17c9aaa3
2025/06/26 22:12:13 Received ReadPost request for ID: 7d6614c2-eb84-4253-8f24-f4ce17c9aaa3
2025/06/26 22:12:13 Error: Post not found: 7d6614c2-eb84-4253-8f24-f4ce17c9aaa3


go run main.go
Created Post: postID:"7d6614c2-eb84-4253-8f24-f4ce17c9aaa3"  title:"My First Blog"  content:"Some content"  author:"ABC"  publicationDate:"2025-06-26T22:12:13+05:30"
Read Post: postID:"7d6614c2-eb84-4253-8f24-f4ce17c9aaa3"  title:"My First Blog"  content:"Some content"  author:"ABC"  publicationDate:"2025-06-26T22:12:13+05:30"
Updated Post: postID:"7d6614c2-eb84-4253-8f24-f4ce17c9aaa3"  title:"My First Blog (Updatedd)"  content:"Updated content"  author:"ABC"  publicationDate:"2025-06-26T22:12:13+05:30"
Post deleted successfully
As expected, post not found after delete: rpc error: code = NotFound desc = post not found



post1: dff9fd08-78b0-4a64-a547-3c4259de509b
post2: 6b35cb75-9b7c-4943-a0a9-84be4c4d6611

