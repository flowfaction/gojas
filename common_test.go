package gojas


var asserted_json_data string = `{
		"user":{
			"properties":{
				"string":{
					"type": "string",
					"value": "foobar"
				},
				"number":{
					"type": "number",
					"value": 42
				},
				"boolean":{
					"type": "boolean",
					"value": true
				},
				"numberArray":{
					"type": "array",
					"value": [1,1,2,3,5,8]
				},
				"stringArray":{
					"type": "array",
					"value": ["1","1","2","3","5","8"]
				},
				"object":{
					"type": "object",
					"innerObject": {
						"foo" : "bar",
						"baz": 11235,
						"key" : "value"
					}
				}
			},
			"propertypeer": {
				"peer": 1
			}
		},
		"userpeer" :  {
			"peer" : 1
		}
}`


var bad_json_data string = `{
		"user":{
			"properties":{
				"object":{
					"type": "object"
					"innerObject": {
						"foo" : "bar",
						"baz": 11235,
						"key" : "value"
					}
				}
			},
		},
		"userpeer" :  {
			"peer" : 1
		}
}`
