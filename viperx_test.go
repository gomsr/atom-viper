package viperx

import (
	"testing"
)

func TestInitViper(_ *testing.T) {
	InitViper(&struct{}{}, "config.yaml")
}
