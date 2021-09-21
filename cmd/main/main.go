package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"

	"github.com/evatix-go/rediswrapper/redisdefaults"
)

func main() {
	wrapper := redisdefaults.NewClientWrapper()
	collection1 := corestr.NewCollection(100)

	collection1.Add("alim 1").
		Add("alim 2").
		Add("alim 3")
	fmt.Println(collection1)
	charMap := collection1.HashsetAsIs()
	fmt.Println(charMap.String())
	wrapper.SaveStrHashset("B", charMap)
	collect2 := wrapper.GetStrHashset("B")
	collect2.ErrorWrapper.HandleError()

	fmt.Println(collect2)

	fmt.Println(wrapper.GetString("A").Value)

	wrapper.DeleteManyKeys("alist", "aSet1")
	wrapper.AddListStringItems("alist", "alim 1")

	fmt.Println(wrapper.ListAsCollection("alist").Collection)

	wrapper.AddUniqueAnysToKey("aSet1", "alim 1", "alim 2")

	fmt.Println("Aset", wrapper.UniqueListAsCollection("aSet1").Collection)

	akey2 := "akey2"
	wrapper.MarshalSaveMany(akey2, collect2, collect2, collect2)
	fmt.Println(akey2, wrapper.GetString(akey2))

	collect3 := corestr.EmptyHashset()

	wrapper.GetUnmarshalMany(akey2, collect3)

	fmt.Println(collect3)

}
