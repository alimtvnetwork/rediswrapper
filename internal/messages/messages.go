package messages

var (
	StringCmdNull = "string / status / string slice cmd : no status returned after " +
		"saving or reading the bytes from redis server, may be communication lost between redis server."
	IntCmdNull               = "int cmd : no status returned after saving or reading the bytes from redis server, may be communication lost between redis server."
	JsonResultNullPointer    = "JsonResult nil pointer or empty, cannot process it."
	CannotUnmarshallEmpty    = "cannot unmarshal empty bytes to object."
	CannotUnmarshallNilItems = "cannot unmarshal nil items."
	CannotMarshallNilOrEmpty = "cannot marshal empty or nil object to serialized data."
)
