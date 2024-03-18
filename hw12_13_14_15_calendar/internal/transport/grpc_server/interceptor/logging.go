package interceptor

import (
	"context"
	"fmt"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"
)

func Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		result, err := handler(ctx, req)

		logger.UseLogger().Info(
			fmt.Sprintf("%s %s %s",
				info.FullMethod[strings.LastIndex(info.FullMethod, "/"):],
				startTime.Format(time.DateTime),
				getUserAgent(ctx),
			))

		return result, err
	}
}
func getUserAgent(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	agents, ok := md["user-agent"]
	if !ok || len(agents) == 0 {
		return ""
	}

	return fmt.Sprintf("\"%s\"", agents[0])
}
