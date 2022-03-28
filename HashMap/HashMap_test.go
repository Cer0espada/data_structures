package hashmap

import "testing"

var testKeys = []string{"foo", "bar", "croo", "octopus", "rave", "What time is it?"}
var testValues = []string{"bar", "foo", "12:45", "hammer", "that stuff", "party time"}
func TestAddAndGet(t *testing.T){

	t.Run("testing nilcase", func(t *testing.T){
			var HM HashMap[string, string]

		val, err := HM.Get("foobar")

		if err == nil{
			t.Errorf("Expected to get an Error, we got a value of: %v", val)
		}
	})
	
	var HM HashMap[string, string]

	for index, test := range testKeys{
		HM.Add(test,testValues[index])

		val, err := HM.Get(test)

		if err != nil{
			t.Error(err)
		}
		
		if HM.Size != index +1 {
			t.Errorf("Size not being maintained as expected, Expected: %v, got: %v",index+1, HM.Size)
		}

		if testValues[index] != val{
			t.Errorf("Expected to get the same value")
		}
	}
}

func TestSearch(t *testing.T){

	t.Run("testing nilcase", func(t *testing.T){
		var HM HashMap[string, string]
		HM.Add("poo", "snoo")

		val := HM.Search("croo")

		if val != false{
			t.Errorf("expected to see false as the return value in the nilcase")
		}
	})

	var HM HashMap[string, string]

	for index, test := range testKeys{
		HM.Add(test,testValues[index])

		val := HM.Search(test)

		
		if HM.Size != index +1 {
			t.Errorf("Size not being maintained as expected, Expected: %v, got: %v",index+1, HM.Size)
		}

		if !val{
			t.Errorf("Expected value to be present in the hash table ")
		}
	}
}

func TestDelete(t *testing.T){

	t.Run("testing nil case", func(t *testing.T){
		var HM HashMap[string, string]

		err := HM.Delete("stuff")

		if err == nil{
			t.Error("Expected error in nilcase")
		}

	})

	var HM HashMap[string, string]

	for index, test := range testKeys{
		HM.Add(test,testValues[index])

		err := HM.Delete(test)

		if err != nil{
			t.Error(err)
		}

		if HM.Size != 0 {
			t.Errorf("Size not being maintained as expected, Expected: %v, got: %v",0, HM.Size)
		}

		if HM.Size != 0 {
			t.Errorf("Size")
		}
	}

}