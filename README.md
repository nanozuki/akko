# Akko

An annotation and code generation based web framework

## Usage guide

### First CRUD Example

1. Just Write your actual logic function

   ```go
   type Book struct {
       ID       int64    `json:"id"`
       Name     string   `json:"name"`
       AuthorID int64    `json:"author_id"`
       ISBN     string   `json:"ISBN"`
       Tags     []string `json:"tags"`
   }

   type BookPatch struct {
       Name   string
       Author int64
       ISBN   string
       NewTag string
   }

   type BookFilter struct {
       Tag   string `json:"tag"`
       Limit int    `json:"limit`
   }

   func (a *Applicaton) GetBookByID(ctx context.Context, id int64) (*Book, error)
   func (a *Applicaton) ListBooks(ctx context.Context, filter *BookFilter) ([]*Book, error)
   func (a *Applicaton) CreateBook(ctx context.Context, book *Book) (*Book, error)
   func (a *Applicaton) UpdateBook(ctx context.Context, id int64, patch *BookPatch) (*Book, error)
   func (a *Applicaton) DeleteBook(ctx context.Context, id int64) error
   ```

1. add annotations to logic function

   ```go
   // GetBookByID find a book by bookID, if book not exist, return NotFound error
   // [GET=/books/:id]
   func (a *Applicaton) GetBookByID(ctx context.Context, id int64) (*Book, error)
   // ListBooks get a list of books through filter, default limit is 10
   // [GET=/books query->filter]
   func (a *Applicaton) ListBooks(ctx context.Context, filter *BookFilter) ([]*Book, error)
   // Or you can define function like this
   // [GET=/books query.tag->tag,limit]
   func (a *Applicaton) ListBooks(ctx context.Context, tag string, limit int) ([]*Book, error)
   // [POST=/books body|json->book]
   func (a *Applicaton) CreateBook(ctx context.Context, book *Book) (*Book, error)
   // [PUT=/books/:id body|json->patch, return.1|json->body]
   func (a *Applicaton) UpdateBook(ctx context.Context, id int64, patch *BookPatch) (*Book, error)
   // [DELETE=/books/:id]
   func (a *Applicaton) DeleteBook(ctx context.Context, id int64) error
   ```

1. response

   Akko will try to find these three method for response and error:

   ```go
   func (r Response) Status() int
   func (r Response) Body()   io.Reader
   func (r Response) Header() http.Header
   ```

   At default, will treat non-err response to json with 200 status code, and convert error to error string with 500.

## Annotation Pipeline Specification

### syntax:

_Input_[. _Field_][| *Process*] -> _Destination_[. _Field_]

### Input

Input for request:

| input        | return type               | note                                                |
| ------------ | ------------------------- | --------------------------------------------------- |
| request      | \*http.Request            |                                                     |
| param        | []\*akko.Param            | ordered list                                        |
| param.`a`    | string                    | can be ignore when the name is equal to Destination |
| query        | url.Values                | alias to 'map[string][]string'                      |
| query.`a`    | string, []string          |                                                     |
| header       | http.Header               | alias to 'map[string][]string'                      |
| header.`a`   | string, []string          |                                                     |
| cookie       | []\*http.Cookie           |                                                     |
| cookie.`a`   | string, \*http.Cookie     |                                                     |
| body         | []byte                    |                                                     |
| body \| json | - (ready for json decode) |                                                     |
| body \| form | url.Values                |                                                     |

Input for response:

| input          | return type              | note                                                      |
| -------------- | ------------------------ | --------------------------------------------------------- |
| return         | handler return type      |                                                           |
| return.`a`     | return's field or method |                                                           |
| return \| json | []byte                   | Marshal return value to json, set application/json header |
| error          | [T akko.Error]           | error of handler returned                                 |

_TODO: MultipartForm_

relative types:

```go
type Param struct {
    Key string
    Value string
}
```

### Destination

#### Fixed destinations:

| destination | receive type                   | note                                    |
| ----------- | ------------------------------ | --------------------------------------- |
| body        | []byte, io.Writer              |                                         |
| status      | int                            | Http status code. Default is 200 or 500 |
| header      | http.Header, map[string]string |                                         |
| header.`a`  | string                         |                                         |
| cookie.`a`  | \*http.Cookie                  |                                         |

#### Handler parameters destination:

Each parameters (except context.Context) of handle can be a destination, for example:

```go
// [GET=/books query.tag->tag,limit]
func (a *Applicaton) ListBooks(ctx context.Context, tag string, limit int) ([]*Book, error)
```

the destination can be 'ctx', 'tag', 'limit'. To avoid name conflict, the params can not be 'body', 'status', 'header', 'cookie'.

#### Type convert

The type of destination can be the input type, or can be convert from input type. For different inputs, the convertible destination type 'T' needs to satisfy the following conditions:

1. \*http.Request

   T must implemented interface akko.RequestReceiver:

   ```go
   type RequestReceiver interface {
       ParseRequest(*http.Request) error
   }
   ```

1. \*akko.Param

   T must implmented interface akko.ParamUnmarshaler:

   ```go
   type ParamUnmarshaler interface {
       UnmarshalParam([]*akko.Param) error
   }
   ```

1. string

   T can be []byte, all integer types, and boolean, or implemented encoding.TextUnmarshaler

1. []string

   T can be `[]U`, which U is "string convertible", or T implemented akko.StringsUnmarshaler:

   ```go
   type StringsUnmarshaler interface {
       UnmarlshalStrings([]string) error
   }
   ```

1. map[string][]string, url.Values, http.Header

   T can be a struct, which every fields are "[]string convertiable", or T implemented akko.ValuesUnmarshaler:

   ```go
   type ValuesUnmarshaler interface {
       UnmarlshalValues(map[string][]string) erro
   }

   ```

1. \*http.Cookie

   T can be a struct, which every fields are "string convertible", or T implemented akko.CookieUnmarshaler:

   ```go
   type CookieUnmarshaler {
       UnmarshalCookie([]*http.Cookie) error
   }
   ```

1. body

   T must implemented encoding.TextUnmarshaler

1. body|json

   T is a type can be use to json.Unmarshal
