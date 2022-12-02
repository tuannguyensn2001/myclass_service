package routes

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"picket/src/config"
	auth_transport "picket/src/features/auth/transport"
	auth_usecase "picket/src/features/auth/usecase"
	class_repository "picket/src/features/class/repository"
	class_transport "picket/src/features/class/transport"
	class_usecase "picket/src/features/class/usecase"
	media_transport "picket/src/features/media/transport"
	media_usecase "picket/src/features/media/usecase"
	test_repository "picket/src/features/test/repository"
	test_transport "picket/src/features/test/transport"
	test_usecase "picket/src/features/test/usecase"
	user_repository "picket/src/features/user/repository"
	user_transport "picket/src/features/user/transport"
	user_usecase "picket/src/features/user/usecase"
	authpb "picket/src/pb/auth"
	classpb "picket/src/pb/class"
	testpb "picket/src/pb/test"
	userpb "picket/src/pb/user"
	oauth_service "picket/src/services/oauth"
	storage_service "picket/src/services/storage"
)

type handler = func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

func RouteGrpc(ctx context.Context, s *grpc.Server, config config.IConfig) {
	oauthService := oauth_service.New(config)

	userRepository := user_repository.New(config.GetDB())
	userUsecase := user_usecase.New(userRepository)
	userTransport := user_transport.New(ctx, userUsecase)

	authUsecase := auth_usecase.New(nil, oauthService, userUsecase, config)
	authTransport := auth_transport.New(ctx, authUsecase)

	classRepository := class_repository.New(config.GetDB())
	classUsecase := class_usecase.New(classRepository)
	classTransport := class_transport.New(ctx, classUsecase)

	testRepository := test_repository.New(config.GetDB())
	testUsecase := test_usecase.New(testRepository)
	testTransport := test_transport.New(ctx, testUsecase)

	authpb.RegisterAuthServiceServer(s, authTransport)
	userpb.RegisterUserServiceServer(s, userTransport)
	classpb.RegisterClassServiceServer(s, classTransport)
	testpb.RegisterTestServiceServer(s, testTransport)
}

func RouteGw(ctx context.Context, gw *runtime.ServeMux, conn *grpc.ClientConn) {
	lists := []handler{authpb.RegisterAuthServiceHandler, userpb.RegisterUserServiceHandler, classpb.RegisterClassServiceHandler, testpb.RegisterTestServiceHandler}

	for _, item := range lists {
		err := item(ctx, gw, conn)
		if err != nil {
			zap.S().Fatalln(err)
		}
	}

	storageService := storage_service.New("fs", "src/storage")
	mediaUsecase := media_usecase.New(storageService)
	mediaTransport := media_transport.New(ctx, mediaUsecase)

	gw.HandlePath(http.MethodPost, "/api/v1/media/upload", mediaTransport.Upload)
	gw.HandlePath(http.MethodGet, "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		res := map[string]string{
			"message": "server is running 123",
		}
		json.NewEncoder(w).Encode(res)
	})
}
