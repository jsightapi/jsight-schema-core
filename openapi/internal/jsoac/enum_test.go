package jsoac

import (
	"testing"
)

func Test_newOpenAPIEnumType(t *testing.T) {
	tests := []struct {
		jsight  string
		openapi string
	}{
		{
			`"white" // { type: "enum", enum: ["white", "blue", "red", ""] }`,
			`{
				"example": "white",
				"enum": ["white", "blue", "red", ""]
			}`,
		},
		{
			`"white" // { enum: ["white", "blue", "red"] }`,
			`{
				"example": "white",
				"enum": ["white", "blue", "red"]
			}`,
		},
		{
			`-2.1 // { type: "enum", enum: [-3, -2.1, 1.2, true, false, "-3", "0", "1.2", "string", "true", null, "null"] }`,
			`{
				"example": -2.1,
				"enum": [-3, -2.1, 1.2, true, false, "-3", "0", "1.2", "string", "true", null, "null"]
			}`,
		},
		{
			`"http://wierd.com/.././%2D?query:=; #%fragment/api" // { "enum": ["http://wierd.com/.././%2D?query:=; #%fragment/api"] }`,
			`{
						"example": "http://wierd.com/.././%2D?query:=; #%fragment/api",
						"enum": ["http://wierd.com/.././%2D?query:=; #%fragment/api"]
					}`,
		},
		{
			`"https://username@[1080:0:0:0:8:800:200C:417A]:80/ABCDEFGHIJKLMNOPQRSTUVWXYZ?abcdefghijklmnopqrstuvwxyz!$&'()*+,;=#0123456789-._~ " // { "enum": ["https://username@[1080:0:0:0:8:800:200C:417A]:80/ABCDEFGHIJKLMNOPQRSTUVWXYZ?abcdefghijklmnopqrstuvwxyz!$&'()*+,;=#0123456789-._~ "] }`,
			`{
						"example": "https://username@[1080:0:0:0:8:800:200C:417A]:80/ABCDEFGHIJKLMNOPQRSTUVWXYZ?abcdefghijklmnopqrstuvwxyz!$&'()*+,;=#0123456789-._~ ",
						"enum": ["https://username@[1080:0:0:0:8:800:200C:417A]:80/ABCDEFGHIJKLMNOPQRSTUVWXYZ?abcdefghijklmnopqrstuvwxyz!$&'()*+,;=#0123456789-._~ "]
					}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.jsight, func(t *testing.T) {
			jsightToOpenAPI(t, tt.jsight, tt.openapi)
		})
	}
}
