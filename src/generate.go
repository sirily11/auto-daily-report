package src

//go:generate wire ./wire/wire.go

//go:generate mockgen -source=./repositories/merchant/merchant_repository.go -destination=./repositories/merchant/mock_merchant_repository.go -package=merchant
//go:generate mockgen -source=./services/merchant/merchant_service.go -destination=./services/merchant/mock_merchant_service.go -package=services

//go:generate mockgen -source=./repositories/item/item_repository.go -destination=./repositories/item/mock_item_repository.go -package=item
//go:generate mockgen -source=./services/item/item_service.go -destination=./services/item/mock_item_service.go -package=services

//go:generate mockgen -source=./repositories/authentication/auth_repository.go -destination=./repositories/authentication/mock_auth_repository.go -package=auth
//go:generate mockgen -source=./services/authentication/auth_service.go -destination=./services/authentication/mock_auth_service.go -package=services

//go:generate mockgen -source=./repositories/user/user_repository.go -destination=./repositories/user/mock_user_repository.go -package=user
//go:generate mockgen -source=./services/user/user_service.go -destination=./services/user/mock_user_service.go -package=services

//go:generate mockgen -source=./repositories/order/order_repository.go -destination=./repositories/order/mock_order_repository.go -package=order
//go:generate mockgen -source=./services/order/order_service.go -destination=./services/order/mock_order_service.go -package=services

//go:generate mockgen -source=./repositories/payment/payment_repository.go -destination=./repositories/payment/mock_payment_repository.go -package=payment
//go:generate mockgen -source=./services/payment/payment_service.go -destination=./services/payment/mock_payment_service.go -package=services

//go:generate mockgen -source=./services/chatroom/chatroom_service.go -destination=./services/chatroom/mock_chatroom_service.go -package=chatroom
//go:generate mockgen -source=./repositories/chatroom/chatroom_repository.go -destination=./repositories/chatroom/mock_chatroom_repository.go -package=chatroom

//go:generate mockgen -source=./services/s3/s3_service.go -destination=./services/s3/mock_s3_service.go -package=s3
//go:generate mockgen -source=./services/s3/s3.go -destination=./services/s3/mock_s3.go -package=s3
//go:generate mockgen -source=./services/s3/s3_interface.go -destination=./services/s3/mock_s3_interface.go -package=s3
