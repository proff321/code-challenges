package linkedlist

import "fmt"

func Hello() string {
  return "Hello world"
}

func main() {
	fmt.Println(Hello())
}

type node struct {
  value string
  next *node
}

type LinkedList struct {
  root *node
}

func (ll *LinkedList) Add (newValue string) {
  if (ll.root == nil) {
    ll.root = &node{value: newValue}
  } else {
    // Get to the last node
    iterator := ll.root
    for ; iterator.next != nil; iterator = iterator.next {
    }

    iterator.next = &node{value: newValue}
  }
}

func (ll *LinkedList) Remove (removeValue string) {
  var previous *node  
  for iterator := ll.root; iterator != nil;  {
    if (iterator.value == removeValue) {
      if (previous != nil) {
        if(iterator.next != nil) {
          // Middle
          previous.next = iterator.next
        } else {
          // End
          previous.next = nil
        }
      } else {
        if (iterator.next != nil) {
          // First
          ll.root = iterator.next
        } else {
          // Last
          ll.root = nil
        }
      }

      return
    }
    previous = iterator
    iterator = iterator.next
  }
}

func (ll LinkedList) Count () int {

  total := 0
  iterator := ll.root
  for ; iterator != nil; iterator = iterator.next {
    total++
  }

  return total;
}

func (ll LinkedList) Find (searchTerm string) (numberFound int, result string) {
  for iterator := ll.root; iterator != nil; iterator = iterator.next {
    if (iterator.value == searchTerm) {
      numberFound++
      result = iterator.value
    }
  }
  return
}
