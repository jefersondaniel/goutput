# Goutput

```go
import "github.com/jefersondaniel/goutput/output"
import "github.com/jefersondaniel/goutput/output/resources"
import "github.com/revel/revel"

type Notification struct {
    Text string
    Url string
}

func (self *Notification) SetText(text string) {
    self.Text = text
}

func (self *Notification) SetUrl(url string) {
    self.Url = url
}

func (c NotificationController) Create() revel.Result {
    notification := Notification{}
    notification.SetText("Hello")
    notification.SetUrl("http://url")

    transformerFunction := func (data interface{}) map[string]interface{} {
        notification := data.(*entity.Notification)

        return map[string]interface{}{
            "text": notification.GetText(),
            "link": notification.GetLink(),
        }
    }

    manager := output.CreateManager()
    transformer := output.Transformer{transformerFunction}
    resource := resources.CreateItemResource("data", notification, transformer)

    return c.RenderJSON(manager.CreateData(resource).GetData())
}
```
