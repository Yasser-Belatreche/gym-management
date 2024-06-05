package memberships

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/components"
	bills2 "gym-management/src/components/memberships/core/usecases/bills"
	"gym-management/src/components/memberships/core/usecases/bills/get_bill"
	"gym-management/src/components/memberships/core/usecases/bills/get_bills"
	"gym-management/src/components/memberships/core/usecases/bills/mark_bill_as_paid"
	customers2 "gym-management/src/components/memberships/core/usecases/customers"
	"gym-management/src/components/memberships/core/usecases/customers/change_customer_plan"
	"gym-management/src/components/memberships/core/usecases/customers/create_customer"
	"gym-management/src/components/memberships/core/usecases/customers/delete_customer"
	"gym-management/src/components/memberships/core/usecases/customers/get_customer"
	"gym-management/src/components/memberships/core/usecases/customers/get_customers"
	"gym-management/src/components/memberships/core/usecases/customers/restrict_customer"
	"gym-management/src/components/memberships/core/usecases/customers/unrestrict_customer"
	"gym-management/src/components/memberships/core/usecases/customers/update_customer"
	memberships2 "gym-management/src/components/memberships/core/usecases/memberships"
	"gym-management/src/components/memberships/core/usecases/memberships/cancel_membership"
	"gym-management/src/components/memberships/core/usecases/memberships/get_membership"
	"gym-management/src/components/memberships/core/usecases/memberships/get_memberships"
	"gym-management/src/components/memberships/core/usecases/memberships/renew_membership"
	plans2 "gym-management/src/components/memberships/core/usecases/plans"
	"gym-management/src/components/memberships/core/usecases/plans/create_plan"
	"gym-management/src/components/memberships/core/usecases/plans/delete_plan"
	"gym-management/src/components/memberships/core/usecases/plans/get_plan"
	"gym-management/src/components/memberships/core/usecases/plans/get_plans"
	"gym-management/src/components/memberships/core/usecases/plans/update_plan"
	training_sessions2 "gym-management/src/components/memberships/core/usecases/training_sessions"
	"gym-management/src/components/memberships/core/usecases/training_sessions/end_training_session"
	"gym-management/src/components/memberships/core/usecases/training_sessions/get_training_session"
	"gym-management/src/components/memberships/core/usecases/training_sessions/get_training_sessions"
	"gym-management/src/components/memberships/core/usecases/training_sessions/start_training_session"
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/base"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/bills"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/customers"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/memberships"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/plans"
	"gym-management/src/web/gin/v1/controllers/memberships/contracts/training_sessions"
	"gym-management/src/web/gin/v1/utils"
	"net/http"
)

func CreatePlanHandler(c *gin.Context) {
	var url plans.CreatePlanUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request plans.CreatePlanRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Memberships().CreatePlan(&create_plan.CreatePlanCommand{
		Name:            request.Name,
		Featured:        *request.Featured,
		SessionsPerWeek: request.SessionsPerWeek,
		WithCoach:       *request.WithCoach,
		MonthlyPrice:    request.MonthlyPrice,
		GymId:           url.GymId,
		Session:         utils.ExtractUserSession(c),
	})

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, plans.CreatePlanResponse{Id: response.Id})
}

func UpdatePlanHandler(c *gin.Context) {
	var url plans.UpdatePlanUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request plans.UpdatePlanRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Memberships().UpdatePlan(&update_plan.UpdatePlanCommand{
		Id:              url.PlanId,
		Name:            request.Name,
		Featured:        *request.Featured,
		SessionsPerWeek: request.SessionsPerWeek,
		WithCoach:       *request.WithCoach,
		MonthlyPrice:    request.MonthlyPrice,
		Session:         utils.ExtractUserSession(c),
	})

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, plans.UpdatePlanResponse{Id: response.Id})
}

func DeletePlanHandler(c *gin.Context) {
	var url plans.DeletePlanUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().DeletePlan(&delete_plan.DeletePlanCommand{
		Id:      url.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, plans.DeletePlanResponse{Id: res.Id})
}

func GetPlansHandler(c *gin.Context) {
	var url plans.GetPlansUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request plans.GetPlansRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetPlans(&get_plans.GetPlansQuery{
		PaginatedQuery: application_specific.PaginatedQuery{
			Page:    request.Page,
			PerPage: request.PerPage,
		},
		Id:       request.Id,
		GymId:    []string{url.GymId},
		Featured: request.Featured,
		Deleted:  request.Deleted,
		Session:  utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.GetHttpPaginatedResponse(res, func(x plans2.PlanToReturn) base.Plan {
		return base.Plan{
			Id:             x.Id,
			Name:           x.Name,
			Featured:       x.Featured,
			SessionPerWeek: x.SessionPerWeek,
			WithCoach:      x.WithCoach,
			MonthlyPrice:   x.MonthlyPrice,
			GymId:          x.GymId,
			CreatedBy:      x.CreatedBy,
			UpdatedBy:      x.UpdatedBy,
			CreatedAt:      x.CreatedAt,
			UpdatedAt:      x.UpdatedAt,
			DeletedBy:      x.DeletedBy,
			DeletedAt:      x.DeletedAt,
		}
	}))
}

func GetPlanHandler(c *gin.Context) {
	var url plans.GetPlanUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetPlan(&get_plan.GetPlanQuery{
		Id:      url.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, base.Plan{
		Id:             res.Id,
		Name:           res.Name,
		Featured:       res.Featured,
		SessionPerWeek: res.SessionPerWeek,
		WithCoach:      res.WithCoach,
		MonthlyPrice:   res.MonthlyPrice,
		GymId:          res.GymId,
		CreatedBy:      res.CreatedBy,
		UpdatedBy:      res.UpdatedBy,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
		DeletedBy:      res.DeletedBy,
		DeletedAt:      res.DeletedAt,
	})
}

func CancelMembershipHandler(c *gin.Context) {
	var url memberships.CancelMembershipUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().CancelMembership(&cancel_membership.CancelMembershipCommand{
		MembershipId: url.Id,
		Session:      utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, memberships.CancelMembershipResponse{Id: res.Id})
}

func RenewMembershipHandler(c *gin.Context) {
	var url memberships.RenewMembershipUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request memberships.RenewMembershipRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().RenewMembership(&renew_membership.RenewMembershipCommand{
		MembershipId: url.Id,
		EndDate:      request.EndDate,
		Session:      utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, memberships.RenewMembershipResponse{Id: res.Id})
}

func GetMembershipsHandler(c *gin.Context) {
	var url memberships.GetMembershipsUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetMemberships(&get_memberships.GetMembershipsQuery{
		PaginatedQuery: application_specific.PaginatedQuery{
			Page:    url.Page,
			PerPage: url.PerPage,
		},
		Id:         url.Id,
		GymId:      []string{url.GymId},
		CustomerId: url.CustomerId,
		PlanId:     url.PlanId,
		Enabled:    url.Enabled,
		Session:    utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.GetHttpPaginatedResponse(res, func(x memberships2.MembershipToReturn) base.Membership {
		return base.Membership{
			Id:              x.Id,
			StartDate:       x.StartDate,
			EndDate:         x.EndDate,
			Enabled:         x.Enabled,
			DisabledFor:     x.DisabledFor,
			SessionsPerWeek: x.SessionsPerWeek,
			WithCoach:       x.WithCoach,
			MonthlyPrice:    x.MonthlyPrice,
			Customer: base.MembershipCustomer{
				Id:        x.Customer.Id,
				FirstName: x.Customer.FirstName,
				LastName:  x.Customer.LastName,
			},
			Plan: base.MembershipPlan{
				Id:   x.Plan.Id,
				Name: x.Plan.Name,
			},
			GymId:     x.GymId,
			CreatedBy: x.CreatedBy,
			CreatedAt: x.CreatedAt,
			UpdatedBy: x.UpdatedBy,
			UpdatedAt: x.UpdatedAt,
			RenewedAt: x.RenewedAt,
		}
	}))
}

func GetMembershipHandler(c *gin.Context) {
	var url memberships.GetMembershipUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetMembership(&get_membership.GetMembershipQuery{
		MembershipId: url.Id,
		Session:      utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, base.Membership{
		Id:              res.Id,
		StartDate:       res.StartDate,
		EndDate:         res.EndDate,
		Enabled:         res.Enabled,
		DisabledFor:     res.DisabledFor,
		SessionsPerWeek: res.SessionsPerWeek,
		WithCoach:       res.WithCoach,
		MonthlyPrice:    res.MonthlyPrice,
		Customer: base.MembershipCustomer{
			Id:        res.Customer.Id,
			FirstName: res.Customer.FirstName,
			LastName:  res.Customer.LastName,
		},
		Plan: base.MembershipPlan{
			Id:   res.Plan.Id,
			Name: res.Plan.Name,
		},
		GymId:     res.GymId,
		CreatedBy: res.CreatedBy,
		CreatedAt: res.CreatedAt,
		UpdatedBy: res.UpdatedBy,
		UpdatedAt: res.UpdatedAt,
		RenewedAt: res.RenewedAt,
	})
}

func MarkBillAsPaid(c *gin.Context) {
	var url bills.MarkBillAsPaidUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().MarkBillAsPaid(&mark_bill_as_paid.MarkBillAsPaidCommand{
		BillId:       url.BillId,
		MembershipId: url.MembershipId,
		Session:      utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, bills.MarkBillAsPaidResponse{Id: res.Id})
}

func GetBillsHandler(c *gin.Context) {
	var url bills.GetBillsUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetBills(&get_bills.GetBillsQuery{
		PaginatedQuery: application_specific.PaginatedQuery{
			Page:    url.Page,
			PerPage: url.PerPage,
		},
		Id:           url.BillId,
		MembershipId: []string{url.MembershipId},
		CustomerId:   url.CustomerId,
		GymId:        []string{url.GymId},
		PlanId:       url.PlanId,
		Paid:         url.Paid,
		Session:      utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.GetHttpPaginatedResponse(res, func(x bills2.BillToReturn) base.Bill {
		return base.Bill{
			Id:     x.Id,
			Amount: x.Amount,
			Paid:   x.Paid,
			Customer: base.BillCustomer{
				Id:        x.Customer.Id,
				FirstName: x.Customer.FirstName,
				LastName:  x.Customer.LastName,
			},
			Plan: base.BillPlan{
				Id:   x.Plan.Id,
				Name: x.Plan.Name,
			},
			MembershipId: x.MembershipId,
			GymId:        x.GymId,
			PaidAt:       x.PaidAt,
			DueDate:      x.DueDate,
			CreatedAt:    x.CreatedAt,
		}
	}))
}

func GetBillHandler(c *gin.Context) {
	var url bills.GetBillUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetBill(&get_bill.GetBillQuery{
		BillId:  url.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, base.Bill{
		Id:     res.Id,
		Amount: res.Amount,
		Paid:   res.Paid,
		Customer: base.BillCustomer{
			Id:        res.Customer.Id,
			FirstName: res.Customer.FirstName,
			LastName:  res.Customer.LastName,
		},
		Plan: base.BillPlan{
			Id:   res.Plan.Id,
			Name: res.Plan.Name,
		},
		MembershipId: res.MembershipId,
		GymId:        res.GymId,
		PaidAt:       res.PaidAt,
		DueDate:      res.DueDate,
		CreatedAt:    res.CreatedAt,
	})
}

func StartTrainingSessionHandler(c *gin.Context) {
	var url training_sessions.StartTrainingSessionUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().StartTrainingSession(&start_training_session.StartTrainingSessionCommand{
		MembershipId: url.MembershipId,
		Session:      utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, training_sessions.StartTrainingSessionResponse{Id: res.Id})
}

func EndTrainingSessionHandler(c *gin.Context) {
	var url training_sessions.EndTrainingSessionUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().EndTrainingSession(&end_training_session.EndTrainingSessionCommand{
		MembershipId: url.MembershipId,
		Session:      utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, training_sessions.EndTrainingSessionResponse{Id: res.Id})
}

func GetTrainingSessionsHandler(c *gin.Context) {
	var url training_sessions.GetTrainingSessionsUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetTrainingSessions(&get_training_sessions.GetTrainingSessionsQuery{
		PaginatedQuery: application_specific.PaginatedQuery{},
		Id:             url.Id,
		CustomerId:     url.CustomerId,
		MembershipId:   []string{url.MembershipId},
		GymId:          []string{url.GymId},
		Ended:          url.Ended,
		Session:        utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.GetHttpPaginatedResponse(res, func(x training_sessions2.TrainingSessionToReturn) base.TrainingSession {
		return base.TrainingSession{
			Id:        x.Id,
			StartedAt: x.StartedAt,
			EndedAt:   x.EndedAt,
			Customer: base.TrainingSessionCustomer{
				Id:        x.Customer.Id,
				FirstName: x.Customer.FirstName,
				LastName:  x.Customer.LastName,
			},
			GymId: x.GymId,
		}
	}))
}

func GetTrainingSessionHandler(c *gin.Context) {
	var url training_sessions.GetTrainingSessionUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetTrainingSession(&get_training_session.GetTrainingSessionQuery{
		Id:      url.SessionId,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, base.TrainingSession{
		Id:        res.Id,
		StartedAt: res.StartedAt,
		EndedAt:   res.EndedAt,
		Customer: base.TrainingSessionCustomer{
			Id:        res.Customer.Id,
			FirstName: res.Customer.FirstName,
			LastName:  res.Customer.LastName,
		},
		GymId: res.GymId,
	})
}

func CreateCustomerHandler(c *gin.Context) {
	var url customers.CreateCustomerUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request customers.CreateCustomerRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Memberships().CreateCustomer(&create_customer.CreateCustomerCommand{
		FirstName:         request.FirstName,
		LastName:          request.LastName,
		Email:             request.Email,
		PhoneNumber:       request.PhoneNumber,
		BirthYear:         request.BirthYear,
		Gender:            request.Gender,
		Password:          request.Password,
		PlanId:            request.PlanId,
		MembershipEndDate: request.MembershipEndDate,
		Session:           utils.ExtractUserSession(c),
	})

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, customers.CreateCustomerResponse{
		CustomerId:   response.CustomerId,
		MembershipId: response.MembershipId,
	})
}

func UpdateCustomerHandler(c *gin.Context) {
	var url customers.UpdateCustomerUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request customers.UpdateCustomerRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Memberships().UpdateCustomer(&update_customer.UpdateCustomerCommand{
		Id:          url.Id,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		BirthYear:   request.BirthYear,
		Gender:      request.Gender,
		NewPassword: request.NewPassword,
		Session:     utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customers.UpdateCustomerResponse{Id: response.Id})
}

func ChangeCustomerPlanHandler(c *gin.Context) {
	var url customers.ChangeCustomerPlanUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request customers.ChangeCustomerPlanRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	response, err := components.Memberships().ChangeCustomerPlan(&change_customer_plan.ChangeCustomerPlanCommand{
		CustomerId: url.CustomerId,
		PlanId:     request.PlanId,
		EndDate:    request.EndDate,
		Session:    utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customers.ChangeCustomerPlanResponse{
		CustomerId:   response.CustomerId,
		MembershipId: response.MembershipId,
	})
}

func RestrictCustomerHandler(c *gin.Context) {
	var url customers.RestrictCustomerUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().RestrictCustomer(&restrict_customer.RestrictCustomerCommand{
		Id:      url.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customers.RestrictCustomerResponse{Id: res.Id})
}

func UnrestrictCustomerHandler(c *gin.Context) {
	var url customers.UnrestrictCustomerUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().UnrestrictCustomer(&unrestrict_customer.UnrestrictCustomerCommand{
		Id:      url.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customers.UnrestrictCustomerResponse{Id: res.Id})
}

func DeleteCustomerHandler(c *gin.Context) {
	var url customers.DeleteCustomerUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().DeleteCustomer(&delete_customer.DeleteCustomerCommand{
		Id:      url.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customers.DeleteCustomerResponse{Id: res.Id})
}

func GetCustomersHandler(c *gin.Context) {
	var url customers.GetCustomersUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	var request customers.GetCustomersRequest
	if err := c.ShouldBind(&request); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetCustomers(&get_customers.GetCustomersQuery{
		PaginatedQuery: application_specific.PaginatedQuery{
			Page:    request.Page,
			PerPage: request.PerPage,
		},
		Id:           request.Id,
		GymId:        []string{url.GymId},
		MembershipId: request.MembershipId,
		PlanId:       request.PlanId,
		Restricted:   request.Restricted,
		Deleted:      request.Deleted,
		Session:      utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, utils.GetHttpPaginatedResponse(res, func(x customers2.CustomerToReturn) base.Customer {
		return base.Customer{
			Id:          x.Id,
			FirstName:   x.FirstName,
			LastName:    x.LastName,
			Email:       x.Email,
			PhoneNumber: x.PhoneNumber,
			Restricted:  x.Restricted,
			BirthYear:   x.BirthYear,
			Gender:      x.Gender,
			CreatedBy:   x.CreatedBy,
			UpdatedBy:   x.UpdatedBy,
			Membership: base.CustomerMembership{
				Id:              x.Membership.Id,
				Enabled:         x.Membership.Enabled,
				SessionsPerWeek: x.Membership.SessionsPerWeek,
				WithCoach:       x.Membership.WithCoach,
				MonthlyPrice:    x.Membership.MonthlyPrice,
				Plan: base.CustomerMembershipPlan{
					Id:   x.Membership.Plan.Id,
					Name: x.Membership.Plan.Name,
				},
			},
			GymId:     x.GymId,
			CreatedAt: x.CreatedAt,
			UpdatedAt: x.UpdatedAt,
			DeletedBy: x.DeletedBy,
			DeletedAt: x.DeletedAt,
		}
	}))
}

func GetCustomerHandler(c *gin.Context) {
	var url customers.GetCustomerUrl
	if err := c.ShouldBindUri(&url); err != nil {
		utils.HandleError(c, utils.NewInvalidRequestError(err))
		return
	}

	res, err := components.Memberships().GetCustomer(&get_customer.GetCustomerQuery{
		Id:      url.Id,
		Session: utils.ExtractUserSession(c),
	})
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, base.Customer{
		Id:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		Restricted:  res.Restricted,
		BirthYear:   res.BirthYear,
		Gender:      res.Gender,
		CreatedBy:   res.CreatedBy,
		UpdatedBy:   res.UpdatedBy,
		Membership: base.CustomerMembership{
			Id:              res.Membership.Id,
			Enabled:         res.Membership.Enabled,
			SessionsPerWeek: res.Membership.SessionsPerWeek,
			WithCoach:       res.Membership.WithCoach,
			MonthlyPrice:    res.Membership.MonthlyPrice,
			Plan: base.CustomerMembershipPlan{
				Id:   res.Membership.Plan.Id,
				Name: res.Membership.Plan.Name,
			},
		},
		GymId:     res.GymId,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		DeletedBy: res.DeletedBy,
		DeletedAt: res.DeletedAt,
	})
}
