package iopipe

import (
  "context"
  "strconv"
  "github.com/iopipe/iopipe-go"
  "github.com/arienmalec/alexa-go"

)

type LambdaHandler func (context.Context, alexa.Request) (alexa.Response, error)

var (
  agent = iopipe.NewAgent(iopipe.Config{})
)

func WrapHandler(handler LambdaHandler) (interface{}) {
  return agent.WrapHandler(handler)
}

func Tag(ctx context.Context, req alexa.Request) {
   context, _ := iopipe.FromContext(ctx)
   intent_name := req.Body.Intent.Name
   if ( intent_name == "" ) {
     intent_name = "(blank)"
   }

   context.IOpipe.Label(intent_name)

   //Metrics
   context.IOpipe.Metric("version", req.Version)
   context.IOpipe.Metric("type", req.Body.Type)
   context.IOpipe.Metric("Timestamp", req.Body.Timestamp)
   context.IOpipe.Metric("intent_name", req.Body.Intent.Name)
   context.IOpipe.Metric("RequestID", req.Body.RequestID)
   context.IOpipe.Metric("session.isNew", strconv.FormatBool(req.Session.New))
   context.IOpipe.Metric("session.ID", req.Session.SessionID)
   context.IOpipe.Metric("session.AppId", req.Session.Application.ApplicationID)
   context.IOpipe.Metric("user.ID", req.Session.User.UserID)
   context.IOpipe.Metric("device.ID", req.Context.System.Device.DeviceID)

}
