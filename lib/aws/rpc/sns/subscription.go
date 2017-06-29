// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/resource/plugin"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

/* RPC stubs for Subscription resource provider */

// SubscriptionToken is the type token corresponding to the Subscription package type.
const SubscriptionToken = tokens.Type("aws:sns/subscription:Subscription")

// SubscriptionProviderOps is a pluggable interface for Subscription-related management functionality.
type SubscriptionProviderOps interface {
    Check(ctx context.Context, obj *Subscription, property string) error
    Create(ctx context.Context, obj *Subscription) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Subscription, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Subscription, new *Subscription, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Subscription, new *Subscription, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// SubscriptionProvider is a dynamic gRPC-based plugin for managing Subscription resources.
type SubscriptionProvider struct {
    ops SubscriptionProviderOps
}

// NewSubscriptionProvider allocates a resource provider that delegates to a ops instance.
func NewSubscriptionProvider(ops SubscriptionProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &SubscriptionProvider{ops: ops}
}

func (p *SubscriptionProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(SubscriptionToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return plugin.NewCheckResponse(err), nil
    }
    var failures []error
    if failure := p.ops.Check(ctx, obj, ""); failure != nil {
        failures = append(failures, failure)
    }
    unks := req.GetUnknowns()
    if !unks["name"] {
        if failure := p.ops.Check(ctx, obj, "name"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Subscription", "name", failure))
        }
    }
    if !unks["topic"] {
        if failure := p.ops.Check(ctx, obj, "topic"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Subscription", "topic", failure))
        }
    }
    if !unks["protocol"] {
        if failure := p.ops.Check(ctx, obj, "protocol"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Subscription", "protocol", failure))
        }
    }
    if !unks["endpoint"] {
        if failure := p.ops.Check(ctx, obj, "endpoint"); failure != nil {
            failures = append(failures,
                resource.NewPropertyError("Subscription", "endpoint", failure))
        }
    }
    if len(failures) > 0 {
        return plugin.NewCheckResponse(resource.NewErrors(failures)), nil
    }
    return plugin.NewCheckResponse(nil), nil
}

func (p *SubscriptionProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(SubscriptionToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[Subscription_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *SubscriptionProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(SubscriptionToken))
    obj, _, err := p.Unmarshal(req.GetProperties())
    if err != nil {
        return nil, err
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{Id: string(id)}, nil
}

func (p *SubscriptionProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(SubscriptionToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: plugin.MarshalProperties(
            nil, resource.NewPropertyMap(obj), plugin.MarshalOptions{}),
    }, nil
}

func (p *SubscriptionProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(SubscriptionToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("topic") {
            replaces = append(replaces, "topic")
        }
        if diff.Changed("protocol") {
            replaces = append(replaces, "protocol")
        }
        if diff.Changed("endpoint") {
            replaces = append(replaces, "endpoint")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *SubscriptionProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SubscriptionToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SubscriptionProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SubscriptionToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SubscriptionProvider) Unmarshal(
    v *pbstruct.Struct) (*Subscription, resource.PropertyMap, error) {
    var obj Subscription
    props := plugin.UnmarshalProperties(nil, v, plugin.MarshalOptions{RawResources: true})
    return &obj, props, mapper.MapIU(props.Mappable(), &obj)
}

/* Marshalable Subscription structure(s) */

// Subscription is a marshalable representation of its corresponding IDL type.
type Subscription struct {
    Name *string `lumi:"name,optional"`
    Topic resource.ID `lumi:"topic"`
    Protocol Protocol `lumi:"protocol"`
    Endpoint string `lumi:"endpoint"`
}

// Subscription's properties have constants to make dealing with diffs and property bags easier.
const (
    Subscription_Name = "name"
    Subscription_Topic = "topic"
    Subscription_Protocol = "protocol"
    Subscription_Endpoint = "endpoint"
)

/* Typedefs */

type (
    Protocol string
)

/* Constants */

const (
    ApplicationSubscription Protocol = "application"
    EmailJSONSubscription Protocol = "email-json"
    EmailSubscription Protocol = "email"
    HTTPSSubscription Protocol = "https"
    HTTSubscription Protocol = "http"
    LambdaSubscription Protocol = "lambda"
    SMSSubscription Protocol = "sms"
    SQSSubscription Protocol = "sqs"
)


