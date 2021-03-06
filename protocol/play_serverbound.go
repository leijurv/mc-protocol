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

//go:generate protocol-builder $GOFILE Play serverbound

package protocol

// This is a Minecraft packet
// ID: 0x00
type TeleportConfirm struct {
	TeleportID VarInt
}

// TabComplete is sent by the client when the client presses tab in
// the chat box.
//
// This is a Minecraft packet
// ID: 0x01
type TabComplete struct {
	Text         string
	AssumeComman bool
	HasTarget    bool
	Target       Position `if:".HasTarget==true"`
}

// ChatMessage is sent by the client when it sends a chat message or
// executes a command (prefixed by '/').
//
// This is a Minecraft packet
// ID: 0x02
type ChatMessage struct {
	Message string
}

// ClientStatus is sent to update the client's status
//
// This is a Minecraft packet
// ID: 0x03
type ClientStatus struct {
	ActionID VarInt
}

// ClientSettings is sent by the client to update its current settings.
//
// This is a Minecraft packet
// ID: 0x04
type ClientSettings struct {
	Locale             string
	ViewDistance       byte
	ChatMode           byte
	ChatColors         bool
	DisplayedSkinParts byte
	MainHand           VarInt
}

// ConfirmTransactionServerbound is a reply to ConfirmTransaction.
//
// This is a Minecraft packet
// ID: 0x05
type ConfirmTransactionServerbound struct {
	ID           byte
	ActionNumber int16
	Accepted     bool
}

// EnchantItem is sent when the client enchants an item.
//
// This is a Minecraft packet
// ID: 0x06
type EnchantItem struct {
	ID          byte
	Enchantment byte
}

// ClickWindow is sent when the client clicks in a window.
//
// This is a Minecraft packet
// ID: 0x07
type ClickWindow struct {
	ID           byte
	Slot         int16
	Button       byte
	ActionNumber int16
	Mode         byte
	ClickedItem  ItemStack `as:"raw"`
}

// CloseWindow is sent when the client closes a window.
//
// This is a Minecraft packet
// ID: 0x08
type CloseWindow struct {
	ID byte
}

// PluginMessageServerbound is used for custom messages between the client
// and server. This is mainly for plugins/mods but vanilla has a few channels
// registered too.
//
// This is a Minecraft packet
// ID: 0x09
type PluginMessageServerbound struct {
	Channel string
	Data    []byte `length:"remaining"`
}

// UseEntity is sent when the user interacts (right clicks) or attacks
// (left clicks) an entity.
//
// This is a Minecraft packet
// ID: 0x0A
type UseEntity struct {
	TargetID VarInt
	Type     VarInt
	TargetX  float32 `if:".Type==2"`
	TargetY  float32 `if:".Type==2"`
	TargetZ  float32 `if:".Type==2"`
	Hand     VarInt  `if:".Type==0 .Type==2"`
}

// KeepAliveServerbound is sent by a client as a response to a
// KeepAliveClientbound. If the client doesn't reply the server
// may disconnect the client.
//
// This is a Minecraft packet
// ID: 0x0B
type KeepAliveServerbound struct {
	ID int64
}

// Player is used to update the player's on ground status.
//
// This is a Minecraft packet
// ID: 0x0C
type Player struct {
	OnGround bool
}

// PlayerPosition is used to update the player's position.
//
// This is a Minecraft packet
// ID: 0x0D
type PlayerPosition struct {
	X, Y, Z  float64
	OnGround bool
}

// PlayerPositionLook is a combination of PlayerPosition and
// PlayerLook.
//
// This is a Minecraft packet
// ID: 0x0E
type PlayerPositionLook struct {
	X, Y, Z    float64
	Yaw, Pitch float32
	OnGround   bool
}

// PlayerLook is used to update the player's rotation.
//
// This is a Minecraft packet
// ID: 0x0F
type PlayerLook struct {
	Yaw, Pitch float32
	OnGround   bool
}

// PlayerVehicleMove is used to update the location of the player's vehicle.
//
// This is a Minecraft packet
// ID: 0x10
type PlayerVehicleMove struct {
	X, Y, Z    float64
	Yaw, Pitch float32
}

// SteerBoat is used to update the status of the player's boat.
//
// This is a Minecraft packet
// ID: 0x11
type SteerBoat struct {
	LeftPaddle, RightPaddle bool
}

// CraftRecipeRequest is used to craft items.
//
// This is a Minecraft packet
// ID: 0x12
type CraftRecipeRequest struct {
	WindowID byte
	RecipeID VarInt
	MakeAll  bool
}

// ClientAbilities is used to modify the players current abilities.
// Currently flying is the only one
//
// This is a Minecraft packet
// ID: 0x13
type ClientAbilities struct {
	Flags        byte
	FlyingSpeed  float32
	WalkingSpeed float32
}

// PlayerDigging is sent when the client starts/stops digging a block.
// It also can be sent for droppping items and eating/shooting.
//
// This is a Minecraft packet
// ID: 0x14
type PlayerDigging struct {
	Status   byte
	Location Position
	Face     byte
}

// PlayerAction is sent when a player preforms various actions.
//
// This is a Minecraft packet
// ID: 0x15
type PlayerAction struct {
	EntityID  VarInt
	ActionID  VarInt
	JumpBoost VarInt
}

// SteerVehicle is sent by the client when steers or preforms an action
// on a vehicle.
//
// This is a Minecraft packet
// ID: 0x16
type SteerVehicle struct {
	Sideways float32
	Forward  float32
	Flags    byte
}

// This is a Minecraft packet
// ID: 0x17
type CraftingBookData struct {
	Type               VarInt
	RecipeID           float32 `if:".Type==0"`
	CraftingBookOpen   bool    `if:".Type==1"`
	CraftingBookFilter bool    `if:".Type==1"`
}

// ResourcePackStatus informs the server of the client's current progress
// in activating the requested resource pack
//
// This is a Minecraft packet
// ID: 0x18
type ResourcePackStatus struct {
	Result VarInt
}

// This is a Minecraft packet
// ID: 0x19
type AdvancementTab struct {
	Action VarInt
	TabID  VarInt `if:".Action==0"`
}

// HeldItemChange is sent when the player changes the currently active
// hotbar slot.
//
// This is a Minecraft packet
// ID: 0x1A
type HeldItemChange struct {
	Slot int16
}

// CreativeInventoryAction is sent when the client clicks in the creative
// inventory. This is used to spawn items in creative.
//
// This is a Minecraft packet
// ID: 0x1B
type CreativeInventoryAction struct {
	Slot        int16
	ClickedItem ItemStack `as:"raw"`
}

// SetSign sets the text on a sign after placing it.
//
// This is a Minecraft packet
// ID: 0x1C
type SetSign struct {
	Location Position
	Line1    string
	Line2    string
	Line3    string
	Line4    string
}

// ArmSwing is sent by the client when the player left clicks (to swing their
// arm).
//
// This is a Minecraft packet
// ID: 0x1D
type ArmSwing struct {
	Hand VarInt
}

// SpectateTeleport is sent by clients in spectator mode to teleport to a player.
//
// This is a Minecraft packet
// ID: 0x1E
type SpectateTeleport struct {
	Target UUID `as:"raw"`
}

// PlayerBlockPlacement is sent when the client tries to place a block.
//
// This is a Minecraft packet
// ID: 0x1F
type PlayerBlockPlacement struct {
	Location                  Position
	Face                      VarInt
	Hand                      VarInt
	CursorX, CursorY, CursorZ float32
}

// UseItem is sent when the client tries to use an item.
//
// This is a Minecraft packet
// ID: 0x20
type UseItem struct {
	Hand VarInt
}
