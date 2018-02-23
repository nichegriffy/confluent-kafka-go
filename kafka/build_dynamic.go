// +build !static
// +build !static_all

package kafka

// #cgo darwin linux CFLAGS: -I${SRCDIR}/clib/include
// #cgo darwin linux LDFLAGS: -L${SRCDIR}/clib/lib -lrdkafka
// #cgo !darwin,!linux pkg-config: rdkafka
// #cgo !darwin,!linux LDFLAGS: -lrdkafka
import "C"
