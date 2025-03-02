package progbar

import (
	"fmt"
	"testing"
	"time"
	// "github.com/crolbar/lipbalm/assert"
)

func TestMain(t *testing.T) {
	now := time.Now()

	pb := NewProgBar("asnotehutsnaosntheut", 10, 60,
		WithReverse(),
		// WithVertical(),
	)


	pb.SetValue(60)

	v := pb.View()

	fmt.Println("elapsed: ", time.Since(now))

	fmt.Println(v)
	// assert.Equal(t, 2, 3)
}
