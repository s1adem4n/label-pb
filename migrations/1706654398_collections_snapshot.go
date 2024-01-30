package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "mlrj8f154zedj6k",
				"created": "2024-01-19 12:00:11.343Z",
				"updated": "2024-01-21 22:05:33.064Z",
				"name": "products",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "srdqha4o",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "crplp6rv",
						"name": "owner",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "vinpx9wr",
						"name": "image",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "eiyalqor8dch1pq",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "bd0pwunz",
						"name": "notes",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = owner.id",
				"viewRule": "@request.auth.id = owner.id",
				"createRule": "@request.auth.verified = true",
				"updateRule": "@request.auth.id = owner.id",
				"deleteRule": "@request.auth.id = owner.id",
				"options": {}
			},
			{
				"id": "rsybzzivhxze7wn",
				"created": "2024-01-19 12:03:14.099Z",
				"updated": "2024-01-21 22:05:16.990Z",
				"name": "batches",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "8txodyml",
						"name": "manufactured",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "0fn4elvo",
						"name": "expires",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "gnae0yyl",
						"name": "quantity",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": true
						}
					},
					{
						"system": false,
						"id": "sddr8x9a",
						"name": "notes",
						"type": "editor",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "dsydohlp",
						"name": "owner",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "9ozinjox",
						"name": "product",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "mlrj8f154zedj6k",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = owner.id",
				"viewRule": "@request.auth.id = owner.id",
				"createRule": "@request.auth.verified = true",
				"updateRule": "@request.auth.id = owner.id",
				"deleteRule": "@request.auth.id = owner.id",
				"options": {}
			},
			{
				"id": "eiyalqor8dch1pq",
				"created": "2024-01-19 12:04:24.388Z",
				"updated": "2024-01-21 22:05:07.797Z",
				"name": "images",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "jsekr64o",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mcxriayi",
						"name": "file",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/png",
								"image/jpeg"
							],
							"thumbs": [
								"500x500f"
							],
							"maxSelect": 1,
							"maxSize": 100000000,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "sdkoeqko",
						"name": "owner",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = owner.id",
				"viewRule": "@request.auth.id = owner.id",
				"createRule": "@request.auth.verified = true",
				"updateRule": "@request.auth.id = owner.id",
				"deleteRule": "@request.auth.id = owner.id",
				"options": {}
			},
			{
				"id": "_pb_users_auth_",
				"created": "2024-01-30 20:38:12.667Z",
				"updated": "2024-01-30 20:48:40.233Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "njkjdggt",
						"name": "admin",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": false,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"onlyVerified": false,
					"requireEmail": false
				}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
