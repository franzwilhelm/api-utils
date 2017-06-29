package base

import (
	"strconv"

	"github.com/astaxie/beego"
)

// ResponseController is used for all ResponseControllers and contains the most basic elements
type ResponseController struct {
	beego.Controller
	InputURL string
	Code     int
	Res      struct {
		Href    string                 `json:"href,omitempty"`
		Message string                 `json:"message,omitempty"`
		Size    int                    `json:"size,omitempty"`
		Offset  int                    `json:"offset,omitempty"`
		Limit   int                    `json:"limit,omitempty"`
		Links   []ResponseLink         `json:"links,omitempty"`
		Items   []interface{}          `json:"items,omitempty"`
		Errors  map[string]interface{} `json:"errors,omitempty"`
	}
}

// ResponseLink is used to reference other rescources in the API
type ResponseLink struct {
	Rel    string `json:"rel,omitempty"`
	Method string `json:"method"`
	Href   string `json:"href" `
}

// Prepare is used to prepare the ResponseController
func (c *ResponseController) Prepare() {
	c.Res.Errors = make(map[string]interface{})
	c.InputURL = c.Ctx.Input.URL()
}

// SetErrorText sets error based on code and string
func (c *ResponseController) SetErrorText(code int, msg string) {
	c.Code = code
	c.Res.Message = msg
	c.SetOutput()
}

// SetError sets error based on code and error interface
func (c *ResponseController) SetError(code int, err error) {
	c.Code = code
	c.Res.Message = err.Error()
	c.SetOutput()
}

// SetOutput sets JSON output formatted for web
func (c *ResponseController) SetOutput() {
	c.Res.Href = c.InputURL
	c.Data["json"] = c.Res
	c.Ctx.Output.SetStatus(c.Code)
	c.ServeJSON()
	c.StopRun()
}

// AppendLinks appends links to the response
func (c *ResponseController) AppendLinks(links ...ResponseLink) {
	c.Res.Links = append(c.Res.Links, links...)
}

// AppendItems appends items to the response
func (c *ResponseController) AppendItems(items ...interface{}) {
	c.Res.Items = append(c.Res.Items, items...)
}

// AppendOffsetLimitSize appends offset and limit to output next/previous/all-links
func (c *ResponseController) AppendOffsetLimitSize(offset, limit, size int) {
	var next, previous int
	link := ResponseLink{Method: "GET"}

	// set size, offset and limit in res, if value is 0, it will get omited
	c.Res.Size = size
	c.Res.Offset = offset
	c.Res.Limit = limit
	limitString := strconv.Itoa(limit)

	// set previous offset if it exists
	if offset > 0 {
		if offset-limit > 0 {
			previous = offset - limit
		} else {
			previous = 0
		}
		link.Rel = "previous"
		link.Href = c.InputURL + "?offset=" + strconv.Itoa(previous) + "limit=" + limitString
		c.AppendLinks(link)
	}

	// set next offset if it exists
	if offset+limit < size {
		next = offset + limit
		link.Rel = "next"
		link.Href = c.InputURL + "?offset=" + strconv.Itoa(next) + "limit=" + limitString
		c.AppendLinks(link)
	}

	c.AppendLinks(ResponseLink{Rel: "all", Method: "GET", Href: c.InputURL + "?limit=" + strconv.Itoa(size)})
}

// GetOffsetLimit checks input offset/limit, and sets default values if not found
func (c *ResponseController) GetOffsetLimit() (offset, limit int) {
	var err error
	if offset, err = c.GetInt("offset"); err != nil {
		offset = 0
	}
	if limit, err = c.GetInt("limit"); err != nil {
		limit = 10
	}
	return
}
