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
   // [GET=/books query->tag,limit]
   func (a *Applicaton) ListBooks(ctx context.Context, tag string, limit int) ([]*Book, error)
   // [POST=/books body.json->book]
   func (a *Applicaton) CreateBook(ctx context.Context, book *Book) (*Book, error)
   // [PUT=/books/:id body.json->patch, return.1.json->body]
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
