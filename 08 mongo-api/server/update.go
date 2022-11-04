package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Manuel11713/mongo-api/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID: %v\n", err))
	}

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update")
	}

	if res.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Cannot find blog with Id")
	}

	return &emptypb.Empty{}, nil
}
