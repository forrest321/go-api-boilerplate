package domain_test

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/vardius/go-api-boilerplate/pkg/common/domain"
)

func ExampleContextWithFlag() {
	ctx := context.Background()
	ctx = domain.ContextWithFlag(ctx, "test")

	fmt.Printf("%v\n", domain.HasFlag(ctx, "test"))

	// Output:
	// true
}

func ExampleHasFlag() {
	ctx := context.Background()

	fmt.Printf("%v\n", domain.HasFlag(ctx, "test"))

	ctx = domain.ContextWithFlag(ctx, "test")

	fmt.Printf("%v\n", domain.HasFlag(ctx, "test"))

	// Output:
	// false
	// true
}

func ExampleFlagsFromContext() {
	ctx := context.Background()
	flags := domain.FlagsFromContext(ctx)

	fmt.Printf("%v\n", flags)

	ctx = domain.ContextWithFlag(ctx, "foo")
	ctx = domain.ContextWithFlag(ctx, "bar")
	flags = domain.FlagsFromContext(ctx)

	sort.Strings(flags)
	fmt.Printf("%v\n", flags)

	// Output:
	// []
	// [bar foo]
}

func ExampleNewEvent() {
	type Test struct {
		Page   int      `json:"page"`
		Fruits []string `json:"fruits"`
	}

	event, _ := domain.NewEvent(
		uuid.New(),
		"streamName",
		0,
		Test{1, []string{"apple", "peach"}},
	)

	fmt.Printf("%v\n", event.Metadata.StreamName)
	fmt.Printf("%v\n", event.Metadata.StreamVersion)
	fmt.Printf("%s\n", event.Payload)

	// Output:
	// streamName
	// 0
	// {"page":1,"fruits":["apple","peach"]}
}

func ExampleMakeEvent() {
	event, _ := domain.MakeEvent(
		domain.EventMetaData{
			Type:          "type",
			StreamID:      uuid.New(),
			StreamName:    "streamName",
			StreamVersion: 0,
			OccurredAt:    time.Now(),
		},
		[]byte(`{"page":1,"fruits":["apple","peach"]}`),
	)

	fmt.Printf("%v\n", event.Metadata.StreamName)
	fmt.Printf("%v\n", event.Metadata.StreamVersion)
	fmt.Printf("%s\n", event.Payload)

	// Output:
	// streamName
	// 0
	// {"page":1,"fruits":["apple","peach"]}
}
