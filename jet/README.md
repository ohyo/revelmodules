# Jet Template Engine for Revel
 The [Jet](https://github.com/CloudyKit/jet) Templating Plugin

 Jet is a template engine developed to be easy to use, powerful, dynamic, yet secure and very fast.

 - supports template inheritance with extends, import and include statements
 - descriptive error messages with filename and line number
 - auto-escape
 - simple C-like expressions
 - very fast execution – Jet can execute templates faster than some pre-compiled template engines
 - very light in terms of allocations and memory footprint
 - simple and familiar syntax
 - easy to use

 Profit: 
 - Extremely fast
 - Perfect syntax
 - Correct (When i developed blog module the pongo2 is working not correct for me)
 - Extendable
 - Custom renderers for model is really cool

## Jet wiki
 The [Jet Wiki](https://github.com/CloudyKit/jet/wiki) Templating Plugin

##### Details
Extending Jet Module
```
import (
  "ohyo.network/modules/jet"
)
....
// Should extend after module registered
// use onAppStart hook or AddInitEventHandler on 
// REVEL_AFTER_MODULES_LOADED State

jet.Engine.AddGlobalFunc("base64", func(a jet.Arguments) reflect.Value {
  a.RequireNumOfArguments("base64", 1, 1)

  buffer := bytes.NewBuffer(nil)
  fmt.Fprint(buffer, a.Get(0))

  return reflect.ValueOf(base64.URLEncoding.EncodeToString(buffer.Bytes()))
})
....

```