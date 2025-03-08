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

func ParseGUID(guid string) *GUID {
	// decode base64 string
	decoded, err := base64.StdEncoding.DecodeString(guid)
	if err != nil {
		return nil
	}

	parts := strings.Split(string(decoded), "/")
	if len(parts) != 2 {
		return nil
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil
	}

	return NewGUID(parts[0], id)
}

func (g *GUID) String() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s/%d", g.nodeType, g.id)))
}

func (g *GUID) UnmarshalGQLContext(ctx context.Context, v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("guid must be a string")
	}

	guid := ParseGUID(s)
	if guid == nil {
		return fmt.Errorf("invalid GUID format: %s", s)
	}

	g.nodeType = guid.nodeType
	g.id = guid.id
	return nil
}

func (g GUID) MarshalGQLContext(ctx context.Context, w io.Writer) error {
	guid := strconv.Quote(g.String())
	w.Write([]byte(guid))
	return nil
}

func (g *GUID) NodeType() string {
	return g.nodeType
}

func (g *GUID) Id() int {
	return g.id
}
