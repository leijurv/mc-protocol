// Copyright 2015 Matthew Collins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate protocol-builder $GOFILE Status serverbound

package protocol

// StatusRequest is sent by the client instantly after
// switching to the Status protocol state and is used
// to signal the server to send a StatusResponse to the
// client
//
// This is a Minecraft packet
// ID: 0x00
type StatusRequest struct {
}

// StatusPing is sent by the client after recieving a
// StatusResponse. The client uses the time from sending
// the ping until the time of recieving a pong to measure
// the latency between the client and the server.
//
// This is a Minecraft packet
// ID: 0x01
type StatusPing struct {
	// The time when the ping was sent
	Time int64
}
