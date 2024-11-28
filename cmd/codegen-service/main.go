package main

import "codegen-service/internal/engine"

// сделать проверку на обратную совместимость с предыдущей версией и
// уточнять у пользователя точно ли он хочет внести изменения

const contract = `{
  "general": {
    "id": 1,
    "name": "testapi",
    "owner": "aaalace",
    "version": "1.0",
    "port": 8070,
    "autoAuth": false
  },
  "models": [
    {
      "name": "Book",
      "fields": [
        {
          "name": "id",
          "type": "int",
          "isUnique": true
        },
        {
          "name": "title",
          "type": "string",
          "isUnique": false
        },
        {
          "name": "isFinished",
          "type": "bool",
          "isUnique": false
        }
      ],
      "methods": [
        {
          "type": "GET",
          "uniqueParam": "id",
          "responseFields": ["id", "title", "isFinished"],
          "privateEndpoint": false
        },
        {
          "type": "GET*",
          "responseFields": ["id", "title"],
          "privateEndpoint": false
        },
        {
          "type": "POST",
          "responseFields": ["id", "title"],
          "privateEndpoint": true
        },
        {
          "type": "DELETE",
          "responseFields": ["id"],
          "privateEndpoint": true
        }
      ]
    }
  ]
}`

func main() {
	eng := engine.NewEngine(contract)
	saf := eng.ParseSourceToSAF()
	// is saf.AutoAuth => generate auth
	eng.Generator.GenerateMain(&saf.General)
}
