package ononoki

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nanozuki/ononoki/prop"
	"github.com/nanozuki/ononoki/typ"
)

func TestSchema_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		schema  *Schema
		wantErr bool
	}{
		{
			name:   "marshal schema",
			schema: newSchema(t),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.schema.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Schema.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("json: %s", got)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Schema.MarshalJSON() = %v, want %v", got, tt.want)
			// }
		})
	}
}

type PolicyListItem struct {
	ID             int
	Amount         float64
	TrackingNumber string
}

func newSchema(t *testing.T) *Schema {
	s := New("CustomerPolicyApi")
	g1 := s.Group("/v1/admin", "admin_auth").Tag("admin")
	{
		g1.GET("GetPolicyList", "/policies").
			Request(
				prop.String("shop_id").Required().Validator("numeric")).
			Response(
				prop.Array("policy_ids", typ.Int()),
				prop.Array("policies", typ.Object("PolicyItem",
					prop.Int("id").Maximum(10000000).Minimum(99999999),
					prop.Float("amount"),
					prop.String("tracking_number")))).
			Use("admin_customize")
	}
	g2 := s.Group("/v1/customer", "customer_auth").Tag("customer")
	{
		g2.POST("CreateShopPolicyList", "/shop/:shop_id/policies").
			Request(
				prop.String("shop_id").InPath().Validator("numeric"),
				prop.String("tracking_number").Required()).
			Response(
				prop.Array("policies", typ.GoType(PolicyListItem{}))).
			Use("customer_customize")
		g2.GET("GetPolicyDetail", "/policies/:id").
			Request(
				prop.String("id").InPath().Title("policy id")).
			Response(
				prop.Object("policy", typ.Object("PolicyDetail",
					prop.String("policy_id"),
					prop.Float("coverage_amount"),
					prop.String("tracking_number"))),
				prop.String("id").Title("policy id"))
	}

	// to json
	bytes, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("schema: %s\n", string(bytes))
	var ss Schema
	if err := json.Unmarshal(bytes, &ss); err != nil {
		t.Fatal(err)
	}
	return &ss
}
