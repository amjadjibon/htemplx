<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# repo

```go
import "htemplx/app/repo"
```

## Index

- [Constants](<#constants>)
- [type ContactsRepo](<#ContactsRepo>)
  - [func NewContactsRepo\(dbx \*dbx.DBX\) \*ContactsRepo](<#NewContactsRepo>)
  - [func \(c \*ContactsRepo\) CreateContacts\(ctx context.Context, contacts \*models.ContactUs\) error](<#ContactsRepo.CreateContacts>)
- [type UsersRepo](<#UsersRepo>)
  - [func NewUsersRepo\(dbx \*dbx.DBX\) \*UsersRepo](<#NewUsersRepo>)
  - [func \(u \*UsersRepo\) CreateUser\(ctx context.Context, user \*models.User\) error](<#UsersRepo.CreateUser>)
  - [func \(u \*UsersRepo\) DeleteUser\(ctx context.Context, id string\) error](<#UsersRepo.DeleteUser>)
  - [func \(u \*UsersRepo\) GetUserByEmail\(ctx context.Context, email string\) \(\*models.User, error\)](<#UsersRepo.GetUserByEmail>)
  - [func \(u \*UsersRepo\) GetUserByID\(ctx context.Context, id string\) \(\*models.User, error\)](<#UsersRepo.GetUserByID>)
  - [func \(u \*UsersRepo\) GetUserList\(ctx context.Context\) \(\[\]\*models.User, error\)](<#UsersRepo.GetUserList>)
  - [func \(u \*UsersRepo\) UpdateUser\(ctx context.Context, user \*models.User\) error](<#UsersRepo.UpdateUser>)


## Constants

<a name="ContactUsTable"></a>

```go
const ContactUsTable = "contact_us"
```

<a name="UsersTableName"></a>

```go
const UsersTableName = "users"
```

<a name="ContactsRepo"></a>
## type [ContactsRepo](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/contacts.go#L14-L16>)



```go
type ContactsRepo struct {
    // contains filtered or unexported fields
}
```

<a name="NewContactsRepo"></a>
### func [NewContactsRepo](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/contacts.go#L18>)

```go
func NewContactsRepo(dbx *dbx.DBX) *ContactsRepo
```



<a name="ContactsRepo.CreateContacts"></a>
### func \(\*ContactsRepo\) [CreateContacts](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/contacts.go#L22>)

```go
func (c *ContactsRepo) CreateContacts(ctx context.Context, contacts *models.ContactUs) error
```



<a name="UsersRepo"></a>
## type [UsersRepo](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/users.go#L11-L13>)



```go
type UsersRepo struct {
    // contains filtered or unexported fields
}
```

<a name="NewUsersRepo"></a>
### func [NewUsersRepo](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/users.go#L15>)

```go
func NewUsersRepo(dbx *dbx.DBX) *UsersRepo
```



<a name="UsersRepo.CreateUser"></a>
### func \(\*UsersRepo\) [CreateUser](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/users.go#L21>)

```go
func (u *UsersRepo) CreateUser(ctx context.Context, user *models.User) error
```



<a name="UsersRepo.DeleteUser"></a>
### func \(\*UsersRepo\) [DeleteUser](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/users.go#L82>)

```go
func (u *UsersRepo) DeleteUser(ctx context.Context, id string) error
```



<a name="UsersRepo.GetUserByEmail"></a>
### func \(\*UsersRepo\) [GetUserByEmail](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/users.go#L95>)

```go
func (u *UsersRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error)
```



<a name="UsersRepo.GetUserByID"></a>
### func \(\*UsersRepo\) [GetUserByID](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/users.go#L49>)

```go
func (u *UsersRepo) GetUserByID(ctx context.Context, id string) (*models.User, error)
```



<a name="UsersRepo.GetUserList"></a>
### func \(\*UsersRepo\) [GetUserList](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/users.go#L35>)

```go
func (u *UsersRepo) GetUserList(ctx context.Context) ([]*models.User, error)
```



<a name="UsersRepo.UpdateUser"></a>
### func \(\*UsersRepo\) [UpdateUser](<https://github.com/amjadjibon/htemplx/blob/main/app/repo/users.go#L64>)

```go
func (u *UsersRepo) UpdateUser(ctx context.Context, user *models.User) error
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)