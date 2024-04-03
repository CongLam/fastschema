# Introduction

FastSchema is an open-source headless Content Management System (CMS) designed to simplify the creation and management of structured content. By leveraging schema definitions, FastSchema automates the generation of databases and provides CRUD (Create, Read, Update, Delete) APIs effortlessly.

> **Note:** FastSchema is currently in beta and under active development. We welcome feedback, contributions, and suggestions from the community to help improve the platform and make it more robust and feature-rich.


## Overview

At the core of FastSchema lies its schema definition, a blueprint that outlines the structure of your content. This schema acts as the foundation upon which FastSchema builds your database tables and API endpoints, streamlining the development process and allowing you to focus on creating rich, dynamic content.

<p style="text-align: center;">
  <img src="https://fastschema.com/img/fastschema.png" alt="FastSchema Overview" />
</p>

## Features

Fastschema offers a comprehensive suite of features designed to streamline and simplify the process of building and managing dynamic web applications. Whether you're a developer, designer, or content creator, our platform provides the tools you need to create, deploy, and maintain powerful web experiences with ease.

- Automated Database Generation.
- RESTful API Generation.
- Dynamic Content Modeling.
- Built-in File Management.
- Built-in Admin Control Panel.
- Database Support: MySQL, PostgreSQL, SQLite.
- Role-Based Access Control.



FastSchema simplifies the process of building and managing structured content, providing developers with a powerful toolset to create dynamic, data-driven applications. With its schema-driven approach, automated database generation, and CRUD API creation, FastSchema accelerates development workflows and empowers teams to focus on delivering exceptional digital experiences.

Get started with FastSchema today and revolutionize the way you manage content in your applications!

## Documentation

For more information on how to get started with FastSchema, check out our [documentation](https://fastschema.com).

### Schema Definition

The schema definition is structured JSON that encapsulates the characteristics of your content model. Let's take a closer look at a sample schema definition:

**post.json**

```json
{
  "name": "post",
  "namespace": "posts",
  "label_field": "name",
  "fields": [
    {
      "type": "string",
      "name": "name",
      "label": "Name",
      "sortable": true
    },
    {
      "type": "relation",
      "name": "category",
      "label": "Category",
      "renderer": {},
      "relation": {
        "schema": "category",
        "field": "posts",
        "type": "o2m"
      },
    }
  ]
}
```

**category.json**

```json
{
  "name": "category",
  "namespace": "categories",
  "label_field": "name",
  "fields": [
    {
      "type": "string",
      "name": "name",
      "label": "Name",
      "optional": true,
      "sortable": true
    },
    {
      "type": "text",
      "name": "content",
      "label": "Content",
      "renderer": {
        "class": "editor"
      },
      "optional": true,
      "sortable": true
    },
    {
      "type": "relation",
      "name": "posts",
      "label": "Posts",
      "optional": true,
      "relation": {
        "schema": "post",
        "field": "category",
        "type": "o2m",
        "owner": true,
        "optional": true
      }
    }
  ]
}
```

### Example

#### Query

```
GET /api/users/?sort=-age&select=name,email,groups.name&filter={filterObject}
```

```json
{
  "name": {
    "$like": "test%",
    "$neq": "test2"
  },
  "$or": [
    {
      "email": {
        "$neq": "test",
        "$like": "test%"
      },
      "age": {
        "$lt": 10
      }
    },
    {
      "age": 5
    },
    {
      "$and": [
        {
          "name": {
            "$neq": "test2"
          }
        },
        {
          "age": 5
        }
      ]
    }
  ]
}

```

#### Update

```
PUT /api/users/1
```

```json
{
  "name": "John Doe",
  "age": 30,
  "room": { "id": 2 },
  "pets": [ { "id": 2 }, { "id": 3 } ],
  "groups": [ { "id": 4 }, { "id": 5 } ],
  "$set": {
    "bio": "Hello World",
    "address": "123 Main St",
    "sub_room": { "id": 2 },
    "sub_pets": [ { "id": 2 }, { "id": 3 } ],
    "sub_groups": [ { "id": 4 }, { "id": 5 } ]
  },
  "$clear": {
    "bio": true,
    "address": true,
    "room": true,
    "sub_pets": true,
    "sub_groups": true,
    "pets": [ { "id": 2 }, { "id": 3 } ],
    "groups": [ { "id": 4 }, { "id": 5 } ]
  },
  "$add": {
    "pets": [ { "id": 2 }, { "id": 3 } ],
    "groups": [ { "id": 4 }, { "id": 5 } ],
    "age": 1,
    "salary": 1000
  },
  "$expr": {
    "bio": "LOWER(`bio`)",
    "address": "CONCAT(`address`, ' ', `city`, ' ', `state`, ' ', `zip`)"
  }
}
```

## Extend

FastSchema is a flexible and extensible application that allows you to customize and extend its functionality to meet your specific requirements. This guide provides an overview of the different ways you can extend FastSchema, including customizing the API, adding new features, and integrating with third-party services.

### Using FastSchema as a module

```go
package main

import (
	"github.com/fastschema/fastschema"
	"github.com/fastschema/fastschema/app"
	"github.com/fastschema/fastschema/db"
	"github.com/fastschema/fastschema/schema"
)

func main() {
	newApp, err := fastschema.New(&fastschema.AppConfig{})

	if err != nil {
		panic(err)
	}

	newApp.AddResource(
		app.NewResource("home", func(c app.Context, _ *any) (any, error) {
			return "Welcome to fastschema", nil
		}, app.Meta{app.GET: ""}),
	)

	newApp.OnAfterDBContentList(
    func(query *db.QueryOptions, entities []*schema.Entity) ([]*schema.Entity, error) {
      if query.Model.Schema().Name != "media" {
        return entities, nil
      }

      for _, entity := range entities {
        entity.Set("custom", true)
      }

      return entities, nil
    },
  )

	newApp.Start()
}
```


## Roadmap

* [ ] Improve documentation and testing.
* [ ] Add auth provider.
* [ ] Plugin system.
* [ ] OpenAPI generator.
* [ ] Real-time updates.
* [ ] GraphQL support.
* [ ] Webhooks.
* [ ] Client SDKs.


## Testing

FastSchema comes with a suite of automated tests to ensure the stability and reliability of the platform. To run the tests, simply execute the following command:

```bash
go test ./...
```

## Dependencies

FastSchema is built using the Go programming language and leverages a number of open-source libraries to provide its core functionality. Some of the key dependencies include:

- [Fiber](https://gofiber.io/)
- [Ent](https://entgo.io/)
- [Rclone](https://rclone.org/)
- [Zap](https://pkg.go.dev/go.uber.org/zap)
- [Next.js](https://nextjs.org/)
- [Shadcn](https://ui.shadcn.com/)
- [TipTap](https://www.tiptap.dev/)

## Contributing

We welcome contributions from the community and encourage developers to get involved in the project. Whether you're a seasoned developer or just getting started, there are plenty of ways to contribute to FastSchema.