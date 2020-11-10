package booking

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Lokaven-App/booking-service/grpc/modules/config"
	kafkaHelper "github.com/Lokaven-App/helpers/kafka"

	hostClient "github.com/Lokaven-App/booking-service/grpc/client/http/hosts"
	tourClientAPI "github.com/Lokaven-App/booking-service/grpc/client/http/packages"
	"github.com/Lokaven-App/booking-service/grpc/client/kafka"
	"github.com/Lokaven-App/booking-service/grpc/client/rabbitmq"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	messageClient "github.com/Lokaven-App/booking-service/grpc/client/messages"
	notifClient "github.com/Lokaven-App/booking-service/grpc/client/notification"
	tourClient "github.com/Lokaven-App/booking-service/grpc/client/packages"
	userClient "github.com/Lokaven-App/booking-service/grpc/client/users"
	"github.com/Lokaven-App/booking-service/grpc/helpers"
	bookingModel "github.com/Lokaven-App/booking-service/grpc/models/booking"
	"github.com/Lokaven-App/booking-service/grpc/validator"
	pb "github.com/Lokaven-App/booking-service/pb/booking"
	notifPB "github.com/Lokaven-App/proto/notification"
	_user "github.com/Lokaven-App/proto/user"
	"github.com/bearbin/go-age"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//CreateBookingByContact :
func CreateBookingByContact(ctx context.Context, db *mongo.Database, bookByContact *bookingModel.ByContact) (*pb.BookingByContactResponse, error) {
	user, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, bookByContact.UserUID)
	uid := user.UID
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != uid {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	err = validator.ValidateCreateBookingByContact(bookByContact)
	if err != nil {
		return nil, err
	}
	tour, err := tourClient.GetPackagesService(ctx, bookByContact.TourID)
	if err != nil {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}

	bookByContact.UserUID = uid

	tourListRecommendation, err := tourClient.GetPackagesRecommendationService(ctx, tour.HostID)
	if err != nil {
		return nil, status.Error(codes.Internal, helpers.SomethingWentWrong)
	}

	res, err := bookByContact.StoreBookingByContact(ctx, db)
	if err != nil {
		return nil, status.Error(codes.Internal, helpers.SomethingWentWrong)
	}

	data := bookingModel.TransformBookingByContact(res)

	return &pb.BookingByContactResponse{
		Code:               http.StatusOK,
		Title:              helpers.SuccessCreateBookingByContact,
		Data:               data,
		TourRecommendation: tourListRecommendation,
	}, nil
}

//CreateBooking :
func CreateBooking(ctx context.Context, db *mongo.Database, booking *bookingModel.Booking) (*pb.BookingResponse, error) {

	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, booking.UserUID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}

	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}

	// Get tour package by RPC
	// tour, err := tourClient.GetPackagesService(ctx, booking.TourID)
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	// }

	//get tour package by REST API
	tour, err := tourClientAPI.GetTransformTourPackageBSON(ctx, db, booking.TourID)
	if err != nil {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}

	if tour.TourID == "" {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}

	booking.TypeTour = tour.TypeTour
	booking.HostID = tour.HostID
	booking.Title = tour.Title
	booking.Location = tour.Location
	var priceList []*bookingModel.PackagePrice
	for _, prices := range tour.Price {
		dataPrice := &bookingModel.PackagePrice{
			PriceID:        prices.PriceID,
			TourID:         prices.TourID,
			Price:          prices.Price,
			KidPrice:       prices.KidPrice,
			MinKidAge:      prices.MinKidAge,
			MaxKidAge:      prices.MaxKidAge,
			MinParticipant: prices.MinParticipant,
			MaxParticipant: prices.MaxParticipant,
		}
		priceList = append(priceList, dataPrice)
	}

	booking.Prices = priceList

	var statusQuota bool

	if booking.TypeTour == "close" {
		booking.Status = helpers.OrderUnpaid
		_, err = validator.ValidateCreateCloseBooking(booking, tour)
		if err != nil {
			return nil, err
		}
		add24Hour := time.Now().AddDate(0, 0, 1)
		booking.EndDate = add24Hour
	} else if booking.TypeTour == "open" {
		if errMaxQuota := validator.CountMaxQuota(booking, ctx, db); errMaxQuota != nil {
			return nil, errMaxQuota
		}

		if errValidate := validator.ValidateCreateOpenBooking(booking, tour); errValidate != nil {
			return nil, errValidate
		}
		status, quotaFulfilled := validator.ValidateStatusOpenBooking(booking, tour, ctx, db)
		booking.Status = status.Status
		statusQuota = quotaFulfilled
	}

	generatePO, err := helpers.GeneratePONumber(ctx, db)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	generateINV := helpers.GenerateINVNumber(generatePO)

	booking.OrderNumber = generatePO
	booking.InvoiceNumber = generateINV

	for _, schedule := range tour.Schedules {
		if schedule.ScheduleID == booking.Schedules[0].ScheduleID {
			if strings.ToLower(schedule.TourStatus) != "idle" {
				return nil, status.Error(codes.InvalidArgument, "Failed booking this schedule, the tour status is not idle")
			}
		}
	}

	resBook, err := booking.CreateBooking(ctx, db)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	data := bookingModel.TransformBooking(resBook)

	//Notification new booking to host
	go NotifyNewBookingOrder(booking.HostID, resBook, tour, db)

	// send notification can already make payment if type tour is open
	if tour.TypeTour == "open" {

		if statusQuota == true {
			go NotifyCanPayment(booking.UserUID, resBook, tour, db)
		}

	} else if tour.TypeTour == "close" {

		err = tourClient.IsBooked(ctx, data.TourId, data.Schedules[0].ScheduleId)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		//publish message to rabbitmq delayed
		rmq := rabbitmq.RMQ

		dataPub, _ := json.Marshal(data)
		if errPub := rmq.PublishWithDelay(dataPub); errPub != nil {
			log.Errorf("error when publish cancellation")
			return nil, status.Error(codes.Internal, errPub.Error())
		}
		if errPub := rmq.PublishWithDelayReopen(dataPub); errPub != nil {
			log.Errorf("error when publish Reopen")
			return nil, status.Error(codes.Internal, errPub.Error())
		}

		err = kafka.Publish(dataPub, tour.TypeTour, "create_xendit_invoice")
		if err != nil {
			log.Errorf("error when publish kafka create_xendit_invoice")
			return nil, err
		}

	}

	return &pb.BookingResponse{
		Code:  http.StatusOK,
		Title: helpers.SuccessCreateBooking,
		Data:  data,
	}, nil
}

//UpdateStatusBooking :
func UpdateStatusBooking(ctx context.Context, db *mongo.Database, booking *pb.Booking) (*pb.StatusUpdateResponse, error) {
	if booking.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, helpers.CantGetBookingUID)
	}
	_, errExist := bookingModel.FindOrdersByID(ctx, db, booking.OrderId)
	if errExist != nil {
		return nil, errExist
	}

	var reopen = bool(false)
	update := primitive.D{}
	switch booking.Status {
	case helpers.OrderAwaiting:
		update = primitive.D{{"status", helpers.OrderAwaiting}}
	case helpers.OrderUnpaid:
		update = primitive.D{{"status", helpers.OrderUnpaid}}
	case helpers.OrderPaid:
		update = primitive.D{{"status", helpers.OrderPaid}}
	case helpers.OrderCancel:

		detailBooking, _ := bookingModel.FindOrdersByID(ctx, db, booking.OrderId)
		tour, err := tourClientAPI.GetTransformTourPackageBSON(ctx, db, detailBooking.TourID)

		if err != nil || tour.TourID == "" {
			return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
		}

		params := primitive.D{{
			"$in", []string{helpers.OrderPaid, helpers.OrderAwaiting, helpers.OrderUnpaid},
		}}
		totalParticipant := bookingModel.CountParamParticipants(ctx, detailBooking.TourID, detailBooking.Schedules[0].ScheduleID, params, db)

		for _, val := range tour.Schedules {
			if val.ScheduleID == detailBooking.Schedules[0].ScheduleID && strings.ToLower(val.TourStatus) == "idle" {
				maxQuota, _ := strconv.Atoi(val.MaxQuota)
				if totalParticipant == maxQuota {
					errUpdate := tourClient.IsUnBooked(ctx, tour.TourID, val.ScheduleID)
					if errUpdate != nil {
						log.Println(errUpdate.Error())
					}
					reopen = true
				}
				break
			}
		}
		id := primitive.NewObjectID()
		cancelAt := id.Timestamp().Local()
		update = primitive.D{{"status", helpers.OrderCancel}, {"reason", booking.Reason}, {"cancel_at", cancelAt}}

	default:
		return nil, status.Error(codes.InvalidArgument, helpers.SomethingWentWrong)
	}

	order := bookingModel.Booking{}
	order.BookingOrdersUID = booking.OrderId
	if err := order.UpdateStatusBooking(ctx, update, db); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.StatusUpdateResponse{
		Code:   http.StatusOK,
		Title:  helpers.SuccessCancelBooking,
		Reopen: reopen,
	}, nil
}

//CancelTourHost :
func CancelTourHost(ctx context.Context, db *mongo.Database, request *pb.RequestCancelTourHost) (*pb.ResponseCancelTourHost, error) {

	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}

	cfg := config.GetConfig()

	opts := options.Find()
	opts.SetSort(bson.M{"created_at": -1})
	orders, err := bookingModel.FilterBooking(&primitive.M{
		"host_id":               request.HostId,
		"schedules.schedule_id": request.ScheduleId,
		"type_tour":             "open",
		"status":                "paid",
	}, opts, ctx, db)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return &pb.ResponseCancelTourHost{
			Code:  http.StatusNotFound,
			Title: helpers.DataNotFound,
		}, nil
	}

	if request.Cancel == true {

		arrayId := []primitive.ObjectID{}

		for _, val := range orders {
			amount, _ := strconv.Atoi(val.TotalPaidPrice)
			_, err := userClient.Refund(ctx, &_user.LogsBalance{
				UserUid: val.UserUID,
				Amount:  int32(amount),
				Type:    "Refund",
				Message: val.Title,
			})
			if err != nil {
				return nil, status.Error(codes.Internal, helpers.SomethingWentWrong)
			}
			bookingId, errObj := primitive.ObjectIDFromHex(val.BookingOrdersUID)
			if errObj != nil {
				log.Error(errObj.Error())
			}
			arrayId = append(arrayId, bookingId)

			// Publish kafka notification & create meta
			msg := kafkaHelper.Init(kafkaHelper.Config{
				Username: cfg.KafkaUsername,
				Password: cfg.KafkaPassword,
				Url:      cfg.KafkaURL,
				Topic:    "notification_order_refund",
			})
			msg.Notification = notifPB.Notification{
				Service: &notifPB.Service{
					Name: "booking_service",
					Id:   val.TourID,
				},
				Body: &notifPB.Body{
					Message: "Booking Anda untuk jadwal pada " + val.Title + " dibatalkan dan di refund oleh Host karena tidak memenuhi quota",
					Sender:  val.HostID,
					Time:    time.Now().Format("15:04"),
					Status:  "Booking Order Cancellation",
					Properties: &notifPB.Properties{
						ScheduleId: val.Schedules[0].ScheduleID,
						OrderId:    val.BookingOrdersUID,
					},
				},
				Participants: []*notifPB.Participants{
					&notifPB.Participants{
						UserId:   val.UserUID,
						DeviceId: "#",
					},
				},
			}
			if errNotif := msg.Publish(ctx, nil, kafkaHelper.MessageNotification); errNotif != nil {
				log.Infof(errNotif.Error())
			}

			// Create notification in service notif
			_, errCreateNotif := notifClient.CreateNotification(ctx, isValidUID.UID, &msg.Notification)
			if errCreateNotif != nil {
				log.Infof(errCreateNotif.Error())
			}

		}
		filter := primitive.D{{"_id", primitive.D{{"$in", arrayId}}}}
		id := primitive.NewObjectID()
		cancelAt := id.Timestamp().Local()
		update := primitive.D{{"$set", primitive.D{
			{"status", helpers.OrderCancel},
			{"cancel_at", cancelAt},
			{"is_refund", true},
			{"reason", "Cancelled by host"},
		}}}
		bookingModel.UpdateMany(ctx, arrayId, filter, update, db)

	} else {

		// TODO Update tour status to awaiting

		for _, val := range orders {

			// Publish kafka notification & create meta
			msg := kafkaHelper.Init(kafkaHelper.Config{
				Username: cfg.KafkaUsername,
				Password: cfg.KafkaPassword,
				Url:      cfg.KafkaURL,
				Topic:    "notification_order_start",
			})
			msg.Notification = notifPB.Notification{
				Service: &notifPB.Service{
					Name: "booking_service",
					Id:   val.TourID,
				},
				Body: &notifPB.Body{
					Message: val.Title,
					Sender:  val.HostID,
					Time:    time.Now().Format("15:04"),
					Status:  "Tour Started",
					Properties: &notifPB.Properties{
						ScheduleId: val.Schedules[0].ScheduleID,
						OrderId:    val.BookingOrdersUID,
					},
				},
				Participants: []*notifPB.Participants{
					&notifPB.Participants{
						UserId:   val.UserUID,
						DeviceId: "#",
					},
				},
			}
			if errNotif := msg.Publish(ctx, nil, kafkaHelper.MessageNotification); errNotif != nil {
				log.Infof(errNotif.Error())
			}

			// Create notification in service notif
			_, errCreateNotif := notifClient.CreateNotification(ctx, isValidUID.UID, &msg.Notification)
			if errCreateNotif != nil {
				log.Infof(errCreateNotif.Error())
			}

		}
	}
	return &pb.ResponseCancelTourHost{
		Code:  http.StatusOK,
		Title: helpers.SuccessCancelBooking,
	}, nil

}

//GetActivityTour : get activity details
func GetActivityTour(ctx context.Context, db *mongo.Database, request *pb.ActiveTour) (*pb.ResponseActiveTour, error) {

	order := bookingModel.Booking{
		BookingOrdersUID: request.OrderId,
	}
	orderResult, err := order.DetailBooking(ctx, db)
	if err != nil || orderResult.Status != helpers.OrderPaid {
		return nil, status.Error(codes.NotFound, helpers.OrdersNotFound)
	}

	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, orderResult.UserUID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}

	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	tour, err := tourClientAPI.GetTransformTourPackageBSON(ctx, db, orderResult.TourID)
	if err != nil || tour.TourID == "" {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}
	scheduleSelected := pb.PackageSchedule{}
	for _, schedule := range tour.Schedules {
		if schedule.ScheduleID == orderResult.Schedules[0].ScheduleID {
			startDate, _ := ptypes.TimestampProto(schedule.StartDate)
			endDate, _ := ptypes.TimestampProto(schedule.EndDate)
			scheduleSelected.ScheduleId = schedule.ScheduleID
			scheduleSelected.Durations = schedule.Durations
			scheduleSelected.StartDate = startDate
			scheduleSelected.EndDate = endDate
			scheduleSelected.TourStatus = schedule.TourStatus
			scheduleSelected.Quota = schedule.Quota
			scheduleSelected.MaxQuota = schedule.MaxQuota
			scheduleSelected.MinQuota = schedule.MinQuota
			scheduleSelected.IsActive = schedule.IsActive
			scheduleSelected.IsBooked = schedule.IsBooked
		}
	}

	addons := []*pb.PackageAddon{}
	for _, addon := range orderResult.AddOns {
		addons = append(addons, &pb.PackageAddon{
			AddonId:     addon.AddonID,
			Addon:       addon.Addon,
			Price:       addon.Price,
			OrderNumber: addon.OrderNumber,
			TourId:      addon.TourID,
		})
	}

	totalParticipant := bookingModel.CountParticipants(ctx, orderResult.TourID, scheduleSelected.ScheduleId, db)

	detailHost, err := hostClient.GetDetailHost(orderResult.HostID)
	if err != nil {
		return nil, status.Error(codes.Internal, helpers.SomethingWentWrong)
	}

	message := []*pb.BroadcastMessage{}
	getMessage, err := messageClient.GetMessageDetails(ctx, isValidUID.UID, orderResult.TourID, scheduleSelected.ScheduleId)
	if err != nil {
		return nil, err
	}

	for _, m := range getMessage.Data {
		objectID, _ := primitive.ObjectIDFromHex(m.Id)
		getTimestamp := objectID.Timestamp()
		createdAt, _ := ptypes.TimestampProto(getTimestamp)
		message = append(message, &pb.BroadcastMessage{
			Host: &pb.UserInfo{
				Name:  detailHost.BusinessName,
				Image: isValidUID.ProfilePictureURL,
			},
			Message:   m.Message,
			CreatedAt: createdAt,
		})
	}

	review := wrappers.BoolValue{
		Value: orderResult.IsReviewed,
	}

	medias := []*pb.PackageMedia{}
	for _, v := range tour.Medias {
		medias = append(medias, &pb.PackageMedia{
			MediaId: v.MediaID,
			Url:     v.URL,
			Type:    v.Type,
			TourId:  v.TourID,
		})
	}

	if scheduleSelected.TourStatus == "ongoing" || scheduleSelected.TourStatus == "awaiting" {

		result := pb.ActiveTour{}
		result = pb.ActiveTour{
			OrderId:          orderResult.BookingOrdersUID,
			Medias:           medias,
			Title:            tour.Title,
			Schedules:        &scheduleSelected,
			Location:         tour.Location,
			Description:      tour.Description,
			Itinerary:        tour.Itinerary,
			Addons:           addons,
			TotalParticipant: int32(totalParticipant),
			TourStatus:       scheduleSelected.TourStatus,
			Host: &pb.UserInfo{
				Name:  detailHost.BusinessName,
				Image: isValidUID.ProfilePictureURL,
			},
			Message: message,
		}
		return &pb.ResponseActiveTour{
			Code:   http.StatusOK,
			Title:  helpers.SuccessGetData,
			Review: &review,
			Data:   &result,
		}, nil
	} else if scheduleSelected.TourStatus == "ended" {
		result := pb.ActiveTour{}
		endDate, _ := ptypes.TimestampProto(orderResult.EndDate)
		result = pb.ActiveTour{
			OrderId:          orderResult.BookingOrdersUID,
			Medias:           medias,
			Title:            tour.Title,
			Schedules:        &scheduleSelected,
			Location:         tour.Location,
			Description:      tour.Description,
			Itinerary:        tour.Itinerary,
			TotalParticipant: int32(totalParticipant),
			TourStatus:       scheduleSelected.TourStatus,
			EndedByHost:      endDate, //still dummy , must finish ticket confirmation of ended the tour
			Host: &pb.UserInfo{
				Name:  detailHost.BusinessName,
				Image: isValidUID.ProfilePictureURL,
			},
			Message: message,
		}
		return &pb.ResponseActiveTour{
			Code:   http.StatusOK,
			Title:  helpers.SuccessGetData,
			Review: &review,
			Data:   &result,
		}, nil
	}
	return nil, status.Error(codes.NotFound, helpers.OrdersNotFound)
}

//ListBooking : List booking for guest
func ListBooking(ctx context.Context, db *mongo.Database, request *pb.RequestParams) (*pb.ResponseListBooking, error) {
	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	page, _ := strconv.Atoi(request.Page)
	perPage, _ := strconv.Atoi(request.PerPage)
	result := []*pb.ListBooking{}
	paginate := pb.Paginate{}
	opts := options.Find()
	opts.SetSort(bson.M{"created_at": -1})
	switch request.Param {
	case "cancelled":
		total, err := bookingModel.FilterBooking(&primitive.M{
			"user_uid": header.UID,
			"$or": []bson.M{
				{"status": "cancelled"},
				{"is_refund": true},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paging, offset := helpers.Pagination(page, perPage, len(total))
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(perPage))
		orders, err := bookingModel.FilterBooking(&primitive.M{
			"user_uid": header.UID,
			"$or": []bson.M{
				{"status": "cancelled"},
				{"is_refund": true},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paginate = *paging
		result = bookingModel.TransformListBooking(ctx, orders, db)
	case "tour":
		total, err := bookingModel.FilterBooking(&primitive.M{
			"user_uid": header.UID,
			"status":   "paid",
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paging, offset := helpers.Pagination(page, perPage, len(total))
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(perPage))
		orders, err := bookingModel.FilterBooking(&primitive.M{
			"user_uid": header.UID,
			"status":   "paid",
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paginate = *paging
		result = bookingModel.TransformListBooking(ctx, orders, db)
	case "activebooking":
		total, err := bookingModel.FilterBooking(&primitive.M{
			"user_uid": header.UID,
			"status": bson.M{
				"$in": []string{"paid", "unpaid", "awaiting more bookings"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paging, offset := helpers.Pagination(page, perPage, len(total))
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(perPage))
		orders, err := bookingModel.FilterBooking(&primitive.M{
			"user_uid": header.UID,
			"status": bson.M{
				"$in": []string{"paid", "unpaid", "awaiting more bookings"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paginate = *paging
		result = bookingModel.TransformActiveListBooking(ctx, orders, db)
	case "activetour":
		total, err := bookingModel.FilterBooking(&primitive.M{
			"user_uid": header.UID,
			"schedules.tour_status": bson.M{
				"$in": []string{"awaiting", "ongoing"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paging, offset := helpers.Pagination(page, perPage, len(total))
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(perPage))
		orders, err := bookingModel.FilterBooking(&primitive.M{
			"user_uid": header.UID,
			"schedules.tour_status": bson.M{
				"$in": []string{"awaiting", "ongoing"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paginate = *paging
		result = bookingModel.TransformActiveListBooking(ctx, orders, db)
	default:
		return &pb.ResponseListBooking{
			Code:  http.StatusNotFound,
			Title: helpers.DataNotFound,
		}, nil
	}
	if len(result) > 0 {
		return &pb.ResponseListBooking{
			Code:     http.StatusOK,
			Title:    helpers.SuccessGetData,
			Data:     result,
			Paginate: &paginate,
		}, nil
	}
	return &pb.ResponseListBooking{
		Code:  http.StatusNotFound,
		Title: helpers.DataNotFound,
	}, nil
}

// TotalParticipant :
func TotalParticipant(ctx context.Context, db *mongo.Database, request *pb.TotalParticipants) (*pb.ResponseTotalParticipant, error) {
	totalParticipant := bookingModel.CountParticipants(ctx, request.TourId, request.ScheduleId, db)
	return &pb.ResponseTotalParticipant{
		Code:  http.StatusOK,
		Title: helpers.SuccessGetData,
		Data: &pb.TotalParticipants{
			Total: int32(totalParticipant),
		},
	}, nil
}

// GetListBokings :
func GetListBokings(ctx context.Context, db *mongo.Database, request *pb.RequestInfoBookings) (*pb.ResponseListInfoBookings, error) {
	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}

	result, err := bookingModel.FilterBooking(&primitive.M{
		"tour_id":               request.TourId,
		"schedules.schedule_id": request.ScheduleId,
		"status":                request.Status,
	}, nil, ctx, db)

	if err != nil {
		return nil, status.Error(codes.NotFound, helpers.DataNotFound)
	}

	listBookings := []*pb.Booking{}
	for _, v := range result {
		listBookings = append(listBookings, &pb.Booking{
			UserUid: v.UserUID,
		})
	}

	return &pb.ResponseListInfoBookings{
		Code:  http.StatusOK,
		Title: helpers.SuccessGetData,
		Data:  listBookings,
	}, nil
}

// ListScheduleHost :
func ListScheduleHost(ctx context.Context, db *mongo.Database, request *pb.RequestParams) (*pb.ResponseListScheduleHost, error) {

	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}

	page, _ := strconv.Atoi(request.Page)
	perPage, _ := strconv.Atoi(request.PerPage)
	paginate := pb.Paginate{}
	listSchedule := []*pb.ListScheduleHost{}
	result := []*pb.ListScheduleHost{}

	if request.TypeTour == "close" {

		order := bookingModel.Booking{
			HostID:   request.HostId,
			TypeTour: request.TypeTour,
			Schedules: []*bookingModel.PackageSchedules{
				&bookingModel.PackageSchedules{ScheduleID: request.ScheduleId},
			},
		}

		orders, err := order.FindBookingBySchedule(ctx, db)
		if err != nil {
			return &pb.ResponseListScheduleHost{
				Code:  http.StatusNotFound,
				Title: helpers.DataNotFound,
			}, nil
		}

		if request.Param == "pending" && orders.Status != "unpaid" {
			return &pb.ResponseListScheduleHost{
				Code:  http.StatusNotFound,
				Title: helpers.DataNotFound,
			}, nil
		} else if request.Param == "paid" && orders.Status != "paid" {
			return &pb.ResponseListScheduleHost{
				Code:  http.StatusNotFound,
				Title: helpers.DataNotFound,
			}, nil
		} else if request.Param == "cancelled" && orders.Status != "cancelled" {
			return &pb.ResponseListScheduleHost{
				Code:  http.StatusNotFound,
				Title: helpers.DataNotFound,
			}, nil
		}

		tour, err := tourClientAPI.GetTransformTourPackageBSON(ctx, db, orders.TourID)
		for _, val := range tour.Schedules {
			if val.ScheduleID == orders.Schedules[0].ScheduleID && !val.IsBooked {
				return &pb.ResponseListScheduleHost{
					Code:  http.StatusNotFound,
					Title: helpers.DataNotFound,
				}, nil
			}
		}

		res, err := bookingModel.TransformListScheduleHost(ctx, orders, db)
		if err != nil {
			return &pb.ResponseListScheduleHost{
				Code:  http.StatusInternalServerError,
				Title: helpers.SomethingWentWrong,
			}, nil
		}
		result = append(listSchedule, res)

	} else if request.TypeTour == "open" {

		opts := options.Find()
		opts.SetSort(bson.M{"created_at": -1})
		query := primitive.M{}

		if request.Param == "pending" {
			query = primitive.M{
				"host_id":               request.HostId,
				"schedules.schedule_id": request.ScheduleId,
				"type_tour":             request.TypeTour,
				"status": bson.M{
					"$in": []string{"unpaid", "awaiting more booking"},
				}}
		} else if request.Param == "paid" {
			query = primitive.M{
				"host_id":               request.HostId,
				"schedules.schedule_id": request.ScheduleId,
				"type_tour":             request.TypeTour,
				"status":                "paid",
			}
		} else if request.Param == "cancelled" {
			query = primitive.M{
				"host_id":               request.HostId,
				"schedules.schedule_id": request.ScheduleId,
				"type_tour":             request.TypeTour,
				"status":                "cancelled",
			}
		} else {
			return nil, status.Error(codes.InvalidArgument, helpers.InvalidArgument)
		}

		totalOrder, err := bookingModel.FilterBooking(&query, opts, ctx, db)
		if err != nil {
			return nil, err
		}

		paging, offset := helpers.Pagination(page, perPage, len(totalOrder))
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(perPage))
		paginate = *paging

		orders, err := bookingModel.FilterBooking(&query, opts, ctx, db)
		if err != nil {
			return nil, err
		}

		for _, order := range orders {
			res, err := bookingModel.TransformListScheduleHost(ctx, order, db)
			if err != nil {
				return &pb.ResponseListScheduleHost{
					Code:  http.StatusInternalServerError,
					Title: helpers.SomethingWentWrong,
				}, nil
			}
			listSchedule = append(listSchedule, res)
		}
		result = listSchedule
	}

	if len(result) > 0 {
		return &pb.ResponseListScheduleHost{
			Code:     http.StatusOK,
			Title:    helpers.SuccessGetData,
			Data:     result,
			Paginate: &paginate,
		}, nil
	}
	return &pb.ResponseListScheduleHost{
		Code:  http.StatusNotFound,
		Title: helpers.DataNotFound,
	}, nil
}

//DetailBooking :
func DetailBooking(ctx context.Context, db *mongo.Database, booking *pb.Booking) (*pb.BookingResponse, error) {
	if booking.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, helpers.CantGetBookingUID)
	}
	order := bookingModel.Booking{}
	order.BookingOrdersUID = booking.OrderId
	resBook, err := order.DetailBooking(ctx, db)
	if err != nil {
		return nil, err
	}
	return &pb.BookingResponse{
		Code:  http.StatusOK,
		Title: helpers.SuccessGetData,
		Data:  bookingModel.TransformBooking(resBook),
	}, nil
}

// DetailBookingGuest :
func DetailBookingGuest(ctx context.Context, db *mongo.Database, booking *pb.Booking) (*pb.ResponseDetailGuestBooking, error) {
	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	order := bookingModel.Booking{
		BookingOrdersUID: booking.OrderId,
		UserUID:          header.UID,
	}
	resultBooking, err := order.FindBookingByGuest(ctx, db)
	if err != nil {
		return nil, err
	}
	tour, err := tourClientAPI.GetTransformTourPackageBSON(ctx, db, resultBooking.TourID)
	if err != nil {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}
	if tour.TourID == "" {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}
	return bookingModel.TransformGuestDetailBooking(ctx, resultBooking, tour, db)
}

// DetailBookingHost :
func DetailBookingHost(ctx context.Context, db *mongo.Database, booking *pb.Booking) (*pb.ResponseDetailHostBooking, error) {

	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		return nil, err
	}
	if header == nil {
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	order := bookingModel.Booking{
		BookingOrdersUID: booking.OrderId,
		HostID:           booking.HostId,
	}
	resultBooking, err := order.FindBookingByHost(ctx, db)
	if err != nil {
		return nil, status.Error(codes.NotFound, helpers.OrdersNotFound)
	}
	tour, err := tourClientAPI.GetTransformTourPackageBSON(ctx, db, resultBooking.TourID)
	if err != nil {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}
	if tour.TourID == "" {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}
	return bookingModel.TransformHostDetailBooking(ctx, resultBooking, tour, db)

}

// DetailParticipants :
func DetailParticipants(ctx context.Context, db *mongo.Database, booking *pb.Booking) (*pb.ResponseParticipantsList, error) {
	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		return nil, err
	}
	if header == nil {
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	order := bookingModel.Booking{
		BookingOrdersUID: booking.OrderId,
	}
	result, err := order.DetailBooking(ctx, db)
	if err != nil {
		return nil, status.Error(codes.NotFound, helpers.OrdersNotFound)
	}
	listParticpants := []*pb.Participant{}
	for _, participant := range result.Participants {
		listParticpants = append(listParticpants, &pb.Participant{
			FirstName: participant.FirstName,
			LastName:  participant.LastName,
			Age:       participant.Age,
		})
	}
	return &pb.ResponseParticipantsList{
		Code:  http.StatusOK,
		Data:  listParticpants,
		Title: helpers.SuccessGetData,
	}, nil
}

// DetailParticipantsHost :
func DetailParticipantsHost(ctx context.Context, db *mongo.Database, request *pb.ParticipantHostRequest) (*pb.ResponseParticipantHost, error) {
	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		return nil, err
	}
	if header == nil {
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}

	page, _ := strconv.Atoi(request.Page)
	perPage, _ := strconv.Atoi(request.PerPage)

	total, err := bookingModel.FilterBooking(&primitive.M{
		"tour_id":               request.TourId,
		"schedules.schedule_id": request.ScheduleId,
	}, nil, ctx, db)
	if err != nil {
		return nil, err
	}
	paging, offset := helpers.Pagination(page, perPage, len(total))
	opts := options.Find()
	opts.SetSort(bson.M{"created_at": -1})
	opts.SetSkip(int64(offset))
	opts.SetLimit(int64(perPage))

	orders, err := bookingModel.FilterBooking(&primitive.M{
		"tour_id":               request.TourId,
		"schedules.schedule_id": request.ScheduleId,
	}, opts, ctx, db)
	if err != nil {
		return nil, err
	}

	participantsList := []*pb.ParticipantHost{}
	for _, val := range orders {
		participants := []*pb.Participant{}
		userInfo, err := userClient.GetUserService(ctx, val.UserUID)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		for _, participant := range val.Participants {
			participants = append(participants, &pb.Participant{
				FirstName: participant.FirstName,
				LastName:  participant.LastName,
				Age:       participant.Age,
			})
		}
		y, m, d := userInfo.DateOfBirth.Date()
		participantsList = append(participantsList, &pb.ParticipantHost{
			Participant: participants,
			UserUid:     val.UserUID,
			UserInfo: &pb.UserInfo{
				Name: userInfo.FirstName + ` ` + userInfo.LastName,
				Age:  fmt.Sprintf("%v", age.Age(time.Date(y, m, d, 0, 0, 0, 0, time.UTC))),
			},
		})
	}
	return &pb.ResponseParticipantHost{
		Title:    "Ok",
		Code:     http.StatusOK,
		Data:     participantsList,
		Paginate: paging,
	}, nil
}

//GetBookingReceipt :
func GetBookingReceipt(ctx context.Context, db *mongo.Database, orderNumber string) (*pb.ResponseReceiptBooking, error) {

	book, err := bookingModel.FindOrdersReceipts(ctx, db, orderNumber)
	if err != nil {
		return nil, err
	}

	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	if header == nil {
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}

	isValidUID, err := userClient.GetUserService(ctx, book.UserUID)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}

	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}

	//get tour package by REST API
	tour, err := tourClientAPI.GetTransformTourPackageBSON(ctx, db, book.TourID)
	if err != nil {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}

	var medias []*bookingModel.PackageMedias
	for _, media := range tour.Medias {
		medias = append(medias, &bookingModel.PackageMedias{
			MediaID: media.MediaID,
			TourID:  media.TourID,
			Type:    media.Type,
			URL:     media.URL,
		})
	}

	book.Medias = medias

	if book.TypeTour == "open" {
		var totalQty []int
		totalOrder, err := bookingModel.FilterBooking(&primitive.M{
			"tour_id":               book.TourID,
			"schedules.schedule_id": book.Schedules[0].ScheduleID,
			"status": bson.M{
				"$in": []string{"paid", "awaiting more bookings", "unpaid", "cancelled"},
			},
		}, nil, ctx, db)
		if err != nil {
			return nil, err
		}

		var minQuota []string
		var maxQuota []string
		for _, total := range totalOrder {
			minQuota = append(minQuota, total.Schedules[0].MinQuota)
			maxQuota = append(maxQuota, total.Schedules[0].MaxQuota)
			qtyAdultInt, _ := strconv.Atoi(total.QtyAdults)
			qtyKidInt, _ := strconv.Atoi(total.QtyKids)
			total := qtyAdultInt + qtyKidInt
			totalQty = append(totalQty, total)
		}

		resultQty := 0
		for _, n := range totalQty {
			resultQty += n
		}

		resultQtyStr := strconv.Itoa(resultQty)
		book.MinQuota = minQuota[0]
		book.MaxQuota = maxQuota[0]
		book.TotalQuota = resultQtyStr
		maxQuotaInt, _ := strconv.Atoi(maxQuota[0])
		quotaLeft := maxQuotaInt - resultQty
		if quotaLeft <= 0 {
			book.QuotaLeft = "0"
		} else if quotaLeft > 900 {
			book.QuotaLeft = "infinity"
		} else {
			quotaLeftStr := strconv.Itoa(quotaLeft)
			book.QuotaLeft = quotaLeftStr
		}
	}

	data := bookingModel.TransformToGetBookingReceipt(book)

	return &pb.ResponseReceiptBooking{
		Code:  http.StatusOK,
		Title: helpers.SuccessGetData,
		Data:  data,
	}, nil

}

//ListIcomingBooking : List booking for host
func ListIcomingBooking(ctx context.Context, db *mongo.Database, request *pb.RequestParamsIncomingList) (*pb.ResponseIncomingBookingList, error) {
	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	getHost, err := hostClient.GetHostByUID(isValidUID.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.InvalidArgument, helpers.YouAreNotaHost)
	}
	page, _ := strconv.Atoi(request.Page)
	perPage, _ := strconv.Atoi(request.PerPage)
	result := []*pb.ListIncomingBooking{}
	paginate := pb.Paginate{}
	opts := options.Find()
	opts.SetSort(bson.M{"created_at": -1})
	if request.Param == "booking" && request.Key == "" {
		total, err := bookingModel.FilterBooking(&primitive.M{
			"host_id": getHost.HostId,
			"status": bson.M{
				"$in": []string{"paid", "unpaid", "awaiting more bookings"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paging, offset := helpers.Pagination(page, perPage, len(total))
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(perPage))
		orders, err := bookingModel.FilterBooking(&primitive.M{
			"host_id": getHost.HostId,
			"status": bson.M{
				"$in": []string{"paid", "unpaid", "awaiting more bookings"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paginate = *paging
		result = bookingModel.TransformActiveIncomingBookingList(ctx, orders, db)
	} else if request.Param == "booking" && request.Key != "" {
		total, err := bookingModel.FilterBooking(&primitive.M{
			"host_id": getHost.HostId,
			"status": bson.M{
				"$in": []string{"paid", "unpaid", "awaiting more bookings"},
			},
			"$or": []bson.M{
				bson.M{"title": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"order_number": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"invoice_number": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"location": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"type_tour": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"status": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"addons.addon": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"participants.first_name": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"participants.last_name": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"participants.email": primitive.Regex{Pattern: request.Key, Options: ""}},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paging, offset := helpers.Pagination(page, perPage, len(total))
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(perPage))
		orders, err := bookingModel.FilterBooking(&primitive.M{
			"host_id": getHost.HostId,
			"status": bson.M{
				"$in": []string{"paid", "unpaid", "awaiting more bookings"},
			},
			"$or": []bson.M{
				bson.M{"title": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"order_number": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"invoice_number": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"location": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"type_tour": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"status": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"addons.addon": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"participants.first_name": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"participants.last_name": primitive.Regex{Pattern: request.Key, Options: ""}},
				bson.M{"participants.email": primitive.Regex{Pattern: request.Key, Options: ""}},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paginate = *paging
		result = bookingModel.TransformActiveIncomingBookingList(ctx, orders, db)
	}
	if len(result) > 0 {
		return &pb.ResponseIncomingBookingList{
			Code:     http.StatusOK,
			Title:    helpers.SuccessGetData,
			Data:     result,
			Paginate: &paginate,
		}, nil
	}
	return &pb.ResponseIncomingBookingList{
		Code:  http.StatusNotFound,
		Title: helpers.DataNotFound,
	}, nil
}

//GetActiveBookingList : List booking for host
func GetActiveBookingList(ctx context.Context, db *mongo.Database, request *pb.RequestParamsActiveList) (*pb.ResponseActiveBookingList, error) {
	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	getHost, err := hostClient.GetHostByUID(isValidUID.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.InvalidArgument, helpers.YouAreNotaHost)
	}
	page, _ := strconv.Atoi(request.Page)
	perPage, _ := strconv.Atoi(request.PerPage)
	result := []*pb.ListActiveBooking{}
	paginate := pb.Paginate{}
	opts := options.Find()
	opts.SetSort(bson.M{"created_at": -1})
	total, err := bookingModel.FilterBooking(&primitive.M{
		"host_id": getHost.HostId,
		"status": bson.M{
			"$in": []string{"paid"},
		},
		"schedules.tour_status": bson.M{
			"$in": []string{"ongoing", "awaiting"},
		},
	}, opts, ctx, db)
	if err != nil {
		return nil, err
	}
	paging, offset := helpers.Pagination(page, perPage, len(total))
	opts.SetSkip(int64(offset))
	opts.SetLimit(int64(perPage))
	orders, err := bookingModel.FilterBooking(&primitive.M{
		"host_id": getHost.HostId,
		"status": bson.M{
			"$in": []string{"paid"},
		},
		"schedules.tour_status": bson.M{
			"$in": []string{"ongoing", "awaiting"},
		},
	}, opts, ctx, db)
	if err != nil {
		return nil, err
	}
	paginate = *paging
	result = bookingModel.TransformActiveBookingList(ctx, orders, db)

	if len(result) > 0 {
		return &pb.ResponseActiveBookingList{
			Code:     http.StatusOK,
			Title:    helpers.SuccessGetData,
			Data:     result,
			Paginate: &paginate,
		}, nil
	}
	return &pb.ResponseActiveBookingList{
		Code:  http.StatusNotFound,
		Title: helpers.DataNotFound,
	}, nil
}

//GetBookingActivitiyList : List booking for host
func GetBookingActivitiyList(ctx context.Context, db *mongo.Database, request *pb.RequestParamsActivityList) (*pb.ResponseBookingActivityList, error) {
	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	getHost, err := hostClient.GetHostByUID(isValidUID.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.InvalidArgument, helpers.YouAreNotaHost)
	}
	page, _ := strconv.Atoi(request.Page)
	perPage, _ := strconv.Atoi(request.PerPage)
	result := []*pb.ListActivityBooking{}
	paginate := pb.Paginate{}
	opts := options.Find()
	opts.SetSort(bson.M{"created_at": -1})
	if request.Param == "ongoing" {
		total, err := bookingModel.FilterBooking(&primitive.M{
			"host_id": getHost.HostId,
			"status": bson.M{
				"$in": []string{"paid"},
			},
			"schedules.tour_status": bson.M{
				"$in": []string{"ongoing"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paging, offset := helpers.Pagination(page, perPage, len(total))
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(perPage))
		orders, err := bookingModel.FilterBooking(&primitive.M{
			"host_id": getHost.HostId,
			"status": bson.M{
				"$in": []string{"paid"},
			},
			"schedules.tour_status": bson.M{
				"$in": []string{"ongoing"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paginate = *paging
		result = bookingModel.TransformActvityBookingList(ctx, orders, db)
	} else if request.Param == "past" {
		total, err := bookingModel.FilterBooking(&primitive.M{
			"host_id": getHost.HostId,
			"status": bson.M{
				"$in": []string{"paid"},
			},
			"schedules.tour_status": bson.M{
				"$in": []string{"ended"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paging, offset := helpers.Pagination(page, perPage, len(total))
		opts.SetSkip(int64(offset))
		opts.SetLimit(int64(perPage))
		orders, err := bookingModel.FilterBooking(&primitive.M{
			"host_id": getHost.HostId,
			"status": bson.M{
				"$in": []string{"paid"},
			},
			"schedules.tour_status": bson.M{
				"$in": []string{"ended"},
			},
		}, opts, ctx, db)
		if err != nil {
			return nil, err
		}
		paginate = *paging
		result = bookingModel.TransformActvityBookingList(ctx, orders, db)
	}

	if len(result) > 0 {
		return &pb.ResponseBookingActivityList{
			Code:     http.StatusOK,
			Title:    helpers.SuccessGetData,
			Data:     result,
			Paginate: &paginate,
		}, nil
	}
	return &pb.ResponseBookingActivityList{
		Code:  http.StatusNotFound,
		Title: helpers.DataNotFound,
	}, nil
}

//GetSchedulesList :
// func GetSchedulesList(ctx context.Context, reqSchedule *pb.RequestPackageSchedule) (*pb.ResponePackageScheduleList, error) {

// 	getschedules, err := tourClient.GetPackageScheduleList(ctx, reqSchedule.TourId)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, status.Error(codes.NotFound, helpers.TourPackageScheduleNotFound)
// 	}

// 	res, err := bookingModel.FindOrders(ctx, reqSchedule.OrderNumber)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, status.Error(codes.NotFound, helpers.OrdersNotFound)
// 	}

// 	header, err := helpers.GetUserFromHeader(ctx)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
// 	if header == nil {
// 		log.Error(err)
// 		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
// 	}
// 	isValidUID, err := userClient.GetUserService(ctx, res.UserUID)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
// 	}

// 	if isValidUID.UID != header.UID {
// 		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
// 	}

// 	return &pb.ResponePackageScheduleList{
// 		Data: getschedules,
// 	}, nil
// }

//GetPricesList :
// func GetPricesList(ctx context.Context, reqPrice *pb.PackagePrice) (*pb.ResponePricelist, error) {
// 	res, err := bookingModel.FindOrders(ctx, reqPrice.OrderNumber)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, status.Error(codes.NotFound, helpers.OrdersNotFound)
// 	}

// 	header, err := helpers.GetUserFromHeader(ctx)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
// 	if header == nil {
// 		log.Error(err)
// 		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
// 	}
// 	isValidUID, err := userClient.GetUserService(ctx, res.UserUID)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
// 	}

// 	if isValidUID.UID != header.UID {
// 		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
// 	}
// 	getPrices, err := tourClient.GetPackagePricesList(ctx, reqPrice.TourId)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, status.Error(codes.NotFound, helpers.TourPackagePricesNotFound)
// 	}
// 	return &pb.ResponePricelist{
// 		Data: getPrices,
// 	}, nil
// }

//GetPriceByTotalParticipant :
func GetPriceByTotalParticipant(ctx context.Context, request *pb.RequestDetailPriceByParticipant) (*pb.ResponseDetailPriceByParticipant, error) {
	header, err := helpers.GetUserFromHeader(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if header == nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.CantGetUID)
	}
	isValidUID, err := userClient.GetUserService(ctx, header.UID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	if isValidUID.UID != header.UID {
		return nil, status.Error(codes.Unauthenticated, helpers.YouOnlyDoYourOwnAction)
	}
	getPriceByParticipant, err := tourClient.GetPackagePriceByParticipant(ctx, request.TourId, request.AdultParticipant, request.KidParticipant)
	if err != nil {
		return nil, status.Error(codes.Internal, helpers.SomethingWentWrong)
	}
	result := bookingModel.TransformPriceByParticipant(getPriceByParticipant)
	return result, nil
}

//UpdateIsReviewed :
func UpdateIsReviewed(ctx context.Context, db *mongo.Database, booking *pb.Booking) (*pb.BookingResponse, error) {
	result, err := bookingModel.UpdateIsReviewed(ctx, db, booking.OrderId, booking.TourId, booking.UserUid)
	if err != nil {
		return nil, err
	}

	data := bookingModel.TransformBooking(result)
	return &pb.BookingResponse{
		Code:  200,
		Data:  data,
		Title: "Ok",
	}, nil
}

//GetBookingBySchedulesTourID :
func GetBookingBySchedulesTourID(ctx context.Context, db *mongo.Database, tourID string, scheduleID string) (*pb.ResponseGetBookingByTourScheduleID, error) {

	book, err := bookingModel.FindOrderBySchedulesTourID(ctx, db, tourID, scheduleID)
	if err != nil {
		return nil, err
	}

	//get tour package by REST API
	_, err = tourClientAPI.GetTransformTourPackageBSON(ctx, db, book.TourID)
	if err != nil {
		return nil, status.Error(codes.NotFound, helpers.TourPackageNotFound)
	}

	var bookings []string
	if book.TypeTour == "open" {
		var totalQty []int
		totalOrder, err := bookingModel.FilterBooking(&primitive.M{
			"tour_id":               book.TourID,
			"schedules.schedule_id": book.Schedules[0].ScheduleID,
			"status": bson.M{
				"$in": []string{"paid", "awaiting more bookings", "unpaid"},
			},
		}, nil, ctx, db)
		if err != nil {
			return nil, err
		}

		var minQuota []string
		var maxQuota []string
		var qtyAdults []int
		var qtyKids []int
		for _, total := range totalOrder {
			bookings = append(bookings, total.OrderNumber)
			minQuota = append(minQuota, total.Schedules[0].MinQuota)
			maxQuota = append(maxQuota, total.Schedules[0].MaxQuota)
			qtyAdultInt, _ := strconv.Atoi(total.QtyAdults)
			qtyKidInt, _ := strconv.Atoi(total.QtyKids)
			qtyAdults = append(qtyAdults, qtyAdultInt)
			qtyKids = append(qtyKids, qtyKidInt)
			total := qtyAdultInt + qtyKidInt
			totalQty = append(totalQty, total)
		}

		resultQty := 0
		for _, n := range totalQty {
			resultQty += n
		}

		totalQtyAdult := 0
		for _, a := range qtyAdults {
			totalQtyAdult += a
		}

		totalQtyKid := 0
		for _, k := range qtyKids {
			totalQtyKid += k
		}

		book.MinQuota = minQuota[0]
		book.MaxQuota = maxQuota[0]
		book.QtyAdults = strconv.Itoa(totalQtyAdult)
		book.QtyKids = strconv.Itoa(totalQtyKid)
		book.TotalQuota = strconv.Itoa(resultQty)
		maxQuotaInt, _ := strconv.Atoi(maxQuota[0])
		quotaLeft := maxQuotaInt - resultQty
		if quotaLeft <= 0 {
			book.QuotaLeft = "0"
		} else if quotaLeft > 900 {
			book.QuotaLeft = "infinity"
		} else {
			quotaLeftStr := strconv.Itoa(quotaLeft)
			book.QuotaLeft = quotaLeftStr
		}

	} else {
		return nil, status.Error(codes.InvalidArgument, helpers.SomethingWentWrong)
	}

	data := bookingModel.TransformToGetBookingByTourScheduleID(book, bookings)

	return &pb.ResponseGetBookingByTourScheduleID{
		Code:  http.StatusOK,
		Title: helpers.SuccessGetData,
		Data:  data,
	}, nil
}

//UpdateBookingScheduleTourStatus : updaate tour status on booking schedule
func UpdateBookingScheduleTourStatus(ctx context.Context, db *mongo.Database, request *pb.RequestBookingScheduleTourStatus) (*pb.ResponseBookingScheduleTourStatus, error) {
	result, err := bookingModel.UpdateBookingTourStatus(ctx, db, request.OrderNumber, request.TourId, request.ScheduleId, request.TourStatus)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//client package to update tour status on package service
	errClient := tourClient.UpdateTourPackageScheduleTourStatus(ctx, request.TourId, request.ScheduleId, request.TourStatus)
	if errClient != nil {
		log.Error(errClient)
		return nil, errClient
	}

	data := bookingModel.TransformBooking(result)

	return &pb.ResponseBookingScheduleTourStatus{
		Data: data,
	}, nil
}
