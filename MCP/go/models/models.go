package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Chartpostrequest represents the Chartpostrequest schema from the OpenAPI specification
type Chartpostrequest struct {
	Height int `json:"height,omitempty"` // The height of the chart in pixels.
	Width int `json:"width,omitempty"` // The width of the chart in pixels.
	Backgroundcolor string `json:"backgroundColor,omitempty"` // The background color of the chart.
	Chart map[string]interface{} `json:"chart,omitempty"` // The chart configuration in Chart.js format (JSON or Javascript).
	Format string `json:"format,omitempty"` // The output format of the chart, 'png', 'jpg', 'svg', or 'webp'.
}

// Qrpostrequest represents the Qrpostrequest schema from the OpenAPI specification
type Qrpostrequest struct {
	Height int `json:"height,omitempty"` // The height of the QR code in pixels.
	Margin int `json:"margin,omitempty"` // The margin around the QR code in pixels.
	Text string `json:"text,omitempty"` // The text to be encoded in the QR code.
	Width int `json:"width,omitempty"` // The width of the QR code in pixels.
	Format string `json:"format,omitempty"` // The output format of the QR code, e.g., 'png' or 'svg'.
}
