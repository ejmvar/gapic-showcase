// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package client_test

import (
	"context"

	client "github.com/googleapis/gapic-showcase/client"
	genprotopb "github.com/googleapis/gapic-showcase/server/genproto"
	"google.golang.org/api/iterator"
)

func ExampleNewTestingClient() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleTestingClient_CreateSession() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &genprotopb.CreateSessionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateSession(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_GetSession() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &genprotopb.GetSessionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetSession(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_ListSessions() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &genprotopb.ListSessionsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListSessions(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleTestingClient_DeleteSession() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &genprotopb.DeleteSessionRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteSession(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleTestingClient_ReportSession() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &genprotopb.ReportSessionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ReportSession(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTestingClient_ListTests() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &genprotopb.ListTestsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListTests(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleTestingClient_DeleteTest() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &genprotopb.DeleteTestRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteTest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleTestingClient_VerifyTest() {
	ctx := context.Background()
	c, err := client.NewTestingClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &genprotopb.VerifyTestRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.VerifyTest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
