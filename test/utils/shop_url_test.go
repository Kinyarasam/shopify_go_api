package utils_test

import (
	"fmt"
	"shopify_go_api/shopify/utils"
	"testing"
)

func TestSanitizeShopDomain(t *testing.T) {
	tests := []struct {
		name            string
		shopDomain      string
		myshopifyDomain string
		expected        string
	}{
		{
			name:            "Good shop domains",
			shopDomain:      "my-shop",
			myshopifyDomain: "myshopify.com",
			expected:        "my-shop.myshopify.com",
		},
		{
			name:            "Valid shop domain with http",
			shopDomain:      "http://my-shop.myshopify.com",
			myshopifyDomain: "myshopify.com",
			expected:        "my-shop.myshopify.com",
		},
		{
			name:            "Valid shop domain with https",
			shopDomain:      "https://my-shop.myshopify.com",
			myshopifyDomain: "myshopify.com",
			expected:        "my-shop.myshopify.com",
		},
		{
			name:            "Bad shop domains",
			shopDomain:      "myshop.com",
			myshopifyDomain: "myshopify.com",
			expected:        "",
		},
		{
			name:            "Custom shop domains",
			shopDomain:      "my-shop",
			myshopifyDomain: "myshopify.io",
			expected:        "my-shop.myshopify.io",
		},
		{
			name:            "Valid custom domain with http",
			shopDomain:      "http://my-shop.myshopify.io",
			myshopifyDomain: "myshopify.io",
			expected:        "my-shop.myshopify.io",
		},
		{
			name:            "None type",
			shopDomain:      "",
			myshopifyDomain: "myshopify.com",
			expected:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.SanitizeShopDomain(tt.shopDomain, tt.myshopifyDomain)
			fmt.Printf("got: %s\n", got)
			if got != tt.expected {
				t.Errorf("SanitizeShopDomain() = %v, want %v", got, tt.expected)
			}
		})
	}
}
