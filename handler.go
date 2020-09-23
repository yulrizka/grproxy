package grproxy

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func handleBidiStream(ctx context.Context, ss grpc.ServerStream, stub grpcdynamic.Stub, mtd *desc.MethodDescriptor, msgFactory *dynamic.MessageFactory) error {
	upstream, err := stub.InvokeRpcBidiStream(ctx, mtd)
	if err != nil {
		return err
	}

	eg, ctx := errgroup.WithContext(ctx)
	// receive message from client send to upstream
	eg.Go(func() error {
		defer func() {
			if err := upstream.CloseSend(); err != nil {
				log.Printf("failed to close send upstream")
			}
		}()

		m := msgFactory.NewMessage(mtd.GetInputType())
		for {

			// receive from client
			err := ss.RecvMsg(m)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return fmt.Errorf("receive request from client: %w", err)
			}

			// send to upstream
			if err = upstream.SendMsg(m); err != nil {
				return fmt.Errorf("send request to upstream: %w", err)
			}
			m.Reset()
		}
	})

	eg.Go(func() error {
		for {
			// receive from upstream
			msg, err := upstream.RecvMsg()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return fmt.Errorf("receive response from upstream: %w", err)
			}

			// send to client
			if err = ss.SendMsg(msg); err != nil {
				return fmt.Errorf("send response to client: %w", err)
			}
		}
	})

	err = eg.Wait()
	return nil
}

func handleClientStream(ctx context.Context, ss grpc.ServerStream, stub grpcdynamic.Stub, mtd *desc.MethodDescriptor, msgFactory *dynamic.MessageFactory) error {
	upstream, err := stub.InvokeRpcClientStream(ctx, mtd)
	if err != nil {
		return err
	}

	var resp proto.Message
	// receive message from client and send to upstream
	m := msgFactory.NewMessage(mtd.GetInputType())
	for err == nil {
		err = ss.RecvMsg(m)
		if err == io.EOF {
			resp, err = upstream.CloseAndReceive()
			break
		}
		if err != nil {
			return fmt.Errorf("receive request from client: %v", err)
		}

		err = upstream.SendMsg(m)
		if err == io.EOF {
			// We get EOF on send if the server says "go away"
			// We have to use CloseAndReceive to get the actual code
			resp, err = upstream.CloseAndReceive()
			break
		}
		m.Reset()
	}

	// finally, process response data
	_, ok := status.FromError(err)
	if !ok {
		// Error codes sent from the server will get printed differently below.
		// So just bail for other kinds of errors here.
		return fmt.Errorf("grpc call for %q failed: %w", mtd.GetFullyQualifiedName(), err)
	}

	// write the data back to the client
	if err = ss.SendMsg(resp); err != nil {
		return fmt.Errorf("send response to client: %w", err)
	}

	return nil
}

func handleServerStream(ctx context.Context, ss grpc.ServerStream, stub grpcdynamic.Stub, mtd *desc.MethodDescriptor, msgFactory *dynamic.MessageFactory) error {
	// get message from the client
	m := msgFactory.NewMessage(mtd.GetInputType())
	if err := ss.RecvMsg(m); err != nil {
		return fmt.Errorf("get request from client: %w", err)
	}

	// send to upstream and forward response to client
	var resp proto.Message
	upstream, err := stub.InvokeRpcServerStream(ctx, mtd, m)
	for err == nil {
		resp, err = upstream.RecvMsg()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("get response from upstream: %w", err)
		}
		if err := ss.SendMsg(resp); err != nil {
			return fmt.Errorf("send response to client: %w", err)
		}
		resp.Reset()
	}

	return nil
}
