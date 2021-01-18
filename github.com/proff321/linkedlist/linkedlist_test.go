package linkedlist
import "testing"

func TestHello(t *testing.T) {
  got := Hello()
  want := "Hello world"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

// The linked list can 
// - Return the proper count of items
// - Find an item that is stored in the list
// - Store a single item
// - Store multiple items
// - Remove a single item
// -- At the beginning
// -- At the middle
// -- At the end
// ***** Extra ****
// - Store duplicates of the same item
// - Gracefully handle removing a missing item
// - Gracefully remove only the first of a duplicate item

func TestLinkedListReturnsProperCount(t *testing.T) {
  ll := LinkedList{}
  ll.Add("hello")

  got := ll.Count()
  want := 1

  if (got != want) {
    t.Errorf("got %q want %q", got, want)
  }

  ll.Add("world")
  got = ll.Count()
  want = 2

  if (got != want) {
    t.Errorf("got %q want %q", got, want)
  }
}

func TestLinkedListCanFindStoredItem(t *testing.T) {
  ll := LinkedList{}
  ll.Add("hello")

  count, result := ll.Find("hello")

  if (count != 1) {
    t.Errorf("Unable to find 'hello' within linked list")
  }

  if (result != "hello") {
    t.Errorf("Found '%s' instead of '%s'", result, "hello")
  }
}

func TestLinkedListCanStoreMultipleItems(t *testing.T) {
  ll := LinkedList{}
  ll.Add("hello")
  ll.Add("world")

  count, result := ll.Find("hello")

  if (count != 1) {
    t.Errorf("Unable to find 'hello' within linked list")
  }

  want := "hello"
  if (result != want) {
    t.Errorf("Found '%s' instead of '%s'", result, want)
  }

  count, result = ll.Find("world")

  if (count != 1) {
    t.Errorf("Unable to find 'hello' within linked list")
  }

  want = "world"
  if (result != want) {
    t.Errorf("Found '%s' instead of '%s'", result, want)
  }
}

func TestLinkedListCanStoreDuplicateItems(t *testing.T) {
  ll := LinkedList{}
  ll.Add("hello")
  ll.Add("hello")

  count, result := ll.Find("hello")

  if (count !=  2) {
    t.Errorf("Unable to find 'hello' within linked list")
  }

  want := "hello"
  if (result != want) {
    t.Errorf("Found '%s' instead of '%s'", result, want)
  }
}

func TestLinkedListCanRemoveItemFromBeginningOfList(t *testing.T) {
  ll := LinkedList{}
  ll.Add("first")
  ll.Add("second")
  ll.Add("thrid")

  if (ll.Count() != 3) {
    t.Errorf("Linked list should contain three items")
  }

  ll.Remove("first")

  count, result := ll.Find("first")

  if (ll.Count() != 2 || count != 0 || result != "") {
    t.Errorf("Element was not deleted from linked list")
  }
}

func TestLinkedListCanRemoveItemFromMiddleOfList(t *testing.T) {
  ll := LinkedList{}
  ll.Add("first")
  ll.Add("second")
  ll.Add("thrid")

  if (ll.Count() != 3) {
    t.Errorf("Linked list should contain three items")
  }

  ll.Remove("second")

  count, result := ll.Find("second")

  if ( ll.Count() != 2 || count != 0 || result != "" ) {
    t.Errorf("Element was not deleted from linked list")
  }
}

func TestLinkedListCanRemoveItemFromEndOfList(t *testing.T) {
  ll := LinkedList{}
  ll.Add("first")
  ll.Add("second")
  ll.Add("thrid")

  if (ll.Count() != 3) {
    t.Errorf("Linked list should contain three items")
  }

  ll.Remove("thrid")

  count, result := ll.Find("thrid")

  if ( ll.Count() != 2 || count != 0 || result != "" ) {
    t.Errorf("Element was not deleted from linked list")
  }
}

func TestShouldRemoveAllItemsFromLinkedList(t *testing.T) {
  ll := LinkedList{}
  ll.Add("first")
  ll.Add("second")
  ll.Add("thrid")

  if (ll.Count() != 3) {
    t.Errorf("Linked list should contain three items")
  }

  ll.Remove("first")
  ll.Remove("second")
  ll.Remove("thrid")

  if ( ll.Count() != 0) {
    t.Errorf("The linked list still has count %v but should be empty", ll.Count())
  }
}