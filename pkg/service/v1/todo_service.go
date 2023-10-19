package v1

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_bri/pkg/api/v1"
	"grpc_bri/pkg/repository/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type toDoServiceServer struct {
	todoRepository repository.ITodoRepository
}

// NewToDoServiceServer creates To Do service
func NewToDoServiceServer(todoRepository repository.ITodoRepository) v1.ToDoServiceServer {
	return &toDoServiceServer{todoRepository}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *toDoServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create new todo task
func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// insert To Do entity data
	todo := &repository.TodoModel{
		Title:       req.ToDo.Title,
		Description: req.ToDo.Description,
		Reminder:    req.ToDo.Reminder.AsTime(),
	}

	if err := s.todoRepository.CreateTodo(ctx, todo); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  int64(todo.ID),
	}, nil
}

// Read todo task
func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	todo, err := s.todoRepository.ReadTodo(ctx, fmt.Sprint())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.ReadResponse{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Id:          int64(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Reminder:    timestamppb.New(todo.Reminder),
		},
	}, nil
}

// Update todo task
func (s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	err := s.todoRepository.UpdateTodo(ctx, &repository.TodoModel{
		ID:          int32(req.ToDo.Id),
		Title:       req.ToDo.Title,
		Description: req.ToDo.Description,
		Reminder:    req.ToDo.Reminder.AsTime(),
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: 1,
	}, nil
}

// Delete todo task
func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	if err := s.todoRepository.Delete(ctx, fmt.Sprint(req.Id)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.DeleteResponse{
		Api:     apiVersion,
		Deleted: 1,
	}, nil
}

// ReadAll to do tasks
func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	todos, err := s.todoRepository.ReadAll(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	toDos := make([]*v1.ToDo, len(*todos))

	for i, model := range *todos {
		toDos[i] = &v1.ToDo{
			Id:          int64(model.ID),
			Title:       model.Title,
			Description: model.Description,
			Reminder:    timestamppb.New(model.Reminder),
		}
	}

	return &v1.ReadAllResponse{
		Api:   apiVersion,
		ToDos: toDos,
	}, nil
}
