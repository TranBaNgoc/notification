package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/golang/protobuf/ptypes/wrappers"
//	"github.com/opentracing/opentracing-go"
//	"notification-liveshopping/logging"
//	notify "notification-liveshopping/proto/notification"
//)
//
//type Handler struct {
//	// statusMap stores the serving status of the services this Server monitors.
//	statusMap map[string]notify.HealthCheckResponse_ServingStatus
//}
//
//func NewHandler(config *Config) (*Handler, error) {
//	return &Handler{
//		statusMap: make(map[string]notify.HealthCheckResponse_ServingStatus),
//	}, nil
//}
//// Check implements `service Health`.
//func (h *Handler) Check(ctx context.Context, req *notify.HealthCheckRequest, rsp *notify.HealthCheckResponse) error {
//	// add tracing
//	logging.Logger.Debug("Thuc thi: Check")
//	span, ctx := opentracing.StartSpanFromContext(ctx, EndpointService+"/check")
//	defer span.Finish()
//
//	if req.Service == "" {
//		// check the server overall health status.
//		rsp.Status = notify.HealthCheckResponse_SERVING
//	} else {
//		if status, ok := h.statusMap[req.Service]; ok {
//			rsp.Status = status
//		} else {
//			rsp.Status = notify.HealthCheckResponse_UNKNOWN
//		}
//	}
//
//	logging.Logger.Debug("Hoan thanh: Check")
//	return nil
//}
//
//func (h *Handler) Send(ctx context.Context, req *notify.NotificationRequest, rsp *notify.NotificationReply) error {
//	logging.Logger.Debug("Thuc thi: Notify.Send")
//	span, ctx := opentracing.StartSpanFromContext(ctx, EndpointService+"/send")
//	defer span.Finish()
//
//	var badge = int(req.Badge)
//
//	notification := PushNotification{
//		Platform:         int(req.Platform),
//		Tokens:           req.Tokens,
//		Message:          req.Message,
//		Title:            req.Title,
//		Topic:            req.Topic,
//		APIKey:           req.Key,
//		Category:         req.Category,
//		Sound:            req.Sound,
//		ContentAvailable: req.ContentAvailable,
//		ThreadID:         req.ThreadID,
//		MutableContent:   req.MutableContent,
//		Image:            req.Image,
//	}
//
//	if badge > 0 {
//		notification.Badge = &badge
//	}
//
//	if req.Alert != nil {
//		notification.Alert = Alert{
//			Title:        req.Alert.Title,
//			Body:         req.Alert.Body,
//			Subtitle:     req.Alert.Subtitle,
//			Action:       req.Alert.Action,
//			ActionLocKey: req.Alert.Action,
//			LaunchImage:  req.Alert.LaunchImage,
//			LocArgs:      req.Alert.LocArgs,
//			LocKey:       req.Alert.LocKey,
//			TitleLocArgs: req.Alert.TitleLocArgs,
//			TitleLocKey:  req.Alert.TitleLocKey,
//		}
//	}
//
//	// add vao Queue
//	counts  := AddNotification(notification)
//
//	rsp.Success = &wrappers.BoolValue{Value: true}
//	rsp.Counts = &wrappers.Int32Value{Value: int32(counts)}
//
//
//	logging.Logger.Debug("Hoan thanh: Notify.Send")
//	fmt.Println("Hoan thanh")
//	return nil
//
//}