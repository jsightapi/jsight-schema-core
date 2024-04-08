package jsoac

import (
	"testing"
)

func Test_description(t *testing.T) {
	tests := []testConverterData{

		{
			`{} // Long description with different symbols... tabs:        ascii: !\"$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~ spec:¹²³$‰↑∞←→—≠€®™ѣѵіѳ′[]≈§°£₽„“”‘’×©↓−«»…´“„•ẞˇ¢·¸¨‘⌀⌘˘˚ѢѴІѲ″{}±–〉〈¿ˆ¼⅓½¡ Cyrillic: йцукенгшщзФЫВАПРОЛДЖЭ Chinese: 你好世界！`,
			`{
				"type": "object",
				"properties": {},
				"additionalProperties": false,
				"description": "Long description with different symbols... tabs: ascii: !\\\"$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~ spec:¹²³$‰↑∞←→—≠€®™ѣѵіѳ′[]≈§°£₽„“”‘’×©↓−«»…´“„•ẞˇ¢·¸¨‘⌀⌘˘˚ѢѴІѲ″{}±–〉〈¿ˆ¼⅓½¡ Cyrillic: йцукенгшщзФЫВАПРОЛДЖЭ Chinese: 你好世界！"
			}`,
		},

		{
			`1 // Long description with different symbols... tabs:        ascii: !\"$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~ spec:¹²³$‰↑∞←→—≠€®™ѣѵіѳ′[]≈§°£₽„“”‘’×©↓−«»…´“„•ẞˇ¢·¸¨‘⌀⌘˘˚ѢѴІѲ″{}±–〉〈¿ˆ¼⅓½¡ Cyrillic: йцукенгшщзФЫВАПРОЛДЖЭ Chinese: 你好世界！`,
			`{
				"type": "integer",
				"example": 1,
				"description": "Long description with different symbols... tabs: ascii: !\\\"$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~ spec:¹²³$‰↑∞←→—≠€®™ѣѵіѳ′[]≈§°£₽„“”‘’×©↓−«»…´“„•ẞˇ¢·¸¨‘⌀⌘˘˚ѢѴІѲ″{}±–〉〈¿ˆ¼⅓½¡ Cyrillic: йцукенгшщзФЫВАПРОЛДЖЭ Chinese: 你好世界！"
			}`,
		},

		{
			`1  /* Multiline
						 annotation
						 in several lines. 
						 Long description with different symbols... tabs:        ascii: !\"$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~ spec:¹²³$‰↑∞←→—≠€®™ѣѵіѳ′[]≈§°£₽„“”‘’×©↓−«»…´“„•ẞˇ¢·¸¨‘⌀⌘˘˚ѢѴІѲ″{}±–〉〈¿ˆ¼⅓½¡ Cyrillic: йцукенгшщзФЫВАПРОЛДЖЭ Chinese: 你好世界！
					  */`,
			`{
				"type": "integer",
				"example": 1,
				"description": "Multiline annotation in several lines. Long description with different symbols... tabs: ascii: !\\\"$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~ spec:¹²³$‰↑∞←→—≠€®™ѣѵіѳ′[]≈§°£₽„“”‘’×©↓−«»…´“„•ẞˇ¢·¸¨‘⌀⌘˘˚ѢѴІѲ″{}±–〉〈¿ˆ¼⅓½¡ Cyrillic: йцукенгшщзФЫВАПРОЛДЖЭ Chinese: 你好世界！"
			}`,
		},

		{
			`1 // Any description string & "quoted string" & \*\/ \*\/`,
			`{
				"type": "integer",
				"example": 1,
				"description": "Any description string & \"quoted string\" & \\*\\/ \\*\\/"
			}`,
		},
		{
			`1 // {min: -99, max: 99} - Some note.`,
			`{
				"type": "integer",
				"example": 1,
				"minimum": -99,
				"maximum": 99,
				"description": "Some note."
			}`,
		},
		{
			`1 // Some note.`,
			`{
				"type": "integer",
				"example": 1,
				"description": "Some note."
			}`,
		},
		{
			`1  /* 
            	Some note.
			*/`,
			`{
				"type": "integer",
				"example": 1,
				"description": "Some note."
			}`,
		},
		{
			`1  /* Multiline
						 annotation
						 in several lines. 
					  */`,
			`{
				"type": "integer",
				"example": 1,
				"description": "Multiline annotation in several lines."
			}`,
		},
		{
			`{
				"prt1": 1 // Type integer.
			}`,
			`{
    			"type": "object",
    			"properties": {
        			"prt1": {
            			"type": "integer",
            			"example": 1,
            			"description": "Type integer."
        			}
    			},
    			"required": [
        			"prt1"
    			],
    			"additionalProperties": false
			}`,
		},
	}
	for _, data := range tests {
		t.Run(data.name(), func(t *testing.T) {
			assertJSightToOpenAPIConverter(t, data)
		})
	}
}
