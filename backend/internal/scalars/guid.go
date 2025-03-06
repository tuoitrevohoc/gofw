package scalars

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type GUID struct {
	nodeType string
	id       int
}

func NewGUID(nodeType string, id int) *GUID {
	return &GUID{
		nodeType: nodeType,
		id:       id,
	}
}

func (g *GUID) String() string {
	return fmt.Sprintf("%s/%d", g.nodeType, g.id)
}

func (g *GUID) UnmarshalGQLContext(ctx context.Context, v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("guid must be a string")
	}
	// parse base64 string
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return fmt.Errorf("invalid GUID format: %s", s)
	}
	parts := strings.Split(string(decoded), "/")
	if len(parts) != 2 {
		return fmt.Errorf("invalid GUID format: %s", s)
	}
	g.nodeType = parts[0]
	g.id, err = strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid GUID format: %s", s)
	}

	return nil
}

func (g GUID) MarshalGQLContext(ctx context.Context, w io.Writer) error {
	guid := g.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(guid))
	w.Write([]byte(strconv.Quote(encoded)))
	return nil
}

func (g *GUID) NodeType() string {
	return g.nodeType
}

func (g *GUID) Id() int {
	return g.id
}
