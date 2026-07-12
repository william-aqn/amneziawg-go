//go:build !android && !ios && !windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2025 WireGuard LLC. All Rights Reserved.
 */

package device

import "github.com/amnezia-vpn/amneziawg-go/v3/conn"

const (
	QueueStagedSize    = conn.IdealBatchSize
	QueueOutboundSize  = 1024
	QueueInboundSize   = 1024
	QueueHandshakeSize = 1024
	MaxSegmentSize     = (1 << 16) - 1 // largest possible UDP datagram
)

// PreallocatedBuffersPerPool is a var instead of a const (like on ios), so that
// memory-constrained hosts — e.g. consumer routers running the standalone daemon —
// can bound the buffer pools without patching the source. Assign it before calling
// NewDevice (the daemon reads WG_PREALLOCATED_BUFFERS_PER_POOL from the environment;
// 0 means unbounded — upstream's stock Linux behavior).
// router-build default: bounded at 1024 (upstream's iOS/Windows profile) — WaitPool.Get
// gives real backpressure instead of `runtime: out of memory` under sustained inbound
// on low-RAM routers; high-RAM boxes get 0 from their launcher for full throughput.
var PreallocatedBuffersPerPool uint32 = 1024
