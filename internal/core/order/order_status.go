package order

type Status string

const (
	StatusPending    Status = "pending"
	StatusConfirmed  Status = "confirmed"
	StatusPreparing  Status = "preparing"
	StatusReady      Status = "ready"
	StatusDelivering Status = "delivering"
	StatusDelivered  Status = "delivered"
	StatusCancelled  Status = "cancelled"
)

// var transitions = map[Status][]Status{
// 	StatusPending:    {StatusConfirmed, StatusCancelled},
// 	StatusConfirmed:  {StatusPreparing, StatusCancelled},
// 	StatusPreparing:  {StatusReady, StatusCancelled},
// 	StatusReady:      {StatusDelivering, StatusCancelled},
// 	StatusDelivering: {StatusDelivered},
// }
