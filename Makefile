user_client:
	kitex -type protobuf -I ./idl/ ./idl/user.proto
server:
	kitex -type protobuf -service comment -I ./idl/ ./idl/comment.proto
	kitex -type protobuf -service favorite -I ./idl/ ./idl/favorite.proto
	kitex -type protobuf -service feed -I ./idl/ ./idl/feed.proto
	kitex -type protobuf -service publish -I ./idl/ ./idl/publish.proto
	kitex -type protobuf -service relation -I ./idl/ ./idl/relation.proto
	kitex -type protobuf -service user -I ./idl/ ./idl/user.proto