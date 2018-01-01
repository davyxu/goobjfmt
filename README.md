# goobjfmt

Marshal struct input text format( protobuf text )

like 
	key : value
	
	key has no ""
	
	string value should have ""

# Example

```golang
	input := &MyData{
		Name:   "genji",
		Type:   MyCar_Pig,
		Uint32: math.MaxUint32,
		Int64:  math.MaxInt64,
		Uint64: math.MaxUint64,
	}

	t.Log(MarshalTextString(input))
```

# Feedback

Star me if you like or use, thanks

开源讨论群: 527430600

知乎: [http://www.zhihu.com/people/sunicdavy](http://www.zhihu.com/people/sunicdavy)
