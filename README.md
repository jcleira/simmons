# simmons

[![circle-ci](https://circleci.com/gh/jcleira/simmons.png?style=shield&circle-token=163296a336ca296ee3b0d76152f6db0bb2859a76)](https://circleci.com/gh/jcleira/simmons)

simmons is an ORM / database abstraction for data driven applications.

![ORM](https://i.imgur.com/HilP6Tp.png)

simmons use case is about not trusting an ORM ([gorm.io](http://gorm.io) on this case), and provides a real and a mocked interface implementation for the following CRUD operations:

- Create
- Update
- Delete
- All
- Find (by id)
- Query (by where clause)

## Usage

```go
type User struct {
  Name string
}

simmons.Init(conn)

user := &User {
  Name: "Simmons",
}

err := simmons.Create(user)

user.Name = "bar"
err := simmons.Update(user)

users := []*User{}
err := simmons.All(users)

users := []*User{}
err := simmons.Query(users, "name like '%sim'")

err := simmons.Find(user, 1)

err := simmons.Update(user)
``` 
